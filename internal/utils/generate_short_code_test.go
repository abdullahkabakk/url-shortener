package utils

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestGenerateShortCode(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Run("Generate Short Code with Length 8", func(t *testing.T) {
		length := 8
		shortCode := GenerateShortCode(length)

		assert.Equal(t, length, len(shortCode))
		for _, char := range shortCode {
			assert.Contains(t, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", string(char))
		}
	})

}
