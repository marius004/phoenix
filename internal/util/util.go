package util

import (
	"math/rand"
	"strings"

	"github.com/marius004/phoenix/internal/models"
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

func IsAuthed(user *models.User) bool {
	return user != nil
}

func IsAdmin(user *models.User) bool {
	return user != nil && user.IsAdmin
}

func IsProposer(user *models.User) bool {
	return user != nil && (user.IsProposer || user.IsAdmin)
}
