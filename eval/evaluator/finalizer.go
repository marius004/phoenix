package evaluator

import (
	"bytes"
	"context"
	"log"
	"os/exec"
	"strconv"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/models"
)

// Finilizer takes care of cleaning up and updating submissions as well
type FinalizerHandler struct {
	config *models.Config
	logger *log.Logger

	submissions chan *models.Submission
	services    *eval.EvaluatorServices
}

func (handler *FinalizerHandler) UpdateSubmission(submission *models.Submission) {
	tests, err := handler.services.SubmissionTestService.GetBySubmissionId(context.Background(), submission.Id)

	if err != nil {
		handler.logger.Print(err)
		return
	}

	score := 0
	for _, test := range tests {
		score += test.Score
	}

	updateSubmission := &models.UpdateSubmissionRequest{
		Score:           score,
		Status:          models.Finished,
		Message:         "Evaluated",
		HasCompileError: &noError,
	}

	if err := handler.services.SubmissionService.Update(context.Background(), int(submission.Id), updateSubmission); err != nil {
		handler.logger.Println(err)
		return
	}
}

// TODO later
func (handler *FinalizerHandler) Cleanup(submission *models.Submission) {
	rmBinCmd := exec.Command("rm", strconv.Itoa(int(submission.Id))+".bin")
	rmBinCmd.Path = handler.config.CompilePath

	var output bytes.Buffer

	rmBinCmd.Stdout = &output
	rmBinCmd.Stderr = &output

	if err := rmBinCmd.Run(); err != nil {
		handler.logger.Println(err)
	}
}

func (handler *FinalizerHandler) Handle(next chan *models.Submission) {
	for submission := range handler.submissions {
		handler.UpdateSubmission(submission)
		handler.Cleanup(submission)

		if next != nil {
			next <- submission
		}
	}
}

func NewFinalizeHandler(config *models.Config, logger *log.Logger,
	channel chan *models.Submission, services *eval.EvaluatorServices) *FinalizerHandler {
	return &FinalizerHandler{
		config:      config,
		logger:      logger,
		submissions: channel,
		services:    services,
	}
}
