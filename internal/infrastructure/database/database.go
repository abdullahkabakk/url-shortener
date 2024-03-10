package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Connector defines an interface for connecting to a database.
type Connector interface {
	Connect(driverName string) (*sql.DB, error)
}

// DBConnector implements Connector for connecting to a MySQL database.
type DBConnector struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

// Connect establishes a connection to the MySQL database using the provided credentials and driver name.
func (m *DBConnector) Connect(driverName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", m.Username, m.Password, m.Host, m.Port)
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("[DB Connection] Failed to connect to Database: %w", err)
	}

	// Ping the database to ensure connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping Database: %v", err)
	}

	// Create tables if they don't exist
	if err := createDatabase(db, m.DBName); err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}

	// Run migrations
	if err := migrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migration queries: %v", err)
	}

	return db, nil
}

// ConnectToDB connects to the database using the provided connector and driver name.
func ConnectToDB(connector Connector, driverName string) (*sql.DB, error) {
	return connector.Connect(driverName)
}

func createDatabase(db *sql.DB, DBName string) error {
	// Create database if it doesn't exist
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + DBName)
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}

	// Use the specified database
	_, err = db.Exec("USE " + DBName)
	if err != nil {
		return fmt.Errorf("failed to use database: %v", err)
	}

	// Ping the database again after reconnecting
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping Database after reconnect: %v", err)
	}
	return nil
}

func migrations(db *sql.DB) error {
	// Define migration queries
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
			password VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`,
		`CREATE TABLE IF NOT EXISTS urls (
			id INT AUTO_INCREMENT PRIMARY KEY,
			original_url TEXT NOT NULL,
			shortened_url VARCHAR(10) NOT NULL UNIQUE,
			user_id INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
			);`,
		`CREATE TABLE IF NOT EXISTS clicks (
			id INT AUTO_INCREMENT PRIMARY KEY,
			url_id INT NOT NULL,
			ip_address VARCHAR(50) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (url_id) REFERENCES urls(id)
			);`,
	}

	// Execute queries
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to execute query: %v", err)
		}
	}

	return nil
}
