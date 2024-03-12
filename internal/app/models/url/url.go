package url_model

import (
	"errors"
	"time"
)

var ErrURLNotFound = errors.New("URL not found")
var ErrShortCodeAlreadyExists = errors.New("short code already exists")
var ErrInvalidToken = errors.New("invalid token")
var ErrClickNotCreated = errors.New("click not created")

// URL represents a URL entity in the application.
type URL struct {
	OriginalURL  string    `json:"original_url"`
	ShortenedURL string    `json:"shortened_url"`
	UserID       uint      `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
}
