package handlers

import (
	"database/sql"
	"os"
	"url-shortener/internal/app/handlers/auth"
	url_handler "url-shortener/internal/app/handlers/url"
	"url-shortener/internal/app/repositories/auth"
	url_repository "url-shortener/internal/app/repositories/url"
	"url-shortener/internal/app/services/auth"
	"url-shortener/internal/app/services/token"
	url_service "url-shortener/internal/app/services/url"
)

// InitializeUserHandlers initializes all the auth handlers.
func InitializeUserHandlers(db *sql.DB) (*auth_handler.Handler, error) {
	userRepository := auth_repository.NewDBAuthRepository(db)
	userService := auth_service.NewAuthService(userRepository)
	tokenService := token_service.NewTokenService(os.Getenv("JWT_SECRET_KEY"))
	userHandler := auth_handler.NewAuthHandler(userService, tokenService)
	return userHandler, nil
}

// InitializeURLHandlers initializes all the URL handlers.
func InitializeURLHandlers(db *sql.DB) (*url_handler.Handler, error) {
	urlRepository := url_repository.NewDBURLRepository(db)
	urlService := url_service.NewURLService(urlRepository)
	tokenService := token_service.NewTokenService(os.Getenv("JWT_SECRET_KEY"))
	urlHandler := url_handler.NewURLHandler(urlService, tokenService)
	return urlHandler, nil
}
