package url_repository

import (
	"database/sql"
	"errors"
	"url-shortener/internal/app/models/url"
)

// Repository defines methods to interact with the URL repository.
type Repository interface {
	CreateURL(originalURL, shortCode string, userId *uint) (string, error)
	GetOriginalURL(shortCode string) (string, error)
}

// DBURLRepository is an implementation of URLRepository for MySQL database.
type DBURLRepository struct {
	// DB is the database connection
	DB *sql.DB
}

// NewDBURLRepository creates a new instance of DBURLRepository.
func NewDBURLRepository(db *sql.DB) *DBURLRepository {
	return &DBURLRepository{DB: db}
}

// CreateURL inserts a new URL record into the database.
func (r *DBURLRepository) CreateURL(originalURL, shortCode string, userID *uint) (string, error) {
	// Prepare SQL statement
	query := ""
	if userID != nil {
		query = "INSERT INTO urls (original_url, shortened_url, user_id) VALUES (?, ?, ?)"
	} else {
		query = "INSERT INTO urls (original_url, shortened_url) VALUES (?, ?)"
	}

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return "", err
	}

	// Defer closing the prepared statement
	defer stmt.Close()

	// Execute SQL statement
	var result sql.Result
	if userID != nil {
		result, err = stmt.Exec(originalURL, shortCode, *userID)
	} else {
		result, err = stmt.Exec(originalURL, shortCode)
	}
	if err != nil {
		return "", err
	}

	// Check the number of rows affected to ensure successful insertion
	numRows, err := result.RowsAffected()
	if err != nil {
		return "", err
	}
	if numRows == 0 {
		return "", errors.New("no rows affected, insertion failed")
	}

	return shortCode, nil
}

// GetOriginalURL retrieves the original URL from the database by short code.
func (r *DBURLRepository) GetOriginalURL(shortCode string) (string, error) {
	// Prepare SQL statement
	query := "SELECT original_url FROM urls WHERE shortened_url = ?"
	row := r.DB.QueryRow(query, shortCode)

	// Initialize a string to store the result
	var originalURL string

	// Scan the result into the originalURL string
	err := row.Scan(&originalURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Return a custom error if the URL with the specified short code is not found
			return "", url_model.ErrURLNotFound
		}
		return "", err
	}

	return originalURL, nil
}
