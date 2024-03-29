package auth_service

import (
	"golang.org/x/crypto/bcrypt"
	"url-shortener/internal/app/models/user"
	"url-shortener/internal/app/repositories/auth"
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
func (s *Service) CreateUser(user user_model.User) (*user_model.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword) // Convert hashed password to string

	// Call the auth repository to create the auth
	userVal, err := s.Repository.Create(&user)
	if err != nil {
		return nil, err
	}

	return userVal, nil
}

// LoginUser authenticates the auth with the provided username and password.
// It returns a token upon successful authentication.
func (s *Service) LoginUser(user user_model.User) (*user_model.User, error) {
	// Retrieve auth from the database
	userVal, err := s.Repository.GetByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(userVal.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	return userVal, nil
}
