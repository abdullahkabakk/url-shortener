package url_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"strings"
	"url-shortener/internal/app/models/url"
	"url-shortener/internal/app/services/token"
	"url-shortener/internal/app/services/url"
)

// Handler handles HTTP requests related to URLs.
type Handler struct {
	// Service is the URL service instance.
	Service      *url_service.Service
	TokenService token_service.TokenRepository
}

// NewURLHandler creates a new instance of URLHandler with the given URL service.
func NewURLHandler(service *url_service.Service, tokenService token_service.TokenRepository) *Handler {
	return &Handler{Service: service, TokenService: tokenService}
}

// ShortenURLHandler handles HTTP requests to shorten a URL.
func (h *Handler) ShortenURLHandler(c echo.Context) error {
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
		id, err := h.TokenService.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}
		userID = &id
	}

	// Parse request body to extract URL data
	var urlData url_model.URL
	if err := c.Bind(&urlData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if urlData.OriginalURL == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Original URL is required"})
	}

	// Check if the provided URL is valid
	_, err := url.ParseRequestURI(urlData.OriginalURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid URL"})
	}

	// Call the URL service to shorten the URL with the user ID
	shortenedURL, err := h.Service.ShortenURL(urlData.OriginalURL, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"shortened_url": shortenedURL})
}

// GetUserUrlsHandler handles HTTP requests to get the URLs of a user.
func (h *Handler) GetUserUrlsHandler(c echo.Context) error {
	// Extract token from request headers or cookies
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is required"})
	}

	// Initialize userID to nil
	var userID uint
	// If token is provided, validate it and get the user ID
	parts := strings.Fields(token)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}
	token = parts[1]
	// Call the authentication service to validate the token and get the user ID
	id, err := h.TokenService.ValidateToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}
	userID = id

	// Call the URL service to get the URLs of the user
	urls, err := h.Service.GetUserURLs(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, urls)
}
