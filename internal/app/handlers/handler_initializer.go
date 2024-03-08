package handlers

import (
	"database/sql"
	"url-shortener/internal/app/handlers/auth"
	"url-shortener/internal/app/repositories/auth"
	"url-shortener/internal/app/services/auth"
)

// InitializeUserHandlers initializes all the auth handlers.
func InitializeUserHandlers(db *sql.DB) (*auth_handler.Handler, error) {
	userRepository := auth_repository.NewDBAuthRepository(db)
	userService := auth_service.NewAuthService(userRepository)
	userHandler := auth_handler.NewAuthHandler(userService)
	return userHandler, nil
}
