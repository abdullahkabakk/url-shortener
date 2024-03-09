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
