package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"url-shortener/db"
)

// NewMySQLConnector creates a new MySQLConnector with the provided parameters.
func NewMySQLConnector() *db.MySQLConnector {
	return &db.MySQLConnector{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// Create MySQLConnector with environment variables
	connector := NewMySQLConnector()

	// Connect to the database
	database, err := db.ConnectToDB(connector)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}(database)

	fmt.Println("Connected to database")
}
