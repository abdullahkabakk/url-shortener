package mocks

import (
	"github.com/stretchr/testify/assert"
	"testing"
	url_model "url-shortener/internal/app/models/url"
	"url-shortener/internal/app/models/user"
)

func TestMockTokenService_GenerateToken(t *testing.T) {
	mockTokenService := NewMockTokenService()

	t.Run("Generate Token Successfully", func(t *testing.T) {
		user := &user_model.User{ID: 1, Username: "testuser"}
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

	t.Run("Invalid User", func(t *testing.T) {
		user := &user_model.User{ID: 0}
		token, err := mockTokenService.GenerateToken(user)

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, url_model.ErrInvalidToken, err)

	})

	t.Run("User with ID 2", func(t *testing.T) {
		user := &user_model.User{ID: 2}
		token, err := mockTokenService.GenerateToken(user)

		assert.NoError(t, err)
		assert.Equal(t, "2", token)

	})
}

func TestMockTokenService_ValidateToken(t *testing.T) {
	mockTokenService := NewMockTokenService()

	t.Run("Validate Token Successfully", func(t *testing.T) {
		userID, err := mockTokenService.ValidateToken("validToken")

		assert.NoError(t, err)
		assert.Equal(t, uint(123), userID)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		userID, err := mockTokenService.ValidateToken("invalid")

		assert.Error(t, err)
		assert.Equal(t, url_model.ErrInvalidToken, err)
		assert.Zero(t, userID)
	})
	t.Run("Expired Token", func(t *testing.T) {
		userID, err := mockTokenService.ValidateToken("expired")

		assert.NoError(t, err)
		assert.Zero(t, userID)
	})

	t.Run("Valid Token", func(t *testing.T) {
		userID, err := mockTokenService.ValidateToken("mockToken")

		assert.NoError(t, err)
		assert.Equal(t, uint(1), userID)
	})
}
