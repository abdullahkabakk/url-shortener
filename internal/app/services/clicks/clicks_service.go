package clicks_service

import (
	clicks_model "url-shortener/internal/app/models/clicks"
	"url-shortener/internal/app/repositories/clicks"
)

// Service provides URL-related functionalities.
type Service struct {
	Repository clicks_repository.Repository
}

// NewClicksService creates a new instance of ClicksService with the given URL repository.
func NewClicksService(repository clicks_repository.Repository) *Service {
	return &Service{Repository: repository}
}

// CreateClick generates a shortened URL for the given original URL.
func (s *Service) CreateClick(shortURL, ipAddress string) error {
	// Save the click in the repository
	err := s.Repository.CreateClick(shortURL, ipAddress)
	if err != nil {
		return err
	}

	return nil
}

// GetClicks retrieves the clicks for the given shortened URL.
func (s *Service) GetClicks(shortURL string) ([]clicks_model.Clicks, error) {
	// Retrieve the clicks from the repository
	clicks, err := s.Repository.GetClicks(shortURL)
	if err != nil {
		return nil, err
	}

	return clicks, nil
}
