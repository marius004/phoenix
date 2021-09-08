package grader

import "github.com/marius004/phoenix/internal/models"

var (
	hasError = true
	noError  = false

	waitingSubmissionFilter = &models.SubmissionFilter{
		Statuses: []models.SubmissionStatus{models.SubmissionStatus("waiting")},
	}

	workingSubmissionUpdate = &models.UpdateSubmissionRequest{
		Status: models.SubmissionStatus("working"),
	}
)
