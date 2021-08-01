package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	userNameValidation = []validation.Rule{validation.Required, validation.Length(3, 20), is.PrintableASCII}
	passwordValidation = []validation.Rule{validation.Required, validation.Length(8, 20)}
	emailValidation    = []validation.Rule{validation.Required, is.Email}
)

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (data SignupRequest) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Username, userNameValidation...),
		validation.Field(&data.Email, emailValidation...),
		validation.Field(&data.Password, passwordValidation...),
	)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (data LoginRequest) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Username, userNameValidation...),
		validation.Field(&data.Password, passwordValidation...),
	)
}
