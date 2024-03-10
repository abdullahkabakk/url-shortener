package auth_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"url-shortener/internal/app/models/user"
	"url-shortener/internal/app/services/auth"
	"url-shortener/internal/app/services/token"
)

// Handler handles HTTP requests related to users.
type Handler struct {
	// Service is the auth service instance.
	Service         *auth_service.Service
	TokenRepository token_service.TokenRepository
}

// NewAuthHandler creates a new instance of UserHandler with the given auth service.
func NewAuthHandler(service *auth_service.Service, tokenRepository token_service.TokenRepository) *Handler {
	return &Handler{Service: service, TokenRepository: tokenRepository}
}

// CreateUserHandler handles HTTP requests to create a new auth.
func (h *Handler) CreateUserHandler(c echo.Context) error {
	// Parse request body to extract auth data
	var user user_model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Call the auth service to create the auth
	userVal, err := h.Service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Generate a token for the created auth
	token, err := h.TokenRepository.GenerateToken(userVal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"token": token})
}

// LoginUserHandler handles HTTP requests for auth login.
func (h *Handler) LoginUserHandler(c echo.Context) error {
	// Parse request body to extract auth data
	var user user_model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Call the auth service to log in the auth
	userVal, err := h.Service.LoginUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Generate a token for the authenticated auth
	token, err := h.TokenRepository.GenerateToken(userVal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// RefreshTokenHandler handles HTTP requests to refresh an auth token.
func (h *Handler) RefreshTokenHandler(c echo.Context) error {
	// Extract token from request headers or cookies
	token := c.Request().Header.Get("Authorization")

	// Initialize userID to nil
	var userID *uint
	// If token is provided, validate it and get the user ID
	if token != "" {
		parts := strings.Fields(token)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}
		token = parts[1]
		// Call the authentication service to validate the token and get the user ID
		id, err := h.TokenRepository.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}
		userID = &id
	}

	// Generate a new token for the authenticated auth
	token, err := h.TokenRepository.GenerateToken(&user_model.User{ID: *userID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
