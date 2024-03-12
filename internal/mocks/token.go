package mocks

import (
	"url-shortener/internal/app/models/url"
	"url-shortener/internal/app/models/user"
)

// MockTokenService is a mock implementation of TokenService for testing purposes.
type MockTokenService struct {
	secretKey string
}

// NewMockTokenService creates a new instance of MockTokenService.
func NewMockTokenService() *MockTokenService {
	return &MockTokenService{
		secretKey: "mockSecret",
	}
}

// GenerateToken mocks the GenerateToken method of TokenService.
func (mts *MockTokenService) GenerateToken(user *user_model.User) (string, error) {
	if user.Username == "error_token" {
		return "", user_model.ErrUserAlreadyExists
	}
	if user.ID == 0 {
		return "", url_model.ErrInvalidToken
	}
	if user.ID == 2 {
		return "2", nil
	}
	// For simplicity in testing, return a fixed token
	return "mockToken", nil
}

// ValidateToken mocks the ValidateToken method of TokenService.
func (mts *MockTokenService) ValidateToken(tokenString string) (uint, error) {
	if tokenString == "invalid" {
		return 0, url_model.ErrInvalidToken
	}
	if tokenString == "expired" {
		return 0, nil
	}
	if tokenString == "mockToken" {
		return 1, nil
	}
	// For simplicity in testing, return a fixed user ID
	return 123, nil
}
