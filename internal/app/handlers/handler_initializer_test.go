package handlers

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestInitializeUserHandlers(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	userHandler, err := InitializeUserHandlers(db)

	if err != nil {
		t.Errorf("Error initializing auth handlers: %s", err)
	}

	if userHandler == nil {
		t.Errorf("User handler is nil")
	}

	mock.ExpectClose()

}
