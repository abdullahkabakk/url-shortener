package token_service

import (
	"testing"
	"time"
	"url-shortener/internal/app/models/user"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

// MockSecretKey represents a mock secret key for testing.
const MockSecretKey = "mockSecretKey"

func TestTokenService_GenerateToken(t *testing.T) {
	// Create a new user for testing
	user := &user_model.User{ID: 123}

	// Create a new instance of Service with the mock secret key
	tokenService := NewTokenService(MockSecretKey)

	// Generate a token
	token, err := tokenService.GenerateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Parse the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(MockSecretKey), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	// Extract claims from the parsed token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)

	// Validate the user ID in the claims
	assert.Equal(t, float64(user.ID), claims["user_id"])
}

func TestTokenService_ValidateToken(t *testing.T) {
	// Define a mock token
	_ = "mockToken"

	// Create a new instance of Service with the mock secret key
	tokenService := NewTokenService(MockSecretKey)

	t.Run("Valid token", func(t *testing.T) {
		// Create a mock token claims
		claims := &Claims{
			UserID:         123,
			StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(24 * time.Hour).Unix()},
		}

		// Generate a signed token string
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, _ := token.SignedString([]byte(MockSecretKey))

		// Validate the token
		userID, err := tokenService.ValidateToken(signedToken)
		assert.NoError(t, err)
		assert.Equal(t, uint(123), userID)
	})

	t.Run("Expired token", func(t *testing.T) {
		// Create a mock token claims with expired expiration time
		claims := &Claims{
			UserID:         123,
			StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Unix() - 3600}, // 1 hour ago
		}

		// Generate a signed token string
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, _ := token.SignedString([]byte(MockSecretKey))

		// Validate the token
		_, err := tokenService.ValidateToken(signedToken)
		assert.Error(t, err)
		assert.Equal(t, "token is expired by 1h0m0s", err.Error())
	})

}
