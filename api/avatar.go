package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marius004/phoenix/internal/models"
	"github.com/marius004/phoenix/internal/util"
)

const Url = "https://www.gravatar.com/avatar/"

type avatarFilter struct {
	size int
}

func (s *API) GetUserGravatar(w http.ResponseWriter, r *http.Request) {
	filter := parseAvatarFilter(r)
	user, err := s.getUser(r)

	if err != nil || user == nil {
		util.EmptyResponse(w, http.StatusNotFound)
		return
	}

	emailHash := s.calculateEmailHash(user.Email)
	resp, err := http.Get(Url + fmt.Sprintf("%s?s=%d&r=pg&d=retro", emailHash, filter.size))

	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, 500, "Something went Wrong", s.logger)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, 500, "Something went Wrong", s.logger)
		return
	}

	util.DataResponse(w, http.StatusOK, struct {
		Image    []byte `json:"image"`
		Username string `json:"username"`
	}{body, user.Username}, s.logger)
}

func parseAvatarFilter(r *http.Request) *avatarFilter {
	ret := avatarFilter{}

	if v, ok := r.URL.Query()["size"]; ok {
		size, err := strconv.Atoi(v[0])

		if err == nil {
			ret.size = size
		}
	} else {
		ret.size = 25 // the default size
	}

	return &ret
}

func (s *API) getUser(r *http.Request) (*models.User, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetById(r.Context(), id)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return user, nil
}
