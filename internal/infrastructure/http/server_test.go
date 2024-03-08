package http

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"url-shortener/internal/app/handlers/auth"
	"url-shortener/internal/app/services/auth"
	"url-shortener/internal/mocks"
)

func TestServer_StartAndShutdown(t *testing.T) {
	// Mock auth handler
	mockUserRepository := mocks.NewMockUserRepository()

	// Create a new instance of AuthService with the mock repository
	userService := auth_service.NewAuthService(mockUserRepository)

	mockUserHandler := auth_handler.NewAuthHandler(userService)

	// Create a new HTTP server
	server := NewServer("localhost", "8080", mockUserHandler)

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
