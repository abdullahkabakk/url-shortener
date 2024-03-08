package helpers

import (
	"testing"
	"url-shortener/internal/app/models"
)

func TestGenerateToken(t *testing.T) {

	user := &models.User{Username: "testuser", Password: "password123"}

	t.Run("Should generate a token", func(t *testing.T) {
		token, err := GenerateToken(user)
		if err != nil {
			t.Error("Error generating token")
		}
		if token == "" {
			t.Error("Token is empty")
		}
	})
}
