package auth_repository

import (
	"database/sql"
	"errors"
	"url-shortener/internal/app/models/user"
)

// Repository defines methods to interact with the auth repository.
type Repository interface {
	Create(user *user_model.User) (*user_model.User, error)
	GetByUsername(username string) (*user_model.User, error)
}

// DBAuthRepository is an implementation of UserRepository for MySQL database.
type DBAuthRepository struct {
	// DB is the database connection
	DB *sql.DB
}

// NewDBAuthRepository creates a new instance of DBUserRepository.
func NewDBAuthRepository(db *sql.DB) *DBAuthRepository {
	return &DBAuthRepository{DB: db}
}

// Create inserts a new auth record into the database.
func (r *DBAuthRepository) Create(user *user_model.User) (*user_model.User, error) {
	// Prepare SQL statement
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	// Defer closing the prepared statement
	defer stmt.Close()

	// Execute SQL statement
	result, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return nil, err
	}

	// Retrieve the ID of the newly inserted auth
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Set the ID of the auth object
	user.ID = uint(userID)

	return user, nil
}

// GetByUsername retrieves an auth record from the database by username.
func (r *DBAuthRepository) GetByUsername(username string) (*user_model.User, error) {
	// Prepare SQL statement
	query := "SELECT id, username, password, created_at FROM users WHERE username = ?"
	row := r.DB.QueryRow(query, username)

	// Initialize a new User object to store the result
	user := &user_model.User{}

	// Scan the result into the User object
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Return a custom error if the auth with the specified username is not found
			return nil, user_model.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
