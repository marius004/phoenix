package api

import (
	"net/http"

	"github.com/marius004/phoenix/internal/util"
)

// GET /api/submission-tests/{submissionId}
func (s *API) GetSubmissionTests(w http.ResponseWriter, r *http.Request) {
	submission := util.SubmissionFromRequestContext(r.Context())
	submissionTests, err := s.submissionTestService.GetBySubmissionId(r.Context(), submission.Id)

	if err != nil {
		s.logger.Println(err)
		util.EmptyResponse(w, http.StatusInternalServerError)
		return
	}

	util.DataResponse(w, http.StatusOK, submissionTests, s.logger)
}

func (s *API) GetSubmissionTestById(w http.ResponseWriter, r *http.Request) {
	submissionTest := util.SubmissionFromRequestContext(r.Context())
	util.DataResponse(w, http.StatusOK, submissionTest, s.logger)
}
