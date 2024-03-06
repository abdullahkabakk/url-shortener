package db

import (
	"database/sql"
	"fmt"
)

// Connector defines an interface for connecting to a database.
type Connector interface {
	Connect() (*sql.DB, error)
}

// MySQLConnector implements Connector for connecting to a MySQL database.
type MySQLConnector struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

// Connect establishes a connection to the MySQL database using the provided credentials.
func (m *MySQLConnector) Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", m.Username, m.Password, m.Host, m.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("[DB Connection] Failed to connect to MySQL: %w", err)
	}

	// Ensure database exists
	if err := m.createDatabaseIfNotExists(db, m.DBName); err != nil {
		return nil, fmt.Errorf("[DB Connection] Failed to ensure database existence: %w", err)
	}

	return db, nil
}

func (m *MySQLConnector) createDatabaseIfNotExists(db *sql.DB, dbname string) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	return err
}

// ConnectToDB connects to the database using the provided credentials.
func ConnectToDB(connector Connector) (*sql.DB, error) {
	return connector.Connect()
}
