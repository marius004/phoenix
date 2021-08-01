package api

import (
	"encoding/json"
	"net/http"

	"github.com/marius004/phoenix/models"
	"github.com/marius004/phoenix/util"
)

// CreateTest is the handler behind POST /api/tests/
func (s *API) CreateTest(w http.ResponseWriter, r *http.Request) {

	var data models.CreateTestRequest

	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()

	if err := jsonDecoder.Decode(&data); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), s.logger)
		return
	}

	problem, err := s.problemService.GetById(r.Context(), data.ProblemId)
	if problem == nil || err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "Problem with the specified id does not exist", s.logger)
		return
	}

	if !s.canManageProblemResources(r, problem) {
		util.ErrorResponse(w, http.StatusInternalServerError, "you cannot do this", s.logger)
		return
	}

	if err := data.Validate(); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), s.logger)
		return
	}

	test := models.NewTest(data)
	if err := s.testService.Create(r.Context(), test); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not create test", s.logger)
		return
	}

	if err := s.testManager.SaveInputTest(uint(test.Id), problem.Name, data.Input); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not create input file", s.logger)
		return
	}

	if err := s.testManager.SaveOutputTest(uint(test.Id), problem.Name, data.Output); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not create output file", s.logger)
		return
	}
}

// UpdateTestById is the handler behind PUT /api/tests/
func (s *API) UpdateTestById(w http.ResponseWriter, r *http.Request) {
	test := util.TestFromRequestContext(r)
	var data models.UpdateTestRequest

	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()

	if err := jsonDecoder.Decode(&data); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), s.logger)
		return
	}

	problem, err := s.problemService.GetById(r.Context(), test.ProblemId)
	if problem == nil || err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "Problem with the specified id does not exist", s.logger)
		return
	}

	if !s.canManageProblemResources(r, problem) {
		util.ErrorResponse(w, http.StatusInternalServerError, "you cannot do this", s.logger)
		return
	}

	if err := data.Validate(); err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, err.Error(), s.logger)
		return
	}

	models.UpdateTest(test, &data)
	if err := s.testService.Update(r.Context(), int(test.Id), int(problem.Id), test); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not update test", s.logger)
		return
	}

	if data.Input != nil && len(data.Input) > 0 {
		if err := s.testManager.SaveInputTest(uint(test.Id), problem.Name, data.Input); err != nil {
			util.ErrorResponse(w, http.StatusInternalServerError, "Could not update input test", s.logger)
			return
		}
	}

	if data.Output != nil && len(data.Output) > 0 {
		if err := s.testManager.SaveOutputTest(uint(test.Id), problem.Name, data.Output); err != nil {
			util.ErrorResponse(w, http.StatusInternalServerError, "Could not update output test", s.logger)
			return
		}
	}

	util.EmptyResponse(w, http.StatusOK)
}

// GetTestById is the handler behind GET /api/tests/{testId}
func (s *API) GetTestById(w http.ResponseWriter, r *http.Request) {
	test := util.TestFromRequestContext(r)

	problem, err := s.problemService.GetById(r.Context(), test.ProblemId)
	if err != nil || problem == nil {
		util.ErrorResponse(w, http.StatusBadRequest, "Problem does not exist", s.logger)
		return
	}

	// the user cannot see the tests so we don't send the input and output tests
	if !s.canManageProblemResources(r, problem) {
		util.ErrorResponse(w, http.StatusUnauthorized, "You can't do this!", s.logger)
		return
	}

	input, err := s.testManager.GetInputTest(uint(test.Id), problem.Name)
	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not fetch input test", s.logger)
	}

	output, err := s.testManager.GetOutputTest(uint(test.Id), problem.Name)
	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not fetch input test", s.logger)
	}

	fullTest := models.NewFullTest(*test, string(input), string(output))
	util.DataResponse(w, http.StatusOK, fullTest, s.logger)
}

// CreateTest is the handler behind GET /api/tests/
func (s *API) GetAllTests(w http.ResponseWriter, r *http.Request) {
	tests, err := s.testService.GetAllTests(r.Context())

	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not retrive the tests", s.logger)
		return
	}

	util.DataResponse(w, http.StatusOK, tests, s.logger)
}

// DeleteTestById is the handler behind DELETE /api/tests/{testId}
func (s *API) DeleteTestById(w http.ResponseWriter, r *http.Request) {
	test := util.TestFromRequestContext(r)

	problem, err := s.problemService.GetById(r.Context(), test.ProblemId)
	if err != nil || problem == nil {
		util.ErrorResponse(w, http.StatusBadRequest, "Problem does not exist", s.logger)
		return
	}

	if !s.canManageProblemResources(r, problem) {
		util.ErrorResponse(w, http.StatusUnauthorized, "You cannot delete this test!", s.logger)
		return
	}

	if err := s.testService.Delete(r.Context(), int(test.Id), int(problem.Id)); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not delete test", s.logger)
		return
	}

	if err := s.testManager.DeleteInputTest(uint(test.Id), problem.Name); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not delete test", s.logger)
		return
	}

	if err := s.testManager.DeleteOutputTest(uint(test.Id), problem.Name); err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "Could not delete test", s.logger)
		return
	}

	util.EmptyResponse(w, http.StatusOK)
}

// GetProblemTests is the handler behind GET /api/tests/{problemName}
// Returns all the tests for the specified problem
func (s *API) GetProblemTests(w http.ResponseWriter, r *http.Request) {
	problem := util.ProblemFromRequestContext(r)

	tests, err := s.testService.GetAllProblemTests(r.Context(), int(problem.Id))

	if err != nil {
		s.logger.Println(err)
		util.ErrorResponse(w, http.StatusInternalServerError, "could not retrive problem tests", s.logger)
		return
	}

	util.DataResponse(w, http.StatusOK, tests, s.logger)
}
