package clicks_service

import (
	"testing"
	"url-shortener/internal/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateClick(t *testing.T) {
	mockRepository := mocks.NewMockClicksRepository()

	// Create a new instance of ClicksService with the mock repository
	clickService := NewClicksService(mockRepository)

	t.Run("Create Click Successfully", func(t *testing.T) {
		// Call the CreateClick method
		err := clickService.CreateClick("test-url", "127.0.0.1")

		// Assertions
		assert.NoError(t, err)
	})

	t.Run("Failed to Create Click", func(t *testing.T) {
		// Set up repository to return an error

		// Call the CreateClick method
		err := clickService.CreateClick("invalid", "127.0.0.1")

		// Assertions
		assert.Error(t, err)
	})

}

func TestGetClicks(t *testing.T) {
	mockRepository := mocks.NewMockClicksRepository()

	// Create a new instance of ClicksService with the mock repository
	clickService := NewClicksService(mockRepository)

	t.Run("Get Clicks Successfully", func(t *testing.T) {
		// Call the GetClicks method
		clicks, err := clickService.GetClicks("test-url")

		// Assertions
		assert.NoError(t, err)
		assert.Empty(t, clicks)
	})

	t.Run("Failed to Get Clicks", func(t *testing.T) {
		// Set up repository to return an error

		// Call the GetClicks method
		clicks, err := clickService.GetClicks("not_valid")

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, clicks)
	})
}
