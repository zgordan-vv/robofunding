package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func NewId() string {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
