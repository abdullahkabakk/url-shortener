package clicks_model

import (
	"errors"
	"time"
)

var ErrClicksNotFound = errors.New("clicks not found")

// Clicks represents a Clicks entity in the application.
type Clicks struct {
	ID        uint      `json:"id"`
	UrlID     string    `json:"url_id"`
	IPAddress string    `json:"ip_address"`
	CreatedAt time.Time `json:"created_at"`
}
