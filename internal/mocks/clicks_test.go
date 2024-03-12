package mocks

import "testing"

func TestCreateClick(t *testing.T) {

	mockRepository := NewMockClicksRepository()

	t.Run("Create Click Successfully", func(t *testing.T) {
		err := mockRepository.CreateClick("test-url", "127.0.0.1")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Failed to Create Click", func(t *testing.T) {
		err := mockRepository.CreateClick("invalid", "127.0.0.1")

		if err == nil {

			t.Errorf("Expected an error, got nil")
		}
	})
}

func TestGetClicks(t *testing.T) {
	mockRepository := NewMockClicksRepository()

	t.Run("Get Clicks Successfully", func(t *testing.T) {
		clicks, err := mockRepository.GetClicks("test-url")

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(clicks) != 0 {
			t.Errorf("Expected no clicks, got %v", len(clicks))
		}
	})

	t.Run("Failed to Get Clicks", func(t *testing.T) {
		clicks, err := mockRepository.GetClicks("not_valid")

		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
		if clicks != nil {
			t.Errorf("Expected no clicks, got %v", clicks)
		}
	})
}
