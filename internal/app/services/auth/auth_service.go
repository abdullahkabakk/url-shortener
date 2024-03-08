package auth_service

import (
	"golang.org/x/crypto/bcrypt"
	"url-shortener/internal/app/models"
	"url-shortener/internal/app/repositories/auth"
	"url-shortener/internal/helpers"
)

// Service handles operations related to users.
type Service struct {
	// UserRepository is an interface that defines methods to interact with the auth repository.
	Repository auth_repository.Repository
}

// NewAuthService creates a new instance of AuthService with the given auth repository.
func NewAuthService(AuthRepository auth_repository.Repository) *Service {
	return &Service{Repository: AuthRepository}
}

// CreateUser creates a new auth with the provided data.
// It hashes the password before storing it in the database.
func (s *Service) CreateUser(username, password string) (string, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create the auth with the hashed password
	user := &models.User{
		Username: username,
		Password: string(hashedPassword), // Convert hashed password to string
	}

	// Call the auth repository to create the auth
	_, err = s.Repository.Create(user)
	if err != nil {
		return "", err
	}

	// Generate a token for the authenticated auth
	token, err := helpers.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// LoginUser authenticates the auth with the provided username and password.
// It returns a token upon successful authentication.
func (s *Service) LoginUser(username, password string) (string, error) {
	// Retrieve auth from the database
	user, err := s.Repository.GetByUsername(username)
	if err != nil {
		return "", err
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// Generate a token for the authenticated auth
	token, err := helpers.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
