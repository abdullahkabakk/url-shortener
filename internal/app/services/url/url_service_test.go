package url_service

import (
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
		_, err := urlService.ShortenURL("error", nil)
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
