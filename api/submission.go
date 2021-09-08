package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/marius004/phoenix/internal/models"
	"github.com/marius004/phoenix/internal/util"
)

// GetSubmissions is the handler behind GET /api/submissions/
// URL parameters could be sent for data filtering
// ex. GET /api/submissions?lang=go&score=100
// params available: userId(int), problemId(int), score(int),
// lang(string), status(string), compileError(boolean)
func (s *API) GetSubmissions(w http.ResponseWriter, r *http.Request) {
	filter := s.parseSubmissionFilter(r)
	submissions, err := s.submissionService.GetByFilter(r.Context(), filter)

	// I know it is not a good practice to do something like this
	// but I hope I will restructure this in the future
	for _, submission := range submissions {
		if problem, err := s.problemService.GetById(r.Context(), submission.ProblemId); err == nil {
			submission.ProblemName = problem.Name
		}

		if user, err := s.userService.GetById(r.Context(), submission.UserId); err == nil {
			submission.EmailHash = s.calculateEmailHash(user.Email)
			submission.Username = user.Username
		}
	}

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

	if problem, err := s.problemService.GetById(r.Context(), submission.ProblemId); err == nil {
		submission.ProblemName = problem.Name
	}

	if user, err := s.userService.GetById(r.Context(), submission.UserId); err == nil {
		submission.EmailHash = s.calculateEmailHash(user.Email)
		submission.Username = user.Username
	}

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

	util.DataResponse(w, http.StatusCreated, submission, s.logger)
}

func (s *API) parseSubmissionFilter(r *http.Request) *models.SubmissionFilter {
	ret := models.SubmissionFilter{}

	if v, ok := r.URL.Query()["username"]; ok {
		last := len(v) - 1
		username := v[last]

		user, err := s.userService.GetByUsername(r.Context(), username)
		if user != nil && err == nil {
			ret.UserId = int(user.Id)
		} else {
			ret.UserId = -1
		}
	}

	if v, ok := r.URL.Query()["lang"]; ok {
		ret.Langs = convertUrlValuesToLangArr(v)
	}

	if v, ok := r.URL.Query()["problem"]; ok {
		last := len(v) - 1
		problemName := v[last]

		problem, err := s.problemService.GetByName(r.Context(), problemName)
		if problem != nil && err == nil {
			ret.ProblemId = int(problem.Id)
		} else {
			ret.ProblemId = -1
		}
	}

	if v, ok := r.URL.Query()["score"]; ok {
		last := len(v) - 1
		if val, err := strconv.Atoi(v[last]); err == nil {
			ret.Score = val
		}
	} else {
		ret.Score = -1
	}

	if v, ok := r.URL.Query()["status"]; ok {
		ret.Statuses = convertUrlValuesToStatusArr(v)
	}

	if v, ok := r.URL.Query()["compileError"]; ok {
		last := len(v) - 1
		if val, err := strconv.ParseBool(v[last]); err == nil {
			ret.CompileError = &val
		}
	}

	if v, ok := r.URL.Query()["limit"]; ok {
		last := len(v) - 1
		if val, err := strconv.Atoi(v[last]); err == nil {
			ret.Limit = val
		}
	}

	if v, ok := r.URL.Query()["offset"]; ok {
		last := len(v) - 1
		if val, err := strconv.Atoi(v[last]); err == nil {
			ret.Offset = val
		}
	}

	return &ret
}
