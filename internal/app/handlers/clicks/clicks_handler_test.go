package clicks_handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/app/services/clicks"
	url_service "url-shortener/internal/app/services/url"
	"url-shortener/internal/mocks"
)

func TestCreateClick(t *testing.T) {

	// Create mock user repository and service
	clickRepository := mocks.NewMockClicksRepository()
	clickService := clicks_service.NewClicksService(clickRepository)
	mockRepository := mocks.NewMockUrlRepository()
	mockService := url_service.NewURLService(mockRepository)
	clickHandler := NewClickHandler(clickService, mockService)

	// Define test user data
	t.Run("Should create click", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")

		c.SetParamNames("id")
		c.SetParamValues("success")

		// Mock GetOriginalURL method
		// Call CreateClickHandler
		err := clickHandler.CreateClickHandler(c)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, http.StatusMovedPermanently, rec.Code)
		assert.JSONEq(t, `{"Location":"https://www.google.com"}`, rec.Body.String())
	})

	t.Run("Should return error for error URL", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")

		c.SetParamNames("id")
		c.SetParamValues("error")

		// Mock GetOriginalURL method
		// Call CreateClickHandler
		err := clickHandler.CreateClickHandler(c)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"error":"URL not found"}`, rec.Body.String())
		assert.NoError(t, err)
	})

	t.Run("Should return error for `invalid` URL", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")

		c.SetParamNames("id")
		c.SetParamValues("invalid")

		// Mock GetOriginalURL method
		// Call CreateClickHandler
		err := clickHandler.CreateClickHandler(c)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		assert.NoError(t, err)
	})

}
