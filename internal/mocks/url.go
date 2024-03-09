package mocks

import (
	"url-shortener/internal/app/models/url"
)

// MockUrlRepository is a mock implementation of UrlRepository interface for testing purposes.
type MockUrlRepository struct {
	Urls map[uint]*url_model.URL
}

// NewMockUrlRepository creates a new instance of MockUrlRepository.
func NewMockUrlRepository() *MockUrlRepository {
	return &MockUrlRepository{
		Urls: make(map[uint]*url_model.URL),
	}
}

// CreateURL simulates creating a new url in the mock database.
func (r *MockUrlRepository) CreateURL(originalUrl, shortCode string, userId *uint) (string, error) {
	if originalUrl == "error" {
		return "", url_model.ErrURLNotFound
	}

	// shortCode should be unique
	for _, u := range r.Urls {
		if u.ShortenedURL == shortCode {
			return "", url_model.ErrShortCodeAlreadyExists
		}
	}

	url := &url_model.URL{
		OriginalURL:  originalUrl,
		ShortenedURL: shortCode,
	}
	r.Urls[uint(len(r.Urls)+1)] = url
	return shortCode, nil
}

// GetOriginalURL simulates retrieving an url by shortCode from the mock database.
func (r *MockUrlRepository) GetOriginalURL(shortCode string) (string, error) {
	// DB error can be simulated here
	if shortCode == "error" {
		return "", url_model.ErrURLNotFound
	}

	if shortCode == "success" {
		return "https://www.google.com", nil
	}

	if shortCode == "invalid" {
		return "https://www.google.com", nil
	}
	// Return an error if url not found
	return "", url_model.ErrURLNotFound
}
