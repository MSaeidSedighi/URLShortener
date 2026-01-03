package utils

import (
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	var sb strings.Builder

	for i := 0; i < length; i++ {
		sb.WriteByte(
			charset[rand.Intn(len(charset))],
		)
	}

	return sb.String()
}
