package url_repository

import (
	_ "database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
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

	t.Run("Error no rows", func(t *testing.T) {
		shortCode := "abc123"

		mock.ExpectQuery("SELECT original_url FROM urls").
			WithArgs(shortCode).
			WillReturnError(errors.New("no rows"))

		url, err := repo.GetOriginalURL(shortCode)

		assert.Error(t, err)

		assert.Empty(t, url)

	})
}
