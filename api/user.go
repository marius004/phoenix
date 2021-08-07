package api

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/marius004/phoenix/util"
)

// GetUserByUserName is the handler behind GET /api/users/{username}
func (s *API) GetUserByUserName(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "userName")
	user, err := s.userService.GetByUsername(r.Context(), username)

	if user == nil || err != nil {
		util.ErrorResponse(w, http.StatusNotFound, "User not found", s.logger)
		return
	}

	if util.IsRAdmin(r) {
		util.DataResponse(w, http.StatusOK, user, s.logger)
		return
	}

	if !user.Visible {
		util.ErrorResponse(w, http.StatusBadRequest, "The page of the requested user is not visible", s.logger)
		return
	}

	util.DataResponse(w, http.StatusOK, user, s.logger)
}

// GetUserByUserName is the handler behind GET /api/users/
func (s *API) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.userService.GetAll(r.Context())

	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not fetch users", s.logger)
		return
	}

	util.DataResponse(w, http.StatusOK, users, s.logger)
}

type UserGravatarResponse struct {
	EmailHash string `json:"emailHash"`
	Username  string `json:"username"`
}

// https://en.gravatar.com/site/implement/images/
func (s *API) GetUserGravatar(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))

	if err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "invalid userId url parameter", s.logger)
		return
	}

	user, err := s.userService.GetById(r.Context(), userId)

	if user == nil || err != nil {
		util.ErrorResponse(w, http.StatusNotFound, "User not found", s.logger)
		return
	}

	emailHash := s.calculateEmailHash(user.Email)

	res := &UserGravatarResponse{
		EmailHash: emailHash,
		Username:  user.Username,
	}

	util.DataResponse(w, http.StatusOK, res, s.logger)
}

// https://en.gravatar.com/site/implement/hash/
func (s *API) calculateEmailHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	hash := md5.Sum([]byte(email))
	str := hex.EncodeToString(hash[:])

	return str
}
