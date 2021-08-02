package api

import (
	"net/http"

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
