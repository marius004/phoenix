package api

import (
	"encoding/json"
	"net/http"

	"github.com/marius004/phoenix/models"
	"github.com/marius004/phoenix/util"
)

// GetProblems is the handler behind GET /api/problems/
// URL parameters could be sent for data filtering
// ex. GET /api/problems?authorId=1&difficulty=easy
// params available: authorId(int), Difficulty(string),
// credits(string), stream(string)
func (s *API) GetProblems(w http.ResponseWriter, r *http.Request) {
	filter := models.ParseProblemFilter(r)
	problems, err := s.problemService.GetByFilter(r.Context(), filter)

	if err != nil {
		s.logger.Println(err)
		util.EmptyResponse(w, http.StatusBadRequest)
		return
	}

	util.DataResponse(w, http.StatusOK, problems, s.logger)
}

// GetProblemByName is the handler behind GET /api/problems/{problemName}
func (s *API) GetProblemByName(w http.ResponseWriter, r *http.Request) {
	problem := util.ProblemFromRequestContext(r)
	util.DataResponse(w, http.StatusOK, problem, s.logger)
}

// CreateProblem is the handler behind POST /api/problems/
func (s *API) CreateProblem(w http.ResponseWriter, r *http.Request) {

	var data models.CreateProblemRequest

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

	if exists, err := s.problemService.ExistsByName(r.Context(), data.Name); exists || err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "problem already exists", s.logger)
		return
	}

	author := util.UserFromRequestContext(r)

	problem := models.NewProblem(data)
	err := s.problemService.Create(r.Context(), problem, int(author.Id))

	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "could not create problem", s.logger)
		return
	}

	util.EmptyResponse(w, http.StatusCreated)
}

// CreateProblem is the handler behind PUT /api/problems/{problemName}
func (s *API) UpdateProblemByName(w http.ResponseWriter, r *http.Request) {

	problem := util.ProblemFromRequestContext(r)
	if !s.canManageProblemResources(r, problem) {
		util.ErrorResponse(w, http.StatusInternalServerError, "you cannot update this problem", s.logger)
		return
	}

	var data models.UpdateProblemRequest

	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()

	if err := jsonDecoder.Decode(&data); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), s.logger)
		return
	}

	if err := s.problemService.Update(r.Context(), int(problem.Id), &data); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "could not update problem", s.logger)
		return
	}

	if problem.Name != data.Name && data.Name != "" {
		if err := s.testManager.RenameProblemDirectory(problem.Name, data.Name); err != nil {
			s.logger.Println(err)
			util.ErrorResponse(w, http.StatusInternalServerError, "Could not update tests directory", s.logger)
			return
		}
	}

	util.DataResponse(w, http.StatusNoContent, "successfully updated the problem", s.logger)
}

// CreateProblem is the handler behind DELETE /api/problems/{problemName}
func (s *API) DeleteByName(w http.ResponseWriter, r *http.Request) {
	problem := util.ProblemFromRequestContext(r)
	if !s.canManageProblemResources(r, problem) {
		util.ErrorResponse(w, http.StatusUnauthorized, "You cannot delete a problem", s.logger)
		return
	}

	if err := s.problemService.Delete(r.Context(), int(problem.Id)); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not delete problem", s.logger)
		return
	}

	if err := s.testManager.DeleteAllTests(problem.Name); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not delete problem tests", s.logger)
		return
	}

	util.EmptyResponse(w, http.StatusOK)
}
