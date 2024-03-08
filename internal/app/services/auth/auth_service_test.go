package auth_service

import (
	"testing"
	"url-shortener/internal/app/models"
	"url-shortener/internal/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mockUserRepository := mocks.NewMockUserRepository()

	// Create a new instance of AuthService with the mock repository
	userService := NewAuthService(mockUserRepository)

	t.Run("Create User Successfully", func(t *testing.T) {
		// Call the CreateUser method
		token, err := userService.CreateUser("testuser", "password123")

		// Assertions
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

	})

	t.Run("Failed to Create User", func(t *testing.T) {
		// Call the CreateUser method
		token, err := userService.CreateUser("testuser", "password123")

		// Assertions
		assert.Error(t, err)
		// Token should be empty
		assert.Empty(t, token)
	})
}

func TestLoginUser(t *testing.T) {
	mockUserRepository := mocks.NewMockUserRepository()

	// Create a new instance of AuthService with the mock repository
	userService := NewAuthService(mockUserRepository)

	user := &models.User{Username: "test", Password: "password123"}

	t.Run("Login User Successfully", func(t *testing.T) {
		_, err := userService.CreateUser(user.Username, user.Password)
		assert.NoError(t, err)

		// Should login successfully
		token, err := userService.LoginUser(user.Username, user.Password)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("Failed to Login User", func(t *testing.T) {
		// Should fail to login
		token, err := userService.LoginUser(user.Username, "wrongpassword")
		assert.Error(t, err)
		assert.Empty(t, token)

	})

	t.Run("User not found", func(t *testing.T) {

		// Should fail to login
		token, err := userService.LoginUser("unknown", "password")
		assert.Error(t, err)
		assert.Empty(t, token)

	})

}
