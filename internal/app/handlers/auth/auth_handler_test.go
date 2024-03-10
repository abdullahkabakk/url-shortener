package auth_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/app/models/user"
	"url-shortener/internal/app/services/auth"
	"url-shortener/internal/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Define endpoints for user-related operations.
var (
	userEndpoint     = "/auth/"
	registerEndpoint = userEndpoint + "register/"
	loginEndpoint    = userEndpoint + "login/"
)

// TestCreateUserHandler tests the CreateUserHandler method of the user handler.
func TestCreateUserHandler(t *testing.T) {
	// Create mock user repository and service
	userRepository := mocks.NewMockUserRepository()
	userService := auth_service.NewAuthService(userRepository)
	tokenService := mocks.NewMockTokenService()
	userHandler := NewAuthHandler(userService, tokenService)

	// Define test user data
	userData := user_model.User{
		Username: "testuser",
		Password: "password123",
	}

	t.Run("Should create auth", func(t *testing.T) {
		// Prepare a mock echo.Context with valid request body
		jsonData, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, registerEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call CreateUserHandler with valid request body
		err := userHandler.CreateUserHandler(c)

		// Check the response
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "token")
		assert.NoError(t, err)
	})

	t.Run("Should return error for invalid body", func(t *testing.T) {
		// Prepare a mock echo.Context with invalid request body
		req := httptest.NewRequest(http.MethodPost, registerEndpoint, bytes.NewReader([]byte("invalid")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call CreateUserHandler with invalid request body
		err := userHandler.CreateUserHandler(c)

		// Check the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid request body")
		assert.NoError(t, err)
	})

	t.Run("Should return error for creating existing auth", func(t *testing.T) {
		// Prepare a mock echo.Context with valid request body
		jsonData, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, registerEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call CreateUserHandler with valid request body for existing user
		err := userHandler.CreateUserHandler(c)

		// Check the response
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), user_model.ErrUserAlreadyExists.Error())
		assert.NoError(t, err)
	})

	t.Run("Should return error for user named 'error'", func(t *testing.T) {
		// Prepare a mock echo.Context with valid request body
		userData.Username = "error_token"
		jsonData, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, registerEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call CreateUserHandler with valid request body
		err := userHandler.CreateUserHandler(c)

		// Check the response
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), user_model.ErrUserAlreadyExists.Error())
		assert.NoError(t, err)

	})
}

// TestLoginUserHandler tests the login functionality of the user handler.
func TestLoginUserHandler(t *testing.T) {
	// Create mock user repository and service
	userRepository := mocks.NewMockUserRepository()
	userService := auth_service.NewAuthService(userRepository)
	tokenService := mocks.NewMockTokenService()
	userHandler := NewAuthHandler(userService, tokenService)

	// Define test user data
	userData := user_model.User{
		Username: "testuser",
		Password: "password123",
	}

	t.Run("Should login auth", func(t *testing.T) {
		// Create a new auth
		user, err := userService.CreateUser(userData)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		// Prepare a mock echo.Context with valid request body
		jsonData, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, loginEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call LoginUserHandler
		err = userHandler.LoginUserHandler(c)

		// Check the response
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "token")
		assert.NoError(t, err)
	})

	t.Run("Should return error for invalid body", func(t *testing.T) {
		// Prepare a mock echo.Context with invalid request body
		req := httptest.NewRequest(http.MethodPost, loginEndpoint, bytes.NewReader([]byte("invalid")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call LoginUserHandler with invalid request body
		err := userHandler.LoginUserHandler(c)

		// Check the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid request body")
		assert.NoError(t, err)
	})

	t.Run("Should return error for invalid auth", func(t *testing.T) {
		// Prepare a mock echo.Context with valid request body
		userData.Username = "error"
		jsonData, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, loginEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call LoginUserHandler with invalid request body
		err := userHandler.LoginUserHandler(c)

		// Check the response
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), user_model.ErrUserNotFound.Error())
		assert.NoError(t, err)
	})

	t.Run("Should return error for non-existing auth", func(t *testing.T) {
		// Prepare a mock echo.Context with valid request body
		userData.Username = "nonexisting"
		jsonData, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, loginEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call LoginUserHandler with invalid request body
		err := userHandler.LoginUserHandler(c)

		// Check the response
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), user_model.ErrUserNotFound.Error())
		assert.NoError(t, err)
	})

	t.Run("Should return error for user named 'error_token'", func(t *testing.T) {
		// Create user with username 'error_token'
		invalidUser := user_model.User{
			Username: "error_token",
			Password: "password123",
		}

		user, err := userService.CreateUser(invalidUser)
		assert.NoError(t, err)

		assert.NotNil(t, user)

		// Prepare a mock echo.Context with valid request body
		userData.Username = "error_token"
		jsonData, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, loginEndpoint, bytes.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call LoginUserHandler with valid request body
		err = userHandler.LoginUserHandler(c)

		// Check the response
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), user_model.ErrUserAlreadyExists.Error())
		assert.NoError(t, err)

	})
}

func TestRefreshTokenHandler(t *testing.T) {
	// Create mock user repository and service
	userRepository := mocks.NewMockUserRepository()
	userService := auth_service.NewAuthService(userRepository)
	tokenService := mocks.NewMockTokenService()
	userHandler := NewAuthHandler(userService, tokenService)

	// Define test user data
	userData := user_model.User{
		Username: "testuser",
		Password: "password123",
	}

	t.Run("Should refresh token", func(t *testing.T) {
		// Create a new auth
		user, err := userService.CreateUser(userData)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		// Prepare a mock echo.Context with valid request body
		req := httptest.NewRequest(http.MethodPost, userEndpoint+"refresh/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer valid_token")
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call RefreshTokenHandler
		err = userHandler.RefreshTokenHandler(c)

		// Check the response
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "token")
		assert.NoError(t, err)
	})

	t.Run("Should return error for invalid token", func(t *testing.T) {
		// Prepare a mock echo.Context with valid request body
		req := httptest.NewRequest(http.MethodPost, userEndpoint+"refresh/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer invalid")
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call RefreshTokenHandler
		err := userHandler.RefreshTokenHandler(c)

		// Check the response
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), `{"error":"Invalid token"}`)
		assert.NoError(t, err)
	})

	t.Run("Should return error for expired token", func(t *testing.T) {
		// Prepare a mock echo.Context with valid request body
		req := httptest.NewRequest(http.MethodPost, userEndpoint+"refresh/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer expired")
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call RefreshTokenHandler
		err := userHandler.RefreshTokenHandler(c)

		fmt.Println(err)
		fmt.Println(rec.Body.String())
		// Check the response
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), `{"error":"invalid token"}`)
		assert.NoError(t, err)
	})

	t.Run("Should return error for non-existing user", func(t *testing.T) {
		// Prepare a mock echo.Context with valid request body
		req := httptest.NewRequest(http.MethodPost, userEndpoint+"refresh/", nil)
		req.Header.Set(echo.HeaderAuthorization, "NotExisting nonexisting")
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		// Call RefreshTokenHandler
		err := userHandler.RefreshTokenHandler(c)

		// Check the response
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), `{"error":"Invalid token"}`)
		assert.NoError(t, err)
	})

}
