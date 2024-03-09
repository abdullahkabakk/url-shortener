package clicks_repository

import (
	_ "database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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
