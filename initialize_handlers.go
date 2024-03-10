package main

import (
	"database/sql"
	"url-shortener/internal/app/handlers"
	"url-shortener/internal/app/handlers/auth"
	"url-shortener/internal/app/handlers/clicks"
	"url-shortener/internal/app/handlers/url"
)

func initializeHandlers(db *sql.DB) (*auth_handler.Handler, *url_handler.Handler, *clicks_handler.Handler) {
	userHandler := handlers.InitializeUserHandlers(db)
	urlHandler := handlers.InitializeURLHandlers(db)
	clicksHandler := handlers.InitializeClickHandlers(db)

	return userHandler, urlHandler, clicksHandler
}
