package clicks_repository

import (
	"database/sql"
	"url-shortener/internal/app/models/clicks"
)

// Repository defines methods to interact with the URL repository.
type Repository interface {
	CreateClick(shortURL, ipAddress string) error
	GetClicks(shortURL string) ([]clicks_model.Clicks, error)
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
	stmt, err := r.DB.Prepare("INSERT INTO clicks (url_id, ip_address) VALUES (?, ?)")
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

// GetClicks retrieves the clicks for the given shortened URL.
func (r *DBClicksRepository) GetClicks(shortURL string) ([]clicks_model.Clicks, error) {
	// Prepare SQL statement
	stmt, err := r.DB.Prepare("SELECT * FROM clicks WHERE url_id = ?")
	if err != nil {
		return nil, err
	}

	// Defer closing the prepared statement
	defer stmt.Close()

	// Execute SQL statement
	rows, err := stmt.Query(shortURL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	var clicks = make([]clicks_model.Clicks, 0)
	for rows.Next() {
		var clicks_m clicks_model.Clicks
		err := rows.Scan(&clicks_m.ID, &clicks_m.UrlID, &clicks_m.IPAddress, &clicks_m.CreatedAt)
		if err != nil {
			return nil, err
		}
		clicks = append(clicks, clicks_m)
	}

	return clicks, nil
}
