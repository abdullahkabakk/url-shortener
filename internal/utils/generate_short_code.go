package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateShortCode generates a random short code for URLs.
func GenerateShortCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)

	charsetLength := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charsetLength)
		code[i] = charset[randomIndex.Int64()]
	}
	return string(code)
}
