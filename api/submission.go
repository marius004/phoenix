package api

import (
	"encoding/json"
	"net/http"

	"github.com/marius004/phoenix/models"
	"github.com/marius004/phoenix/util"
)

// GetSubmissions is the handler behind GET /api/submissions/
// URL parameters could be sent for data filtering
// ex. GET /api/submissions?lang=go&score=100
// params available: userId(int), problemId(int), score(int),
// lang(string), status(string), compileError(boolean)
func (s *API) GetSubmissions(w http.ResponseWriter, r *http.Request) {
	filter := models.ParseSubmissionFilter(r)
	submissions, err := s.submissionService.GetByFilter(r.Context(), filter)

	if err != nil {
		s.logger.Println(err)
		util.EmptyResponse(w, http.StatusBadRequest)
		return
	}

	util.DataResponse(w, http.StatusOK, submissions, s.logger)
}

// GetSubmissions is the handler behind GET /api/submissions/{submissionId}
func (s *API) GetSubmissionById(w http.ResponseWriter, r *http.Request) {
	submission := util.SubmissionFromRequestContext(r)
	util.DataResponse(w, http.StatusOK, submission, s.logger)
}

// CreateSubmission is the handler behind POST /api/submissions/
func (s *API) CreateSubmission(w http.ResponseWriter, r *http.Request) {
	user := util.UserFromRequestContext(r)
	var data models.CreateSubmissionRequest

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

	submission := models.NewSubmission(data, int(user.Id))
	if err := s.submissionService.Create(r.Context(), submission); err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not create file-managers", s.logger)
		return
	}

	util.EmptyResponse(w, http.StatusOK)
}
