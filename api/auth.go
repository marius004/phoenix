package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/marius004/phoenix/internal/models"
	"github.com/marius004/phoenix/internal/util"
)

// Signup is the handler behind POST /api/auth/signup.
func (s *API) Signup(w http.ResponseWriter, r *http.Request) {
	var data models.SignupRequest

	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()

	if err := jsonDecoder.Decode(&data); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), s.logger)
		return
	}

	if err := data.Validate(); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), s.logger)
		return
	}

	if usr, err := s.userService.GetByUsername(r.Context(), data.Username); (err != nil && !errors.Is(err, sql.ErrNoRows)) || (usr != nil) {
		util.ErrorResponse(w, http.StatusBadRequest, "username is taken", s.logger)
		return
	}

	if usr, err := s.userService.GetByEmail(r.Context(), data.Email); (err != nil && !errors.Is(err, sql.ErrNoRows)) || (usr != nil) {
		util.ErrorResponse(w, http.StatusBadRequest, "email already exists", s.logger)
		return
	}

	hashedPassword, err := util.GeneratePasswordHash(data.Password)

	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "could not create user", s.logger)
		return
	}

	user := models.NewUser(data.Username, data.Email, hashedPassword)
	err = s.userService.Create(r.Context(), user)

	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "could not create user", s.logger)
		return
	}

	authCookieLifeTime := time.Duration(s.config.Api.AuthCookieLifeTime)
	token, err := util.GenerateJwtToken(s.config.JwtSecret(), authCookieLifeTime, user)

	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "could not create user", s.logger)
		return
	}

	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "could not create user", s.logger)
		return
	}

	s.setAuthCookies(w, token, user)
	util.EmptyResponse(w, http.StatusCreated)
}

func (s *API) deleteCookies(w http.ResponseWriter) {
	httpCookie := &http.Cookie{
		Name:     HttpOnlyCookieAuthKey,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}

	clientCookie := &http.Cookie{
		Name:     CookieAuthKey,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, httpCookie)
	http.SetCookie(w, clientCookie)
}

// The authentication is done with 2 cookies:
// - the first one is an httponly cookie(cannot be modified by the client). Contains the jwt token.
// - the second one is normal cookie that contains information about the user and the jwt token
func (s *API) setAuthCookies(w http.ResponseWriter, token string, user *models.User) {
	authCookieLifeTime := time.Duration(s.config.Api.AuthCookieLifeTime)

	serverAuthCookie := newServerAuthCookie(token).Cookie(authCookieLifeTime)
	clientAuthCookie := newClientAuthCookie(token, user).Cookie(authCookieLifeTime)

	http.SetCookie(w, serverAuthCookie)
	http.SetCookie(w, clientAuthCookie)
}

// Login is the handler behind POST /api/auth/login.
func (s *API) Login(w http.ResponseWriter, r *http.Request) {
	var data models.LoginRequest

	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()

	if err := jsonDecoder.Decode(&data); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), s.logger)
		return
	}

	if err := data.Validate(); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), s.logger)
		return
	}

	user, err := s.userService.GetByUsername(r.Context(), data.Username)

	if user == nil || err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "user not found", s.logger)
		return
	}

	if !util.CompareHashAndPassword(data.Password, user.Password) {
		util.ErrorResponse(w, http.StatusUnauthorized, "Invalid username or password", s.logger)
		return
	}

	authCookieLifeTime := time.Duration(s.config.Api.AuthCookieLifeTime)
	token, err := util.GenerateJwtToken(s.config.JwtSecret(), authCookieLifeTime, user)

	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "could not create user", s.logger)
		return
	}

	s.setAuthCookies(w, token, user)
	util.EmptyResponse(w, http.StatusOK)
}

func (s *API) Logout(w http.ResponseWriter, r *http.Request) {
	s.deleteCookies(w)
	util.EmptyResponse(w, http.StatusOK)
}
