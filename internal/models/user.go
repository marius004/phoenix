package models

import "time"

type User struct {
	Id        uint64     `json:"id"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`

	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Bio      string `json:"bio"`
	Visible  bool   `json:"visible"`

	IsAdmin    bool `json:"isAdmin" db:"is_admin"`
	IsProposer bool `json:"isProposer" db:"is_proposer"`
	IsBanned   bool `json:"isBanned" db:"is_banned"`

	VerifiedEmail           bool       `json:"verifiedEmail" db:"verified_email"`
	EmailVerificationSentAt *time.Time `json:"emailVerificationSentAt" db:"email_verification_sent_at"`
}

func NewUser(username, email, password string) *User {
	usr := &User{
		Username: username,
		Email:    email,
		Password: password,
		IsAdmin:  true,
	}

	return usr
}
