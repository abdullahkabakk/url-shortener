package mocks

import (
	"url-shortener/internal/app/models"
)

// MockUserRepository is a mock implementation of UserRepository interface for testing purposes.
type MockUserRepository struct {
	Users map[uint]*models.User
}

// NewMockUserRepository creates a new instance of MockUserRepository.
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		Users: make(map[uint]*models.User),
	}
}

// Create simulates creating a new auth in the mock database.
func (r *MockUserRepository) Create(user *models.User) (*models.User, error) {
	// username should be unique
	for _, u := range r.Users {
		if u.Username == user.Username {
			return nil, models.ErrUserAlreadyExists
		}

	}
	user.ID = uint(len(r.Users) + 1) // Simulate auto-incrementing ID
	r.Users[user.ID] = user
	return user, nil
}

// GetByUsername simulates retrieving a auth by username from the mock database.
func (r *MockUserRepository) GetByUsername(username string) (*models.User, error) {
	// DB error can be simulated here
	if username == "error" {
		return nil, models.ErrUserNotFound
	}

	//  Check if auth exists
	for _, user := range r.Users {
		if user.Username == username {
			// Return the auth
			return user, nil
		}
	}

	// Return an error if auth not found
	return nil, models.ErrUserNotFound
}
