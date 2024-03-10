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

	userHandler := InitializeUserHandlers(db)

	if err != nil {
		t.Errorf("Error initializing auth handlers: %s", err)
	}

	if userHandler == nil {
		t.Errorf("User handler is nil")
	}

	mock.ExpectClose()
}

func TestInitializeURLHandlers(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	urlHandler := InitializeURLHandlers(db)

	if err != nil {
		t.Errorf("Error initializing URL handlers: %s", err)
	}

	if urlHandler == nil {
		t.Errorf("URL handler is nil")
	}

	mock.ExpectClose()
}

func TestInitializeClickHandlers(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	clickHandler := InitializeClickHandlers(db)

	if err != nil {
		t.Errorf("Error initializing click handlers: %s", err)
	}

	if clickHandler == nil {
		t.Errorf("Click handler is nil")
	}

	mock.ExpectClose()
}
