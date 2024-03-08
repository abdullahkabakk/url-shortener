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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("[DB Connection] Failed to connect to MySQL: %w", err)
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	return db, nil
}

// ConnectToDB connects to the database using the provided connector and driver name.
func ConnectToDB(connector Connector, driverName string) (*sql.DB, error) {
	return connector.Connect(driverName)
}
