package clicks_repository

import (
	"database/sql"
)

// Repository defines methods to interact with the URL repository.
type Repository interface {
	CreateClick(shortURL, ipAddress string) error
}

// DBClicksRepository is an implementation of ClicksRepository for MySQL database.
type DBClicksRepository struct {
	// DB is the database connection
	DB *sql.DB
}

// NewDBClicksRepository creates a new instance of DBClicksRepository.
func NewDBClicksRepository(db *sql.DB) *DBClicksRepository {
	return &DBClicksRepository{DB: db}
}

// CreateClick inserts a new click record into the database.
func (r *DBClicksRepository) CreateClick(shortURL, ipAddress string) error {
	// Prepare SQL statement
	stmt, err := r.DB.Prepare("INSERT INTO clicks (shortened_url, ip_address) VALUES (?, ?)")
	if err != nil {
		return err
	}

	// Defer closing the prepared statement
	defer stmt.Close()

	// Execute SQL statement
	_, err = stmt.Exec(shortURL, ipAddress)
	if err != nil {
		return err
	}

	return nil
}
