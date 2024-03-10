package url_repository

import (
	"database/sql"
	_ "database/sql"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
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
