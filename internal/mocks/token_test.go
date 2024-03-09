package mocks

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"url-shortener/internal/app/models/user"
)

func TestMockTokenService_GenerateToken(t *testing.T) {
	mockTokenService := NewMockTokenService()

	t.Run("Generate Token Successfully", func(t *testing.T) {
		user := &user_model.User{Username: "testuser"}
		token, err := mockTokenService.GenerateToken(user)

		assert.NoError(t, err)
		assert.Equal(t, "mockToken", token)
	})

	t.Run("Error Generating Token", func(t *testing.T) {
		user := &user_model.User{Username: "error_token"}
		token, err := mockTokenService.GenerateToken(user)

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, user_model.ErrUserAlreadyExists, err)
	})
}

func TestMockTokenService_ValidateToken(t *testing.T) {
	mockTokenService := NewMockTokenService()

	t.Run("Validate Token Successfully", func(t *testing.T) {
		userID, err := mockTokenService.ValidateToken("validToken")

		assert.NoError(t, err)
		assert.Equal(t, uint(123), userID)
	})
}
