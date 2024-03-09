package http

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"testing"
	"url-shortener/internal/app/handlers/auth"
	clicks_handler "url-shortener/internal/app/handlers/clicks"
	url_handler "url-shortener/internal/app/handlers/url"
	"url-shortener/internal/app/services/auth"
	clicks_service "url-shortener/internal/app/services/clicks"
	"url-shortener/internal/app/services/token"
	url_service "url-shortener/internal/app/services/url"
	"url-shortener/internal/mocks"
)

// TestServer_StartAndShutdown tests the functionality of starting and shutting down the HTTP server.
func TestServer_StartAndShutdown(t *testing.T) {
	// Setup
	authService := auth_service.NewAuthService(mocks.NewMockUserRepository())
	urlService := url_service.NewURLService(mocks.NewMockUrlRepository())
	tokenService := token_service.NewTokenService(os.Getenv("JWT_SECRET_KEY"))
	clicksService := clicks_service.NewClicksService(mocks.NewMockClicksRepository())
	userHandler := auth_handler.NewAuthHandler(authService, tokenService) // assuming NewHandler() creates a new instance
	urlHandler := url_handler.NewURLHandler(urlService, tokenService)     // assuming NewHandler() creates a new instance
	clicksHandler := clicks_handler.NewClickHandler(clicksService, urlService)
	server := NewServer("localhost", "8080", userHandler, urlHandler, clicksHandler)

	// Start server
	go func() {
		err := server.Start()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Errorf("server error: %v", err)
			return
		}
	}()
	defer func() {
		err := server.Shutdown(context.Background())
		if err != nil {
			t.Fatalf("shutdown error: %v", err)
		}
	}()

	// Test HTTP request to check server status
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("failed to close response body: %v", err)
		}
	}(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
