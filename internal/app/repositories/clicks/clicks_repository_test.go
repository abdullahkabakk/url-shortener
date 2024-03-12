package clicks_repository

import (
	_ "database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	clicks_model "url-shortener/internal/app/models/clicks"
)

func TestCreateClick(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBClicksRepository(db)
	shortURL := "test-url"
	ipAddress := "127.0.0.1"

	t.Run("Create Click Successfully", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO clicks").
			ExpectExec().
			WithArgs(shortURL, ipAddress).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.CreateClick(shortURL, ipAddress)

		assert.NoError(t, err)
	})

	t.Run("Failed to Prepare SQL Statement", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO clicks").
			WillReturnError(errors.New("prepare error"))

		err := repo.CreateClick(shortURL, ipAddress)

		assert.Error(t, err)
	})

	t.Run("Failed on SQL Execution", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO clicks").
			ExpectExec().
			WithArgs(shortURL, ipAddress).
			WillReturnError(errors.New("execute error"))

		err := repo.CreateClick(shortURL, ipAddress)

		assert.Error(t, err)
	})

}

func TestGetClicks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBClicksRepository(db)
	shortURL := "test-url"

	//t.Run("Get Clicks Successfully", func(t *testing.T) {
	//	rows := sqlmock.NewRows([]string{"id", "url_id", "ip_address", "created_at"}).
	//		AddRow(1, 1, "127.0.0.1", time.Now()).
	//		AddRow(2, 1, "127.0.0.1", time.Now())
	//
	//	mock.ExpectQuery("SELECT * FROM clicks WHERE url_id = ?").
	//		WithArgs(shortURL).
	//		WillReturnRows(rows)
	//
	//	clicks, err := repo.GetClicks(shortURL)
	//
	//	assert.NoError(t, err)
	//
	//	assert.Len(t, clicks, 2)
	//})

	t.Run("Failed to Prepare SQL Statement", func(t *testing.T) {
		mock.ExpectPrepare("SELECT * FROM clicks").
			WillReturnError(errors.New("prepare error"))

		clicks, err := repo.GetClicks(shortURL)

		assert.Error(t, err)
		assert.Nil(t, clicks)

	})

	t.Run("Failed on SQL Execution", func(t *testing.T) {
		mock.ExpectPrepare("SELECT * FROM clicks").
			ExpectQuery().
			WithArgs(shortURL).
			WillReturnError(errors.New("execute error"))

		clicks, err := repo.GetClicks(shortURL)

		assert.Error(t, err)
		assert.Nil(t, clicks)
	})

	t.Run("Failed to Scan Rows", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "url_id", "ip_address", "created_at"}).
			AddRow(1, 1, "abc", "invalid").
			AddRow(2, 1, "acb", "invalid")

		mock.ExpectPrepare("SELECT * FROM clicks").
			ExpectQuery().
			WithArgs(shortURL).
			WillReturnRows(rows)

		clicks, err := repo.GetClicks(shortURL)

		assert.Error(t, err)
		assert.Nil(t, clicks)
	})

}

func TestGetClicksSuccessful(t *testing.T) {
	// Create a new mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %s", err)
	}
	defer db.Close()

	// Create a new DBClicksRepository with the mock DB
	repo := &DBClicksRepository{DB: db}

	// Define test data
	shortURL := "your-shortened-url"
	now := time.Now()

	// Define expected query and result
	expectedRows := sqlmock.NewRows([]string{"id", "url_id", "ip_address", "created_at"}).
		AddRow(1, "url_id_1", "192.168.0.1", now).
		AddRow(2, "url_id_2", "192.168.0.2", now)

	// Expect the query with the short URL
	mock.ExpectPrepare("SELECT \\* FROM clicks WHERE url_id = \\?").
		ExpectQuery().
		WithArgs(shortURL).
		WillReturnRows(expectedRows)

	// Call the method to be tested
	clicks, err := repo.GetClicks(shortURL)
	if err != nil {
		t.Fatalf("error getting clicks: %s", err)
	}

	// Check if the returned clicks match the expected ones
	expectedClicks := []clicks_model.Clicks{
		{ID: 1, UrlID: "url_id_1", IPAddress: "192.168.0.1", CreatedAt: now},
		{ID: 2, UrlID: "url_id_2", IPAddress: "192.168.0.2", CreatedAt: now},
	}
	if len(clicks) != len(expectedClicks) {
		t.Errorf("expected %d clicks, got %d", len(expectedClicks), len(clicks))
	}
	for i, c := range clicks {
		if c != expectedClicks[i] {
			t.Errorf("expected click %+v, got %+v", expectedClicks[i], c)
		}
	}

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetClicksQueryError(t *testing.T) {
	// Create a new mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %s", err)
	}
	defer db.Close()

	// Create a new DBClicksRepository with the mock DB
	repo := &DBClicksRepository{DB: db}

	// Define test data
	shortURL := "your-shortened-url"

	// Expect the query with the short URL
	mock.ExpectPrepare("SELECT \\* FROM clicks WHERE url_id = \\?").
		ExpectQuery().
		WithArgs(shortURL).
		WillReturnError(errors.New("query error"))

	// Call the method to be tested
	clicks, err := repo.GetClicks(shortURL)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if clicks != nil {
		t.Errorf("expected no clicks, got %+v", clicks)
	}

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetClicksScanError(t *testing.T) {
	// Create a new mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %s", err)
	}
	defer db.Close()

	// Create a new DBClicksRepository with the mock DB
	repo := &DBClicksRepository{DB: db}

	// Define test data
	shortURL := "your-shortened-url"

	// Define expected query and result
	expectedRows := sqlmock.NewRows([]string{"id", "url_id", "ip_address", "created_at"}).
		AddRow(1, "url_id_1", "127.0.0.1", time.Now()).
		AddRow(2, "url_id_2", "127.0.0.1", "invalid")

	// Expect the query with the short URL
	mock.ExpectPrepare("SELECT \\* FROM clicks WHERE url_id = \\?").
		ExpectQuery().
		WithArgs(shortURL).
		WillReturnRows(expectedRows)

	// Call the method to be tested
	clicks, err := repo.GetClicks(shortURL)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if clicks != nil {
		t.Errorf("expected no clicks, got %+v", clicks)
	}
}
