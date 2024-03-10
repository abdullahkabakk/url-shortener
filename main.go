package main

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/infrastructure/database"
	"url-shortener/internal/infrastructure/http"
)

func main() {
	// Start a goroutine to periodically display memory usage
	go func() {
		for {
			printMemoryUsage()
			time.Sleep(5 * time.Second) // Adjust the interval as needed
		}
	}()

	// Create DBConnector with environment variables
	connector := config.NewDBConnector()

	// Connect to the database
	db, err := database.ConnectToDB(connector, "mysql")
	if err != nil {
		fmt.Println("[MAIN] Error connecting to database:", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("[MAIN] Error closing database connection:", err)
		}
	}(db)

	// Create auth handler
	userHandler, urlHandler, clicksHandler := initializeHandlers(db)

	server := http.NewServer(os.Getenv("HOST"), os.Getenv("PORT"), userHandler, urlHandler, clicksHandler)

	if err := server.Start(); err != nil {
		fmt.Println("[MAIN] Error starting server:", err)
	}
}
