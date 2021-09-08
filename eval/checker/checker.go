package checker

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/internal"
	"github.com/marius004/phoenix/internal/models"
)

// Checker implements internal.Checker
type Checker struct {
	services *internal.EvaluatorServices

	config *internal.Config
	logger *log.Logger
}

var ctx = context.Background()
var noError = false

func (c *Checker) handleCheckerInternalErr(submission *models.Submission, message string) {
	updateSubmission := models.NewUpdateSubmissionRequest(0, models.Finished, "checker internal error: "+message, &noError)

	if err := c.services.SubmissionService.Update(context.Background(), int(submission.Id), updateSubmission); err != nil {
		c.logger.Println(err)
	}
}

func (c *Checker) openSubmissionOutputFile(submission *models.Submission, test *models.Test) ([]byte, error) {
	path := eval.GetOutputFileName(c.config, submission, test)
	file, err := os.OpenFile(path, os.O_RDONLY, 0664)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	return ioutil.ReadAll(file)
}

func (c *Checker) parseString(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Trim(str, "\n")
	return str
}

func (c *Checker) isValidAnswer(received, expected []byte) bool {
	receivedStr := c.parseString(string(received))
	expectedStr := c.parseString(string(expected))

	return receivedStr == expectedStr
}

func (c *Checker) Check(submission *models.Submission) error {
	problem, err := c.services.ProblemService.GetById(ctx, int(submission.ProblemId))

	if err != nil {
		c.logger.Println("could not fetch problem", err)
		c.handleCheckerInternalErr(submission, "could not fetch problem")
		return err
	}

	tests, err := c.services.TestService.GetAllProblemTests(ctx, submission.ProblemId)

	if err != nil {
		c.logger.Println("could not fetch tests", err)
		c.handleCheckerInternalErr(submission, "could not fetch tests")
		return err
	}

	score := 0
	for _, test := range tests {

		submissionTest, err := c.services.SubmissionTestService.GetBySubmissionAndTestId(context.Background(), submission.Id, test.Id)

		if submissionTest.Message != "" {
			continue
		}

		if err != nil {
			update := &models.UpdateSubmissionTestRequest{
				Score:  0,
				Time:   -1, // -1 == skip
				Memory: -1,

				Message:  "checker internal error: could not get submission test",
				ExitCode: -1,
			}

			if err := c.services.SubmissionTestService.Update(ctx, submission.Id, test.Id, update); err != nil {
				c.logger.Println(err)
			}

			continue
		}

		if submissionTest.Time > problem.TimeLimit {
			update := &models.UpdateSubmissionTestRequest{
				Score:  0,
				Time:   -1, // -1 == skip
				Memory: -1,

				Message:  "Time Limit Exceeded",
				ExitCode: -1,
			}

			if err := c.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, update); err != nil {
				c.logger.Println(err)
			}

			continue
		}

		if submissionTest.Memory > problem.MemoryLimit {
			update := &models.UpdateSubmissionTestRequest{
				Score:  0,
				Time:   -1, // -1 == skip
				Memory: -1,

				Message:  "Memory Limit Exceeded",
				ExitCode: -1,
			}

			if err := c.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, update); err != nil {
				c.logger.Println(err)
			}

			continue
		}

		output, err := c.services.TestManager.GetOutputTest(uint(test.Id), problem.Name)
		if err != nil {
			c.logger.Println(err)

			update := &models.UpdateSubmissionTestRequest{
				Score:  0,
				Time:   -1, // -1 == skip
				Memory: -1,

				Message:  "checker internal error: could not open the output test file",
				ExitCode: -1,
			}

			if err := c.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, update); err != nil {
				c.logger.Println(err)
			}

			continue
		}

		submissionOutput, err := c.openSubmissionOutputFile(submission, test)
		if err != nil {
			c.logger.Println(err)

			update := &models.UpdateSubmissionTestRequest{
				Score:  0,
				Time:   -1, // -1 == skip
				Memory: -1,

				Message:  "checker internal error: could not open the output submission file",
				ExitCode: -1,
			}

			if err := c.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, update); err != nil {
				c.logger.Println(err)
			}

			continue
		}

		var message = "Correct Answer"
		var testScore = test.Score

		if !c.isValidAnswer(submissionOutput, output) {
			message = "Wrong Answer"
			testScore = 0
		}

		score += testScore

		updateTest := &models.UpdateSubmissionTestRequest{
			Score:  testScore,
			Time:   -1,
			Memory: -1, // -1 == skip

			Message:  message,
			ExitCode: -1,
		}

		if err := c.services.SubmissionTestService.Update(context.Background(), submission.Id, test.Id, updateTest); err != nil {
			c.logger.Println(err)
		}
	}

	updateSubmission := &models.UpdateSubmissionRequest{
		Score:           score,
		Status:          models.Finished,
		Message:         "Evaluated",
		HasCompileError: &noError,
	}

	if err := c.services.SubmissionService.Update(context.Background(), int(submission.Id), updateSubmission); err != nil {
		c.logger.Println(err)
		return err
	}

	return nil
}

func NewChecker(services *internal.EvaluatorServices, config *internal.Config, logger *log.Logger) *Checker {
	return &Checker{
		services: services,

		config: config,
		logger: logger,
	}
}
