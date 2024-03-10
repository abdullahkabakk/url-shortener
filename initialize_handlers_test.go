package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestInitializeHandlers(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		_, _, _ = initializeHandlers(db)

		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})
}
