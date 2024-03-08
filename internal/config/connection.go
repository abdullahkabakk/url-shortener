package config

import (
	"os"
	"url-shortener/internal/infrastructure/database"
)

// NewDBConnector creates a new DBConnector with the provided parameters.
func NewDBConnector() *database.DBConnector {
	return &database.DBConnector{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
