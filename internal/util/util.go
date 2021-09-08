package util

import (
	"math/rand"
	"strings"
)

const randomCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(size int) string {
	sb := strings.Builder{}
	sb.Grow(size)

	for ; size > 0; size-- {
		randIndex := rand.Intn(len(randomCharacters))
		sb.WriteByte(randomCharacters[randIndex])
	}

	return sb.String()
}
