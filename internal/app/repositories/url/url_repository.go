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
	GetUserURLs(userID uint) ([]url_model.URL, error)
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

// GetUserURLs retrieves the URLs created by the given user.
func (r *DBURLRepository) GetUserURLs(userID uint) ([]url_model.URL, error) {
	// Prepare SQL statement
	query := "SELECT * FROM urls WHERE user_id = ?"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the result
	var urls []url_model.URL

	// Iterate through the rows and scan the result into URL objects
	for rows.Next() {
		var u url_model.URL
		err := rows.Scan(&u.ID, &u.OriginalURL, &u.ShortenedURL, &u.UserID, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}

	return urls, nil
}
