package url_service

import (
	url_model "url-shortener/internal/app/models/url"
	"url-shortener/internal/app/repositories/url"
	"url-shortener/internal/utils"
)

// Service provides URL-related functionalities.
type Service struct {
	Repository url_repository.Repository
}

// NewURLService creates a new instance of URLService with the given URL repository.
func NewURLService(repository url_repository.Repository) *Service {
	return &Service{Repository: repository}
}

// ShortenURL generates a shortened URL for the given original URL.
func (s *Service) ShortenURL(originalURL string, userID *uint) (string, error) {
	// Generate a unique short code for the URL
	shortCode := utils.GenerateShortCode(8)

	// Save the URL in the repository
	shortenedURL, err := s.Repository.CreateURL(originalURL, shortCode, userID)
	if err != nil {
		return "", err
	}

	return shortenedURL, nil
}

// GetOriginalURL retrieves the original URL corresponding to the given shortened URL.
func (s *Service) GetOriginalURL(shortURL string) (string, error) {
	// Retrieve the original URL from the repository
	originalURL, err := s.Repository.GetOriginalURL(shortURL)
	if err != nil {
		return "", err
	}

	return originalURL, nil
}

// GetUserURLs retrieves the URLs created by the given user.
func (s *Service) GetUserURLs(userID uint) ([]url_model.URL, error) {
	// Retrieve the URLs from the repository
	urls, err := s.Repository.GetUserURLs(userID)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

// GetUserWithShortURL retrieves the user who created the given shortened URL.
func (s *Service) GetUserWithShortURL(userId uint, shortURL string) error {
	// Check if the user is the owner of the short URL
	err := s.Repository.GetUserWithShortURL(userId, shortURL)
	if err != nil {
		return err
	}
	return nil

}
