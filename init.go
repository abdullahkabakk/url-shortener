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

func initializeApp() {
	// Start a goroutine to periodically display memory usage
	go func() {
		for {
			printMemoryUsage()
			time.Sleep(5 * time.Second) // Adjust the interval as needed
		}
	}()

	// Create DBConnector with environment variables
	connector := config.NewDBConnector()
	dbDriver := "mysql"

	// Connect to the database
	db, err := database.ConnectToDB(connector, dbDriver)
	if err != nil {
		fmt.Println("[MAIN] Error connecting to database:", err)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("[MAIN] Error closing database connection:", err)
		}
	}(db)

	// Create auth handler
	userHandler, urlHandler, clicksHandler := initializeHandlers(db)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	server := http.NewServer(host, port, userHandler, urlHandler, clicksHandler)

	err = server.Start()
	if err != nil {
		return
	}
}
