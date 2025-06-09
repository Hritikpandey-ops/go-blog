package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GenerateAuthToken(userID uuid.UUID) (string, error) {
	const tokenSize = 32 // 32 bytes = 256 bits

	bytes := make([]byte, tokenSize)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate auth token: %w", err)
	}

	randomPart := hex.EncodeToString(bytes)

	token := fmt.Sprintf("%s.%s", userID.String(), randomPart)
	return token, nil
}

func ExtractUserIDFromToken(token string) (uuid.UUID, error) {
	parts := strings.SplitN(token, ".", 2)
	if len(parts) != 2 {
		return uuid.Nil, fmt.Errorf("invalid token format: expected 'userID.randomHex'")
	}

	userID, err := uuid.Parse(parts[0])
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID in token: %w", err)
	}

	return userID, nil
}
