package mocks

import (
	"url-shortener/internal/app/models/url"
)

// MockClicksRepository is a mock implementation of UrlRepository interface for testing purposes.
type MockClicksRepository struct {
	Urls map[uint]*url_model.URL
}

func (m MockClicksRepository) CreateClick(shortURL, ipAddress string) error {
	if shortURL == "invalid" {
		return url_model.ErrClickNotCreated
	}

	return nil
}

// NewMockClicksRepository creates a new instance of MockUrlRepository.
func NewMockClicksRepository() *MockClicksRepository {
	return &MockClicksRepository{
		Urls: make(map[uint]*url_model.URL),
	}
}
