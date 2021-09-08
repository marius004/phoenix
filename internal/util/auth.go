package util

import (
	"errors"
	"fmt"

	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/marius004/phoenix/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CompareHashAndPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false
	}

	if err != nil {
		return false
	}

	return true
}

func GenerateJwtToken(jwtSecret string, expires time.Duration, user *models.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24 * expires).Unix(),
	})

	token, err := claims.SignedString([]byte(jwtSecret))

	return token, err
}

func VerifyToken(tokenString, jwtSecret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// IsTokenValid returns nil if the jwt token is valid and an error otherwise.
func IsTokenValid(tokenString, jwtSecret string) error {
	_, err := VerifyToken(tokenString, jwtSecret)

	if err != nil {
		return err
	}

	return nil
}
