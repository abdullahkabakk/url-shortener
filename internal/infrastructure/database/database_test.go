package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestMySQLConnection(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	t.Run("Connect to MySQL Database", func(t *testing.T) {
		connector := &DBConnector{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			DBName:   os.Getenv("DB_NAME"),
		}

		db, err := ConnectToDB(connector, "mysql")
		if err != nil {
			return
		}
		defer db.Close()
	})

	t.Run("Failed to Connect to MySQL Database", func(t *testing.T) {
		connector := &DBConnector{
			Username: "root",
			Password: "password",
			Host:     "localhost",
			Port:     "3306",
			DBName:   "test",
		}

		_, err := ConnectToDB(connector, "mysql")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
