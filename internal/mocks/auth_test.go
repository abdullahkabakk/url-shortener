package mocks

import (
	"errors"
	"testing"
	"url-shortener/internal/app/models/user"

	"github.com/stretchr/testify/assert"
)

func TestMockUserRepository_Create(t *testing.T) {
	repo := NewMockUserRepository()

	t.Run("Create User Successfully", func(t *testing.T) {
		user := &user_model.User{
			Username: "testuser",
			Password: "password123",
		}

		createdUser, err := repo.Create(user)

		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, uint(1), createdUser.ID)
		assert.Equal(t, "testuser", createdUser.Username)
		assert.Equal(t, "password123", createdUser.Password)
	})

	t.Run("Failed to Create User with Existing Username", func(t *testing.T) {
		user := &user_model.User{
			Username: "testuser",
			Password: "password123",
		}

		// Create a auth with the same username
		_, err := repo.Create(user)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, user_model.ErrUserAlreadyExists))
	})
}

func TestMockUserRepository_GetByUsername(t *testing.T) {
	repo := NewMockUserRepository()

	t.Run("Get User Successfully", func(t *testing.T) {
		user := &user_model.User{
			ID:       1,
			Username: "testuser",
			Password: "password123",
		}

		// Add the auth to the repository
		repo.Users[user.ID] = user

		// Retrieve the auth by username
		foundUser, err := repo.GetByUsername(user.Username)

		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, user, foundUser)
	})

	t.Run("Failed to Get User with Non-existent Username", func(t *testing.T) {
		// Attempt to retrieve a auth with a non-existent username
		_, err := repo.GetByUsername("nonexistent")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, user_model.ErrUserNotFound))
	})

	t.Run("Failed to Get User with Error", func(t *testing.T) {
		// Attempt to retrieve a auth with a non-existent username
		_, err := repo.GetByUsername("error")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, user_model.ErrUserNotFound))
	})

}
