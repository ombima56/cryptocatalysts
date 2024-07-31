package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateRandomHash(length int) (string, error) {
	// Ensure that length is a valid value for hexadecimal characters
	if length%2 != 0 {
		return "", fmt.Errorf("length must be even to produce a valid hexadecimal string")
	}

	// Create a byte slice to hold the random data
	bytes := make([]byte, length/2)

	// Generate random bytes
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("error generating random bytes: %v", err)
	}

	// Encode bytes to hexadecimal string
	hash := hex.EncodeToString(bytes)
	return hash, nil
}
