package api

import (
	"encoding/json"
	"net/http"

	"github.com/marius004/phoenix/models"
	"github.com/marius004/phoenix/util"
)

func (s *API) GetBlogPostById(w http.ResponseWriter, r *http.Request) {
	blogPost := util.BlogPostFromRequestContext(r)
	util.DataResponse(w, http.StatusOK, blogPost, s.logger)
}

func (s *API) CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	var data models.CreateBlogPost

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

	author := util.UserFromRequestContext(r)
	post := models.NewBlogPost(&data, author.Id)

	if err := s.blogPostService.Create(r.Context(), post); err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "could not create blog post", s.logger)
		return
	}

	util.EmptyResponse(w, http.StatusOK)
}

func (s *API) DeleteBlogPostById(w http.ResponseWriter, r *http.Request) {
	blogPost := util.BlogPostFromRequestContext(r)

	if err := s.blogPostService.DeleteById(r.Context(), int(blogPost.Id)); err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "could not delete blog post", s.logger)
		return
	}

	util.EmptyResponse(w, http.StatusOK)
}

func (s *API) UpdateBlogPostById(w http.ResponseWriter, r *http.Request) {
	var data models.UpdateBlogPost

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

	blogPost := util.BlogPostFromRequestContext(r)

	if err := s.blogPostService.UpdateById(r.Context(), int(blogPost.Id), &data); err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "could not update blog post", s.logger)
		return
	}

	util.EmptyResponse(w, http.StatusOK)
}
