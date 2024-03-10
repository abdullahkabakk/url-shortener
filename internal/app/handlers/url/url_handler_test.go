package url_handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/app/models/url"
	user_model "url-shortener/internal/app/models/user"
	"url-shortener/internal/app/services/url"
	"url-shortener/internal/mocks"
)

var (
	urlEndpoint     = "/url/"
	shortenEndpoint = urlEndpoint + "shorten/"
)

func TestShortenUrlHandler(t *testing.T) {
	mockRepository := mocks.NewMockUrlRepository()
	mockService := url_service.NewURLService(mockRepository)
	tokenService := mocks.NewMockTokenService()
	mockHandler := NewURLHandler(mockService, tokenService)

	t.Run("Should shorten a URL", func(t *testing.T) {
		urlData := url_model.URL{
			OriginalURL: "https://www.example.com",
		}
		jsonData, _ := json.Marshal(urlData)
		req := httptest.NewRequest(http.MethodPost, shortenEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		err := mockHandler.ShortenURLHandler(c)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "shortened_url")
		assert.NoError(t, err)

	})

	t.Run("Should return error for invalid body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, shortenEndpoint, bytes.NewReader([]byte("invalid")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		err := mockHandler.ShortenURLHandler(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		assert.NoError(t, err)
	})

	t.Run("Invalid token header", func(t *testing.T) {
		urlData := url_model.URL{
			OriginalURL: "https://www.example.com",
		}
		jsonData, _ := json.Marshal(urlData)
		req := httptest.NewRequest(http.MethodPost, shortenEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "invalid")
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		err := mockHandler.ShortenURLHandler(c)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		assert.NoError(t, err)

	})

	t.Run("Invalid token", func(t *testing.T) {
		urlData := url_model.URL{
			OriginalURL: "https://www.example.com",
		}
		jsonData, _ := json.Marshal(urlData)
		req := httptest.NewRequest(http.MethodPost, shortenEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("Authorization", "Bearer invalid")
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		err := mockHandler.ShortenURLHandler(c)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		assert.NoError(t, err)
	})

	t.Run("Invalid url", func(t *testing.T) {
		urlData := url_model.URL{
			OriginalURL: "http://error.com",
		}
		jsonData, _ := json.Marshal(urlData)
		req := httptest.NewRequest(http.MethodPost, shortenEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		err := mockHandler.ShortenURLHandler(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		assert.NoError(t, err)

	})

	t.Run("Valid token", func(t *testing.T) {
		urlData := url_model.URL{
			OriginalURL: "https://www.example.com",
		}
		jsonData, _ := json.Marshal(urlData)
		req := httptest.NewRequest(http.MethodPost, shortenEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		token, err := tokenService.GenerateToken(&user_model.User{ID: 1})
		if err != nil {
			return
		}
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		err = mockHandler.ShortenURLHandler(c)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "shortened_url")
		assert.NoError(t, err)
	})

	t.Run("Invalid url", func(t *testing.T) {
		urlData := url_model.URL{
			OriginalURL: "error",
		}
		jsonData, _ := json.Marshal(urlData)
		req := httptest.NewRequest(http.MethodPost, shortenEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		err := mockHandler.ShortenURLHandler(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		assert.NoError(t, err)
	})

}
