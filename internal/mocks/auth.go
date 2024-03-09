package mocks

import (
	"url-shortener/internal/app/models/user"
)

// MockUserRepository is a mock implementation of UserRepository interface for testing purposes.
type MockUserRepository struct {
	Users map[uint]*user_model.User
}

// NewMockUserRepository creates a new instance of MockUserRepository.
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		Users: make(map[uint]*user_model.User),
	}
}

// Create simulates creating a new auth in the mock database.
func (r *MockUserRepository) Create(user *user_model.User) (*user_model.User, error) {
	// username should be unique
	for _, u := range r.Users {
		if u.Username == user.Username {
			return nil, user_model.ErrUserAlreadyExists
		}

	}
	user.ID = uint(len(r.Users) + 1) // Simulate auto-incrementing ID
	r.Users[user.ID] = user
	return user, nil
}

// GetByUsername simulates retrieving an auth by username from the mock database.
func (r *MockUserRepository) GetByUsername(username string) (*user_model.User, error) {
	// DB error can be simulated here
	if username == "error" {
		return nil, user_model.ErrUserNotFound
	}

	//  Check if auth exists
	for _, user := range r.Users {
		if user.Username == username {
			// Return the auth
			return user, nil
		}
	}

	// Return an error if auth not found
	return nil, user_model.ErrUserNotFound
}
