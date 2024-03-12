package clicks_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"url-shortener/internal/app/services/clicks"
	token_service "url-shortener/internal/app/services/token"
	"url-shortener/internal/app/services/url"
)

// Handler handles HTTP requests related to clicks.
type Handler struct {
	// Service is the click service instance.
	Service      *clicks_service.Service
	UrlService   *url_service.Service
	TokenService token_service.TokenRepository
}

// NewClickHandler creates a new instance of ClickHandler with the given click service.
func NewClickHandler(service *clicks_service.Service, urlService *url_service.Service, tokenService token_service.TokenRepository) *Handler {
	return &Handler{Service: service, UrlService: urlService, TokenService: tokenService}
}

// CreateClickHandler handles HTTP requests to create a new click.
func (h *Handler) CreateClickHandler(c echo.Context) error {
	// Get the shortened URL from the request
	shortURL := c.Param("id")

	// Call the URL service to get the original URL
	originalURL, err := h.UrlService.GetOriginalURL(shortURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Call the click service to create the click
	err = h.Service.CreateClick(shortURL, c.RealIP())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusMovedPermanently, map[string]string{"Location": originalURL})
}

// GetUserClickDetailsHandler handles HTTP requests to get click details for a user.
func (h *Handler) GetUserClickDetailsHandler(c echo.Context) error {
	// Get the shortened URL from the request
	shortURL := c.Param("id")

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

	err = h.UrlService.GetUserWithShortURL(userID, shortURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Call the click service to get click details for the user
	clickDetails, err := h.Service.GetClicks(shortURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, clickDetails)
}
