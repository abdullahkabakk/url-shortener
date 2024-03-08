package models

import (
	"errors"
	"time"
)

var ErrUserNotFound = errors.New("auth not found")
var ErrUserAlreadyExists = errors.New("auth already exists")

// User represents a auth entity in the application.
type User struct {
	ID               uint      `json:"id"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	RegistrationDate time.Time `json:"registration_date"`
}
