package handlers

import (
	"database/sql"
	"os"
	"url-shortener/internal/app/handlers/auth"
	clicks_handler "url-shortener/internal/app/handlers/clicks"
	url_handler "url-shortener/internal/app/handlers/url"
	"url-shortener/internal/app/repositories/auth"
	clicks_repository "url-shortener/internal/app/repositories/clicks"
	url_repository "url-shortener/internal/app/repositories/url"
	"url-shortener/internal/app/services/auth"
	clicks_service "url-shortener/internal/app/services/clicks"
	"url-shortener/internal/app/services/token"
	url_service "url-shortener/internal/app/services/url"
)

// InitializeUserHandlers initializes all the auth handlers.
func InitializeUserHandlers(db *sql.DB) *auth_handler.Handler {
	userRepository := auth_repository.NewDBAuthRepository(db)
	userService := auth_service.NewAuthService(userRepository)
	tokenService := token_service.NewTokenService(os.Getenv("JWT_SECRET_KEY"))
	userHandler := auth_handler.NewAuthHandler(userService, tokenService)
	return userHandler
}

// InitializeURLHandlers initializes all the URL handlers.
func InitializeURLHandlers(db *sql.DB) *url_handler.Handler {
	urlRepository := url_repository.NewDBURLRepository(db)
	urlService := url_service.NewURLService(urlRepository)
	tokenService := token_service.NewTokenService(os.Getenv("JWT_SECRET_KEY"))
	urlHandler := url_handler.NewURLHandler(urlService, tokenService)
	return urlHandler
}

// InitializeClickHandlers initializes all the click handlers.
func InitializeClickHandlers(db *sql.DB) *clicks_handler.Handler {
	clickRepository := clicks_repository.NewDBClicksRepository(db)

	urlRepository := url_repository.NewDBURLRepository(db)
	urlService := url_service.NewURLService(urlRepository)

	clickService := clicks_service.NewClicksService(clickRepository)
	clickHandler := clicks_handler.NewClickHandler(clickService, urlService)
	return clickHandler
}
