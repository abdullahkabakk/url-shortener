package database

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMySQLConnection(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
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
		assert.NotNil(t, db)
		assert.NoError(t, err)
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

	t.Run("Failed to connect for non-existing database", func(t *testing.T) {
		connector := &DBConnector{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			DBName:   "non-existing",
		}

		_, err := ConnectToDB(connector, "not-existing")

		assert.Error(t, err)

	})

	t.Run("Failed to migrate", func(t *testing.T) {
		// Create a mock database connection
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error creating mock database connection: %v", err)
		}
		defer db.Close()

		// Set up expectations for the mock database query to ensure that the migration fails
		mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").WillReturnError(fmt.Errorf("error"))

		// Call the migrations function
		err = migrations(db)
		if err == nil {
			t.Error("expected an error, got nil")
		}

	})
}

func TestMigrations(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database connection: %v", err)
	}
	defer db.Close()

	t.Run("Migrate Successfully", func(t *testing.T) {
		// Set up expectations for the mock database query to ensure that the migration is successful
		mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec("CREATE TABLE IF NOT EXISTS urls").WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec("CREATE TABLE IF NOT EXISTS clicks").WillReturnResult(sqlmock.NewResult(1, 1))

		// Call the migrations function
		err = migrations(db)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Failed to Migrate", func(t *testing.T) {
		// Set up expectations for the mock database query to ensure that the migration fails
		mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").WillReturnError(fmt.Errorf("error"))

		// Call the migrations function
		err = migrations(db)
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})
}

func TestCreateDatabase(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database connection: %v", err)
	}
	defer db.Close()

	t.Run("Create Database Successfully", func(t *testing.T) {
		// Set up expectations for the mock database query to ensure that the database is created successfully
		mock.ExpectExec("CREATE DATABASE IF NOT EXISTS test").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("USE test").WillReturnResult(sqlmock.NewResult(1, 1))

		// Call the createDatabase function
		err := createDatabase(db, "test")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Failed to Create Database", func(t *testing.T) {
		// Set up expectations for the mock database query to ensure that the database creation fails
		mock.ExpectExec("CREATE DATABASE IF NOT EXISTS test").WillReturnError(fmt.Errorf("error"))

		// Call the createDatabase function
		err := createDatabase(db, "test")
		assert.Error(t, err)
	})

	t.Run("Failed to Use Database", func(t *testing.T) {
		// Set up expectations for the mock database query to ensure that the database usage fails
		mock.ExpectExec("CREATE DATABASE IF NOT EXISTS test").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("USE test").WillReturnError(fmt.Errorf("error"))

		// Call the createDatabase function
		err := createDatabase(db, "test")
		assert.Error(t, err)
	})
}
