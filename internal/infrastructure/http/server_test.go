package http

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"url-shortener/internal/app/handlers/auth"
	url_handler "url-shortener/internal/app/handlers/url"
	"url-shortener/internal/app/services/auth"
	"url-shortener/internal/app/services/token"
	url_service "url-shortener/internal/app/services/url"
	"url-shortener/internal/mocks"
)

func TestServer_StartAndShutdown(t *testing.T) {
	// Mock auth handler
	mockUserRepository := mocks.NewMockUserRepository()

	// Create a new instance of AuthService with the mock repository
	userService := auth_service.NewAuthService(mockUserRepository)

	tokenService := token_service.NewTokenService(os.Getenv("JWT_SECRET_KEY"))
	mockUserHandler := auth_handler.NewAuthHandler(userService, tokenService)

	urlService := url_service.NewURLService(mocks.NewMockUrlRepository())
	urlHandler := url_handler.NewURLHandler(urlService, tokenService)

	// Create a new HTTP server
	server := NewServer("localhost", "8080", mockUserHandler, urlHandler)

	// Start the server
	go func() {
		err := server.Start()
		assert.NoError(t, err)
	}()

	// Create a context for graceful shutdown
	ctx := context.Background()

	// Shutdown the server
	err := server.Shutdown(ctx)
	assert.NoError(t, err)
}
