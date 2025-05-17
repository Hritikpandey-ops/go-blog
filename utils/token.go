package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/segmentio/ksuid"
)

// GenerateAuthToken generates a secure token prefixed with the user's ID.
func GenerateAuthToken(userID ksuid.KSUID) (string, error) {
	const tokenSize = 32 // 32 bytes = 256 bits

	bytes := make([]byte, tokenSize)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate auth token: %w", err)
	}

	randomPart := hex.EncodeToString(bytes)

	// Token format: userID.randomHex
	token := fmt.Sprintf("%s.%s", userID.String(), randomPart)
	return token, nil
}

// ExtractUserIDFromToken extracts the userID (KSUID) from a token of format: userID.randomHex
func ExtractUserIDFromToken(token string) (ksuid.KSUID, error) {
	parts := strings.SplitN(token, ".", 2)
	if len(parts) != 2 {
		return ksuid.Nil, fmt.Errorf("invalid token format: expected 'userID.randomHex'")
	}

	userID, err := ksuid.Parse(parts[0])
	if err != nil {
		return ksuid.Nil, fmt.Errorf("invalid user ID in token: %w", err)
	}

	return userID, nil
}
