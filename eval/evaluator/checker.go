package evaluator

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/models"
)

type CheckerHandler struct {
	config *models.Config
	logger *log.Logger

	submissions chan *models.Submission
	services    *eval.CheckerServices
}

type checkerData struct {
	problem *models.Problem
	tests   []*models.Test

	err      error
	internal error
}

func (handler *CheckerHandler) getCheckerData(submission *models.Submission) *checkerData {
	var data checkerData

	data.problem, data.err = handler.services.ProblemService.GetById(context.Background(), submission.ProblemId)

	if data.err != nil {
		data.internal = errors.New("could not fetch problem")
		handler.logger.Println(data.err)
		return &data
	}

	data.tests, data.err = handler.services.TestService.GetAllProblemTests(context.Background(), submission.ProblemId)

	if data.err != nil {
		data.internal = errors.New("could not fetch problem tests")
		handler.logger.Println(data.err)
		return &data
	}

	return &data
}

func (handler *CheckerHandler) openSubmissionOutputFile(submission *models.Submission, test *models.Test) ([]byte, error) {
	path := eval.GetOutputFileName(handler.config, submission, test)
	file, err := os.OpenFile(path, os.O_RDONLY, 0664)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	return ioutil.ReadAll(file)
}

func (c *CheckerHandler) parseString(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Trim(str, "\n")
	return str
}

func (handler *CheckerHandler) Check(received, expected []byte) bool {
	receivedStr := handler.parseString(string(received))
	expectedStr := handler.parseString(string(expected))

	return receivedStr == expectedStr
}

func (handler *CheckerHandler) Handle(next chan *models.Submission) {
	for submission := range handler.submissions {
		data := handler.getCheckerData(submission)

		if data.err != nil {
			updateSubmission := models.NewUpdateSubmissionRequest(0, models.Finished, "evaluator error: "+data.err.Error(), &noError)
			if err := handler.services.SubmissionService.Update(context.Background(), int(submission.Id), updateSubmission); err != nil {
				handler.logger.Println(err)
			}
			return
		}

		for _, test := range data.tests {
			submissionTest, err := handler.services.SubmissionTestService.GetBySubmissionAndTestId(context.Background(), submission.Id, test.Id)

			if err != nil {
				update := &models.UpdateSubmissionTestRequest{
					Score:  0,
					Time:   -1, // -1 == skip
					Memory: -1,

					Message:  "evaluator internal error: could not get submission test",
					ExitCode: -1,
				}

				if err := handler.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, update); err != nil {
					handler.logger.Println(err)
				}

				continue
			}

			if submissionTest.Time > data.problem.TimeLimit {
				update := &models.UpdateSubmissionTestRequest{
					Score:  0,
					Time:   -1, // -1 == skip
					Memory: -1,

					Message:  "Time Limit Exceeded",
					ExitCode: -1,
				}

				if err := handler.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, update); err != nil {
					handler.logger.Println(err)
				}

				continue
			}

			if submissionTest.Memory >= data.problem.MemoryLimit {
				update := &models.UpdateSubmissionTestRequest{
					Score:  0,
					Time:   -1, // -1 == skip
					Memory: -1,

					Message:  "Memory Limit Exceeded",
					ExitCode: -1,
				}

				if err := handler.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, update); err != nil {
					handler.logger.Println(err)
				}

				continue
			}

			output, err := handler.services.TestManager.GetOutputTest(uint(test.Id), data.problem.Name)

			if err != nil {
				handler.logger.Println(err)

				update := &models.UpdateSubmissionTestRequest{
					Score:  0,
					Time:   -1, // -1 == skip
					Memory: -1,

					Message:  "Evaluator Internal Error: Could Not Open The Output Test File",
					ExitCode: -1,
				}

				if err := handler.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, update); err != nil {
					handler.logger.Println(err)
				}

				continue
			}

			submissionOutput, err := handler.openSubmissionOutputFile(submission, test)

			if err != nil {
				handler.logger.Println(err)

				update := &models.UpdateSubmissionTestRequest{
					Score:  0,
					Time:   -1, // -1 == skip
					Memory: -1,

					Message:  "evaluator internal error: could not open the output submission file",
					ExitCode: -1,
				}

				if err := handler.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, update); err != nil {
					handler.logger.Println(err)
				}

				continue
			}

			var updateTest *models.UpdateSubmissionTestRequest

			if handler.Check(submissionOutput, output) {
				updateTest = &models.UpdateSubmissionTestRequest{
					Score:  test.Score,
					Time:   -1,
					Memory: -1, // -1 == skip

					Message:  "Correct Answer",
					ExitCode: -1,
				}
			} else {
				updateTest = &models.UpdateSubmissionTestRequest{
					Score:  0,
					Time:   -1,
					Memory: -1, // -1 == skip

					Message:  "Wrong Answer",
					ExitCode: -1,
				}
			}

			if err := handler.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, updateTest); err != nil {
				handler.logger.Println(err)
			}
		}

		if next != nil {
			next <- submission
		}
	}
}

func NewCheckerHandler(config *models.Config, logger *log.Logger,
	channel chan *models.Submission, services *eval.CheckerServices) *CheckerHandler {
	return &CheckerHandler{
		config: config,
		logger: logger,

		submissions: channel,
		services:    services,
	}
}
