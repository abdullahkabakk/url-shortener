package clicks_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"url-shortener/internal/app/services/clicks"
	"url-shortener/internal/app/services/url"
)

// Handler handles HTTP requests related to clicks.
type Handler struct {
	// Service is the click service instance.
	Service    *clicks_service.Service
	UrlService *url_service.Service
}

// NewClickHandler creates a new instance of ClickHandler with the given click service.
func NewClickHandler(service *clicks_service.Service, urlService *url_service.Service) *Handler {
	return &Handler{Service: service, UrlService: urlService}
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
