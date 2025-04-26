package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateAuthToken() (string, error) {
	bytes := make([]byte, 32) // 32 bytes = 64 hex characters
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
