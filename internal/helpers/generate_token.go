package helpers

import (
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"time"
	"url-shortener/internal/app/models"
)

// GenerateToken generates a JWT token for the provided auth.
func GenerateToken(user *models.User) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which include the auth ID and expiration time
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   strconv.FormatUint(uint64(user.ID), 10),
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	// Create the token with the claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
