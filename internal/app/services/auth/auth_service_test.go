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

	userData := user_model.User{
		Username: "testuser",
		Password: "password123",
	}

	t.Run("Create User Successfully", func(t *testing.T) {
		// Call the CreateUser method
		user, err := userService.CreateUser(userData)

		// Assertions
		assert.NoError(t, err)

		// Check if the user is created
		assert.Equal(t, userData.Username, user.Username)
	})

	t.Run("Failed to Create User", func(t *testing.T) {
		// Call the CreateUser method
		user, err := userService.CreateUser(userData)

		// Assertions
		assert.Error(t, err)

		// Check if the user is not created
		assert.Nil(t, user)
	})

	t.Run("Should return error if password longer than 72 characters", func(t *testing.T) {
		// Call the CreateUser method
		longUserData := user_model.User{
			Username: "testuser",
			Password: "arandompasswordthatislongerthan72characterslongarandompasswordthatislongerthan72characterslongarandompasswordthatislongerthan72characterslong",
		}
		user, err := userService.CreateUser(longUserData)

		// Assertions
		assert.Error(t, err)

		// Check if the user is not created
		assert.Nil(t, user)
	})

}

func TestLoginUser(t *testing.T) {
	mockUserRepository := mocks.NewMockUserRepository()

	// Create a new instance of AuthService with the mock repository
	userService := NewAuthService(mockUserRepository)

	user := user_model.User{Username: "test", Password: "password123"}

	t.Run("Login User Successfully", func(t *testing.T) {
		userVal, err := userService.CreateUser(user)
		assert.NoError(t, err)

		assert.Equal(t, user.Username, userVal.Username)

		// Should log in successfully
		returnVal, err := userService.LoginUser(user)

		assert.NoError(t, err)

		assert.Equal(t, user.Username, returnVal.Username)
	})

	t.Run("Failed to Login User", func(t *testing.T) {
		// Should fail to log in
		user.Password = "wrongpassword"
		userVal, err := userService.LoginUser(user)
		assert.Error(t, err)

		assert.Nil(t, userVal)
	})

	t.Run("User not found", func(t *testing.T) {

		notFoundUser := user_model.User{Username: "unknown", Password: "password123"}
		// Should fail to log in
		userReturn, err := userService.LoginUser(notFoundUser)
		assert.Error(t, err)

		assert.Nil(t, userReturn)

	})

}
