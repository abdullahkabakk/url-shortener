package auth_repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"
	"url-shortener/internal/app/models/user"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMySQLUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBAuthRepository(db)
	user := &user_model.User{
		Username: "testuser",
		Password: "password123",
	}

	t.Run("Create User Successfully", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO users").
			ExpectExec().
			WithArgs(user.Username, user.Password).
			WillReturnResult(sqlmock.NewResult(1, 1))

		createdUser, err := repo.Create(user)

		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, uint(1), createdUser.ID)
	})

	t.Run("Failed to Prepare SQL Statement", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO users").
			WillReturnError(errors.New("prepare error"))

		createdUser, err := repo.Create(user)

		assert.Error(t, err)
		assert.Nil(t, createdUser)
	})

	t.Run("Failed on sql statement", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO users").
			ExpectExec().
			WithArgs(user.Username, user.Password).
			WillReturnError(errors.New("execute error"))

		createdUser, err := repo.Create(user)

		assert.Error(t, err)
		assert.Nil(t, createdUser)

	})

	t.Run("Failed to close prepared statement", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO users").
			ExpectExec().
			WithArgs(user.Username, user.Password).
			WillReturnResult(sqlmock.NewResult(1, 1))

		createdUser, err := repo.Create(user)

		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, uint(1), createdUser.ID)

	})

	t.Run("Failed to Execute SQL Statement", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO users").
			ExpectExec().
			WithArgs(user.Username, user.Password).
			WillReturnError(errors.New("execute error"))

		createdUser, err := repo.Create(user)

		assert.Error(t, err)
		assert.Nil(t, createdUser)
	})

	t.Run("Failed to Retrieve Last Inserted ID", func(t *testing.T) {
		mock.ExpectPrepare("INSERT INTO users").
			ExpectExec().
			WithArgs(user.Username, user.Password).
			WillReturnResult(sqlmock.NewErrorResult(errors.New("last insert ID error")))

		createdUser, err := repo.Create(user)

		assert.Error(t, err)
		assert.Nil(t, createdUser)

	})

}

func TestDBUserRepositoryGetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewDBAuthRepository(db)
	username := "testuser"

	t.Run("Get User Successfully", func(t *testing.T) {
		expectedUser := &user_model.User{
			ID:               1,
			Username:         username,
			Password:         "password123",
			RegistrationDate: time.Now(),
		}

		mock.ExpectQuery("SELECT id, username, password, registration_date FROM users").
			WithArgs(username).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "registration_date"}).
				AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Password, expectedUser.RegistrationDate))

		user, err := repo.GetByUsername(username)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)

	})

	t.Run("Failed to Get User (No Rows Returned)", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, username, password, registration_date FROM users").
			WithArgs(username).
			WillReturnError(sql.ErrNoRows)

		user, err := repo.GetByUsername(username)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.True(t, errors.Is(err, user_model.ErrUserNotFound))

	})

	t.Run("Failed to Get User (Scan Error)", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, username, password, registration_date FROM users").
			WithArgs(username).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "registration_date"}).
				AddRow(nil, nil, nil, nil))

		user, err := repo.GetByUsername(username)
		assert.Error(t, err)
		assert.Nil(t, user)

	})

}
