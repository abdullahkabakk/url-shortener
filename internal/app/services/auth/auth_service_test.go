package auth_service

import (
	"testing"
	"url-shortener/internal/app/models/user"
	"url-shortener/internal/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mockUserRepository := mocks.NewMockUserRepository()

	// Create a new instance of AuthService with the mock repository
	userService := NewAuthService(mockUserRepository)

	t.Run("Create User Successfully", func(t *testing.T) {
		// Call the CreateUser method
		err := userService.CreateUser("testuser", "password123")

		// Assertions
		assert.NoError(t, err)

	})

	t.Run("Failed to Create User", func(t *testing.T) {
		// Call the CreateUser method
		err := userService.CreateUser("testuser", "password123")

		// Assertions
		assert.Error(t, err)
	})

	t.Run("Should return error if password longer than 72 characters", func(t *testing.T) {
		// Call the CreateUser method
		err := userService.CreateUser("testuser", "arandompasswordthatislongerthan72characterslongarandompasswordthatislongerthan72characterslongarandompasswordthatislongerthan72characterslong")

		// Assertions
		assert.Error(t, err)
	})

}

func TestLoginUser(t *testing.T) {
	mockUserRepository := mocks.NewMockUserRepository()

	// Create a new instance of AuthService with the mock repository
	userService := NewAuthService(mockUserRepository)

	user := &user_model.User{Username: "test", Password: "password123"}

	t.Run("Login User Successfully", func(t *testing.T) {
		err := userService.CreateUser(user.Username, user.Password)
		assert.NoError(t, err)

		// Should log in successfully
		err = userService.LoginUser(user.Username, user.Password)
		assert.NoError(t, err)
	})

	t.Run("Failed to Login User", func(t *testing.T) {
		// Should fail to log in
		err := userService.LoginUser(user.Username, "wrongpassword")
		assert.Error(t, err)
	})

	t.Run("User not found", func(t *testing.T) {

		// Should fail to log in
		err := userService.LoginUser("unknown", "password")
		assert.Error(t, err)

	})

}
