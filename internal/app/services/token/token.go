package token_service

import (
	"errors"
	"time"
	"url-shortener/internal/app/models/user"

	"github.com/golang-jwt/jwt"
)

// Service handles JWT token generation and validation.
type Service struct {
	secretKey string
}

type TokenRepository interface {
	GenerateToken(user *user_model.User) (string, error)
	ValidateToken(tokenString string) (uint, error)
}

// NewTokenService creates a new instance of Service.
func NewTokenService(secretKey string) *Service {
	return &Service{secretKey: secretKey}
}

// Claims represents the JWT claims.
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for the provided user.
func (ts *Service) GenerateToken(user *user_model.User) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which include the user ID and expiration time
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create the token with the claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(ts.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates the provided JWT token and extracts the user ID.
func (ts *Service) ValidateToken(tokenString string) (uint, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ts.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return 0, errors.New("invalid token signature")
		}
		return 0, err
	}

	// Check if token is valid
	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Extract user ID from claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return claims.UserID, nil
}
