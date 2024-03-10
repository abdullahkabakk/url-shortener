package url_service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"url-shortener/internal/mocks"
)

func TestShortenURL(t *testing.T) {
	mockRepo := mocks.NewMockUrlRepository()

	urlService := NewURLService(mockRepo)

	t.Run("Shorten URL Successfully", func(t *testing.T) {
		url, err := urlService.ShortenURL("https://www.example.com", nil)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.NotEmpty(t, url)
	})

	t.Run("Should return error for invalid URL", func(t *testing.T) {
		_, err := urlService.ShortenURL("http://error.com", nil)
		assert.Error(t, err)
	})

}

func TestGetOriginalURL(t *testing.T) {
	mockRepo := mocks.NewMockUrlRepository()

	urlService := NewURLService(mockRepo)

	t.Run("Get Original URL Successfully", func(t *testing.T) {
		url, err := urlService.GetOriginalURL("success")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.NotEmpty(t, url)
	})

	t.Run("Should return error for invalid URL", func(t *testing.T) {
		_, err := urlService.GetOriginalURL("error")
		assert.Error(t, err)
	})

	t.Run("Should return error for nonexistent URL", func(t *testing.T) {
		_, err := urlService.GetOriginalURL("nonexistent")
		assert.Error(t, err)
	})
}

func TestGetUserUrls(t *testing.T) {
	mockRepo := mocks.NewMockUrlRepository()

	urlService := NewURLService(mockRepo)

	t.Run("Get User URLs Successfully", func(t *testing.T) {
		user := uint(1)
		_, err := mockRepo.CreateURL("https://www.example.com", "abc123", &user)
		if err != nil {
			return
		}
		urls, err := urlService.GetUserURLs(1)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		fmt.Println(urls)
		assert.NotEmpty(t, urls)
	})

	t.Run("Should return error for nonexistent user", func(t *testing.T) {
		_, err := urlService.GetUserURLs(2)
		assert.Error(t, err)
	})

}
