package main

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"time"
	"url-shortener/internal/app/handlers"
	"url-shortener/internal/config"
	"url-shortener/internal/infrastructure/database"
	"url-shortener/internal/infrastructure/database/migrations"
	"url-shortener/internal/infrastructure/http"

	"github.com/joho/godotenv"
)

func main() {
	// Start a goroutine to periodically display memory usage
	go func() {
		for {
			printMemoryUsage()
			time.Sleep(5 * time.Second) // Adjust the interval as needed
		}
	}()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[MAIN] Error loading .env file:", err)
		return
	}

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

	fmt.Println("[DATABASE] Connected to database")

	// Run migrations
	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	err = migrations.RunMigrations(db, migrationsDir)
	if err != nil {
		fmt.Println("[MAIN] Error running migrations:", err)
		return
	}

	// Create auth handler
	userHandler, err := handlers.InitializeUserHandlers(db)
	if err != nil {
		fmt.Println("[MAIN] Error initializing auth handlers:", err)
		return
	}

	urlHandler, err := handlers.InitializeURLHandlers(db)

	if err != nil {
		fmt.Println("[MAIN] Error initializing auth handlers:", err)
		return
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	server := http.NewServer(host, port, userHandler, urlHandler)

	err = server.Start()
	if err != nil {
		return
	}

	fmt.Println("[SERVER] Server started at", host+":"+port)
}

func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("Memory Usage:")
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
