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
		_, err := repo.CreateURL("http://error.com", "error", nil)
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

	t.Run("Success - Invalid URL", func(t *testing.T) {
		originalURL, err := repo.GetOriginalURL("invalid")
		assert.NoError(t, err)
		assert.Equal(t, "https://www.google.com", originalURL)
	})
}

func TestMockUrlRepository_GetUserUrls(t *testing.T) {
	repo := NewMockUrlRepository()
	userID := uint(1)
	_, err := repo.CreateURL("https://www.example.com", "abc123", &userID)
	if err != nil {
		return
	}

	t.Run("Success", func(t *testing.T) {
		urls, err := repo.GetUserURLs(userID)
		assert.NoError(t, err)
		assert.NotEmpty(t, urls)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		urls, err := repo.GetUserURLs(2)
		assert.Error(t, err)
		assert.Empty(t, urls)
	})
}

func TestMockUrlRepository_GetUserWithShortURL(t *testing.T) {
	repo := NewMockUrlRepository()
	userID := uint(1)
	_, err := repo.CreateURL("https://www.example.com", "abc123", &userID)
	if err != nil {
		return
	}

	t.Run("Success", func(t *testing.T) {
		err := repo.GetUserWithShortURL(userID, "abc123")
		assert.NoError(t, err)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		err := repo.GetUserWithShortURL(2, "abc123")
		assert.Nil(t, err)
	})
	t.Run("Error - Short URL not found", func(t *testing.T) {
		err := repo.GetUserWithShortURL(userID, "invalid")
		assert.Error(t, err)
	})
}
