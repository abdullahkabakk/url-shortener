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
	if originalUrl == "http://error.com" {
		return "", url_model.ErrURLNotFound
	}

	// shortCode should be unique
	for _, u := range r.Urls {
		if u.ShortenedURL == shortCode {
			return "", url_model.ErrShortCodeAlreadyExists
		}
	}

	var userID uint
	if userId != nil {
		userID = *userId
	}

	url := &url_model.URL{
		OriginalURL:  originalUrl,
		ShortenedURL: shortCode,
		UserID:       userID,
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

// GetUserURLs simulates retrieving all urls created by a user from the mock database.
func (r *MockUrlRepository) GetUserURLs(userId uint) ([]url_model.URL, error) {
	urls := make([]url_model.URL, 0)
	for _, u := range r.Urls {
		if u.UserID == userId {
			urls = append(urls, *u)
		}
	}
	if len(urls) == 0 {
		return nil, url_model.ErrURLNotFound
	}
	return urls, nil
}
