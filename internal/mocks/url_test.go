package mocks

import (
	"testing"
	"url-shortener/internal/app/models/url"

	"github.com/stretchr/testify/assert"
)

func TestMockUrlRepository_CreateURL(t *testing.T) {
	repo := NewMockUrlRepository()

	t.Run("Success", func(t *testing.T) {
		shortCode, err := repo.CreateURL("https://example.com", "abc123", nil)
		assert.NoError(t, err)
		assert.Equal(t, "abc123", shortCode)
	})

	t.Run("Error - Short code already exists", func(t *testing.T) {
		repo.Urls = map[uint]*url_model.URL{
			1: {ShortenedURL: "existingShortCode"},
		}
		_, err := repo.CreateURL("https://example.com", "existingShortCode", nil)
		assert.Error(t, err)
		assert.Equal(t, url_model.ErrShortCodeAlreadyExists, err)
	})
	t.Run("Error- Should return error if URL is 'error'", func(t *testing.T) {
		_, err := repo.CreateURL("error", "error", nil)
		assert.Error(t, err)
		assert.Equal(t, url_model.ErrURLNotFound, err)

	})
}

func TestMockUrlRepository_GetOriginalURL(t *testing.T) {
	repo := NewMockUrlRepository()

	t.Run("Success", func(t *testing.T) {
		originalURL, err := repo.GetOriginalURL("success")
		assert.NoError(t, err)
		assert.Equal(t, "https://www.google.com", originalURL)
	})

	t.Run("Error - URL not found", func(t *testing.T) {
		_, err := repo.GetOriginalURL("notFound")
		assert.Error(t, err)
		assert.Equal(t, url_model.ErrURLNotFound, err)
	})

	t.Run("Error - Short code not found", func(t *testing.T) {
		_, err := repo.GetOriginalURL("error")
		assert.Error(t, err)
		assert.Equal(t, url_model.ErrURLNotFound, err)
	})
}
