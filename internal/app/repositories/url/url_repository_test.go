package url_repository

import (
	"database/sql"
	_ "database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"url-shortener/internal/app/models/url"
)

func TestDBURLRepository_CreateURL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBURLRepository(db)

	t.Run("Create URL Successfully", func(t *testing.T) {
		originalURL := "https://www.example.com"
		shortCode := "abc123"
		userID := uint(1)

		mock.ExpectPrepare("INSERT INTO urls").
			ExpectExec().
			WithArgs(originalURL, shortCode, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		createdShortCode, err := repo.CreateURL(originalURL, shortCode, &userID)

		assert.NoError(t, err)
		assert.Equal(t, shortCode, createdShortCode)
	})

	t.Run("Failed to Prepare SQL Statement", func(t *testing.T) {
		originalURL := "https://www.example.com"
		shortCode := "abc123"
		userID := uint(1)

		mock.ExpectPrepare("INSERT INTO urls").
			WillReturnError(errors.New("prepare error"))

		createdShortCode, err := repo.CreateURL(originalURL, shortCode, &userID)

		assert.Error(t, err)
		assert.Empty(t, createdShortCode)
	})

	t.Run("Create URL for Non-Registered User", func(t *testing.T) {
		originalURL := "https://www.example.com"
		shortCode := "abc123"

		mock.ExpectPrepare("INSERT INTO urls").
			ExpectExec().
			WithArgs(originalURL, shortCode).
			WillReturnResult(sqlmock.NewResult(1, 1))

		createdShortCode, err := repo.CreateURL(originalURL, shortCode, nil)

		assert.NoError(t, err)
		assert.Equal(t, shortCode, createdShortCode)
	})

	t.Run("No Rows Affected", func(t *testing.T) {
		originalURL := "https://www.example.com"
		shortCode := "abc123"
		userID := uint(1)

		mock.ExpectPrepare("INSERT INTO urls").
			ExpectExec().
			WithArgs(originalURL, shortCode, userID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		createdShortCode, err := repo.CreateURL(originalURL, shortCode, &userID)

		assert.Error(t, err)
		assert.Equal(t, errors.New("no rows affected, insertion failed"), err)
		assert.Empty(t, createdShortCode)
	})

	t.Run("Failed to Execute SQL Statement", func(t *testing.T) {
		originalURL := "https://www.example.com"
		shortCode := "abc123"
		userID := uint(1)

		mock.ExpectPrepare("INSERT INTO urls").
			ExpectExec().
			WithArgs(originalURL, shortCode, userID).
			WillReturnError(errors.New("execute error"))

		createdShortCode, err := repo.CreateURL(originalURL, shortCode, &userID)

		assert.Error(t, err)
		assert.Empty(t, createdShortCode)
	})

}

func TestDBURLRepository_GetOriginalURL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBURLRepository(db)

	t.Run("Get Original URL Successfully", func(t *testing.T) {
		shortCode := "abc123"
		originalURL := "https://www.example.com"

		rows := sqlmock.NewRows([]string{"original_url"}).
			AddRow(originalURL)

		mock.ExpectQuery("SELECT original_url FROM urls").
			WithArgs(shortCode).
			WillReturnRows(rows)

		url, err := repo.GetOriginalURL(shortCode)

		assert.NoError(t, err)
		assert.Equal(t, originalURL, url)
	})

	t.Run("Failed to Prepare SQL Statement", func(t *testing.T) {
		shortCode := "abc123"

		mock.ExpectPrepare("SELECT original_url FROM urls").
			WillReturnError(errors.New("prepare error"))

		url, err := repo.GetOriginalURL(shortCode)

		assert.Error(t, err)
		assert.Empty(t, url)
	})

	t.Run("URL Not Found", func(t *testing.T) {
		shortCode := "abc123"

		mock.ExpectQuery("SELECT original_url FROM urls").
			WithArgs(shortCode).
			WillReturnError(errors.New("no rows found"))

		url, err := repo.GetOriginalURL(shortCode)

		assert.Error(t, err)
		assert.Empty(t, url)
	})

	t.Run("Failed to Execute SQL Statement", func(t *testing.T) {
		shortCode := "abc123"

		mock.ExpectQuery("SELECT original_url FROM urls").
			WithArgs(shortCode).
			WillReturnError(errors.New("execute error"))

		url, err := repo.GetOriginalURL(shortCode)

		assert.Error(t, err)
		assert.Empty(t, url)
	})

	t.Run("No Rows Returned", func(t *testing.T) {
		shortCode := "abc123"

		mock.ExpectQuery("SELECT original_url FROM urls").
			WithArgs(shortCode).
			WillReturnRows(sqlmock.NewRows([]string{"original_url"}))

		url, err := repo.GetOriginalURL(shortCode)

		assert.Error(t, err)
		assert.Empty(t, url)
	})

}

func TestDBURLRepository_GetOriginalURL_ErrorNoRows(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %v", err)
	}
	defer db.Close()

	// Create a new instance of DBURLRepository with the mock DB connection
	repo := NewDBURLRepository(db)

	// Prepare the mock query and result
	mock.ExpectQuery("SELECT original_url FROM urls WHERE shortened_url = ?").
		WithArgs("nonexistent").
		WillReturnError(sql.ErrNoRows)

	// Call the GetOriginalURL method with a nonexistent short code
	_, err = repo.GetOriginalURL("nonexistent")

	// Assert that the correct error is returned
	assert.ErrorIs(t, err, url_model.ErrURLNotFound)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}

func TestDBUrlRepositoryRowsAffected(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBURLRepository(db)

	t.Run("No Rows Affected", func(t *testing.T) {
		originalURL := "https://www.example.com"
		shortCode := "abc123"
		userID := uint(1)

		mock.ExpectPrepare("INSERT INTO urls").
			ExpectExec().
			WithArgs(originalURL, shortCode, userID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		createdShortCode, err := repo.CreateURL(originalURL, shortCode, &userID)

		assert.Error(t, err)
		assert.Equal(t, errors.New("no rows affected, insertion failed"), err)
		assert.Empty(t, createdShortCode)
	})
	t.Run("No Rows Affected", func(t *testing.T) {
		originalURL := "https://www.example.com"
		shortCode := "abc123"
		userID := uint(1)

		// Define the query and expected arguments
		query := "INSERT INTO urls"
		args := []driver.Value{originalURL, shortCode, userID}

		mock.ExpectPrepare(query).
			ExpectExec().
			WithArgs(args...).
			WillReturnResult(sqlmock.NewResult(0, 0))

		// Call the function under test
		createdShortCode, err := repo.CreateURL(originalURL, shortCode, &userID)

		// Check if the error matches the expected one
		assert.Error(t, err)
		assert.Equal(t, "no rows affected, insertion failed", err.Error())
		assert.Empty(t, createdShortCode)
	})
}

func TestDBURLRepository_GetUserURLs(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBURLRepository(db)

	t.Run("Failed to Prepare SQL Statement", func(t *testing.T) {
		userID := uint(1)

		mock.ExpectPrepare("SELECT * FROM urls").
			WillReturnError(errors.New("prepare error"))

		urls, err := repo.GetUserURLs(userID)

		assert.Error(t, err)
		assert.Empty(t, urls)
	})

	t.Run("Failed to Execute SQL Statement", func(t *testing.T) {
		userID := uint(1)

		mock.ExpectQuery("SELECT * FROM urls").
			WithArgs(userID).
			WillReturnError(errors.New("execute error"))

		urls, err := repo.GetUserURLs(userID)

		assert.Error(t, err)
		assert.Empty(t, urls)
	})

	t.Run("No Rows Returned", func(t *testing.T) {
		userID := uint(1)

		mock.ExpectQuery("SELECT * FROM urls").
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"original_url", "shortened_url"}))

		urls, err := repo.GetUserURLs(userID)

		assert.Error(t, err)
		assert.Empty(t, urls)
	})

}

func TestReturnSuccessfulUrlReturn(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBURLRepository(db)

	t.Run("Create URL Successfully", func(t *testing.T) {

		// Define the expected SQL query and results
		expectedUserID := uint(1)
		expectedRows := sqlmock.NewRows([]string{"id", "original_url", "shortened_url", "user_id", "created_at"}).
			AddRow(1, "http://example.com", "http://short.com", expectedUserID, time.Now()).
			AddRow(2, "http://example2.com", "http://short2.com", expectedUserID, time.Now())

		// Expect the query with the given user ID
		mock.ExpectQuery("SELECT \\* FROM urls WHERE user_id = \\?").WithArgs(expectedUserID).WillReturnRows(expectedRows)

		// Call the function to be tested
		urls, err := repo.GetUserURLs(expectedUserID)
		if err != nil {
			t.Errorf("error was not expected while fetching user URLs: %s", err)
		}

		// Check if the returned URLs match the expected ones
		expectedURLs := []url_model.URL{
			{ID: 1, OriginalURL: "http://example.com", ShortenedURL: "http://short.com"},
			{ID: 2, OriginalURL: "http://example2.com", ShortenedURL: "http://short2.com"},
		}
		if len(urls) != len(expectedURLs) {
			t.Errorf("expected %d URLs, got %d", len(expectedURLs), len(urls))
		}

		for i, u := range urls {
			if u.ID != expectedURLs[i].ID || u.OriginalURL != expectedURLs[i].OriginalURL || u.ShortenedURL != expectedURLs[i].ShortenedURL {
				t.Errorf("expected URL %d to be %+v, got %+v", i+1, expectedURLs[i], u)
			}
		}

		// Check if all expected calls were made
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Row Scan Error", func(t *testing.T) {
		// Define the expected user ID
		expectedUserID := uint(1)

		// Expect the query with the given user ID
		mockRows := sqlmock.NewRows([]string{"id", "original_url", "shortened_url"}).
			AddRow(1, "http://example.com", "http://short.com").
			AddRow(2, "http://example2.com", "http://short2.com").
			RowError(0, fmt.Errorf("error scanning row"))

		mock.ExpectQuery("SELECT \\* FROM urls WHERE user_id = \\?").WithArgs(expectedUserID).WillReturnRows(mockRows).WillReturnError(fmt.Errorf("error"))

		// Call the function to be tested
		urls, err := repo.GetUserURLs(expectedUserID)

		assert.Error(t, err)

		assert.Empty(t, urls)

	})

}
