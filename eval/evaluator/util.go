package evaluator

import (
	"log"
	"os"

	"github.com/marius004/phoenix/models"
)

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

func newLogger(path string) *log.Logger {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	logger := &log.Logger{}
	logger.SetFlags(log.LstdFlags | log.Ldate | log.Llongfile)
	logger.SetOutput(file)

	return logger
}

func UpdateSubmissionErr(message string) *models.UpdateSubmissionRequest {
	return &models.UpdateSubmissionRequest{
		Score:           0,
		Status:          models.SubmissionStatus("finished"),
		Message:         message,
		HasCompileError: &hasError,
	}
}

func UpdateSubmissionInternalErr(message string) *models.UpdateSubmissionRequest {
	return &models.UpdateSubmissionRequest{
		Score:           0,
		Status:          models.SubmissionStatus("finished"),
		Message:         "Internal server error: " + message,
		HasCompileError: &noError,
	}
}
