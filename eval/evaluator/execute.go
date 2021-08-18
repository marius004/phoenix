package evaluator

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/eval/tasks"
	"github.com/marius004/phoenix/models"
	"golang.org/x/sync/semaphore"
)

type ExecuteHandler struct {
	config *models.Config
	logger *log.Logger

	semaphore *semaphore.Weighted
	// submissions is the channel we receive new submissions from.
	submissions chan *models.Submission

	services       *eval.ExecuteServices
	sandboxManager eval.SandboxManager

	debug bool
}

// submissionData is the data needed to execute a submission
type submissionData struct {
	problem *models.Problem
	tests   []*models.Test

	err      error
	internal error
}

func (handler *ExecuteHandler) getSubmissionData(submission *models.Submission) *submissionData {
	var data submissionData

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

func (handler *ExecuteHandler) createExecuteTask(submission *models.Submission, data *submissionData, test *models.Test) (*tasks.ExecuteTask, error) {

	input, err := handler.services.TestManager.GetInputTest(uint(test.Id), data.problem.Name)

	if err != nil {
		handler.logger.Println(err)
		return nil, err
	}

	execute := &tasks.ExecuteTask{
		Config: handler.config,
		Logger: handler.logger,

		Request: &eval.ExecuteRequest{
			ID: int(submission.Id),

			SubmissionId: int(submission.Id),
			TestId:       int(test.Id),

			Limit: eval.Limit{
				Time:   data.problem.TimeLimit,
				Memory: data.problem.MemoryLimit,
				Stack:  data.problem.StackLimit,
			},

			Input:     input,
			IsConsole: data.problem.IsConsoleProblem(),

			ProblemName: data.problem.Name,
			Lang:        string(submission.Lang),
			BinaryPath:  eval.GetBinaryName(handler.config, int(submission.Id)),
		},

		Response: &eval.ExecuteResponse{},
	}

	return execute, nil
}

func (handler *ExecuteHandler) wasTaskExecutedSuccesfully(task *tasks.ExecuteTask) bool {
	return task.Response.ExitCode == 0
}

func (handler *ExecuteHandler) Handle(next chan *models.Submission) {
	for submission := range handler.submissions {

		handler.logger.Printf("Executing submission %d\n", submission.Id)
		data := handler.getSubmissionData(submission)

		if data.err != nil {
			updateSubmission := models.NewUpdateSubmissionRequest(0, models.Finished, "evaluator error: "+data.err.Error(), &noError)
			if err := handler.services.SubmissionService.Update(context.Background(), int(submission.Id), updateSubmission); err != nil {
				handler.logger.Println(err)
			}
			return
		}

		var wg sync.WaitGroup

		for _, test := range data.tests {

			wg.Add(1)

			go func(test *models.Test) {
				defer wg.Done()
				execute, err := handler.createExecuteTask(submission, data, test)

				if err != nil {
					handler.logger.Println(err)

					submissionTest := models.NewSubmissionTest(submission.Id, test.Id)
					if err := handler.services.SubmissionTestService.Create(context.Background(), submissionTest); err != nil {
						handler.logger.Println(err)
					}

					return
				}

				if err := handler.sandboxManager.RunTask(context.Background(), execute); err != nil {
					handler.logger.Println(err)

					submissionTest := models.NewSubmissionTest(submission.Id, test.Id)
					submissionTest.Message = "internal evaluator error: could not execute test"

					if err := handler.services.SubmissionTestService.Create(context.Background(), submissionTest); err != nil {
						handler.logger.Println(err)
					}

					return
				}

				if err := handler.sandboxManager.RunTask(context.Background(), execute); err != nil {
					handler.logger.Println(err)

					submissionTest := models.NewSubmissionTest(submission.Id, test.Id)
					submissionTest.Message = "internal evaluator error: could not execute test"

					if err := handler.services.SubmissionTestService.Create(context.Background(), submissionTest); err != nil {
						handler.logger.Println(err)
					}

					return
				}

				var submissionTest *models.SubmissionTest

				if !handler.wasTaskExecutedSuccesfully(execute) {
					submissionTest = &models.SubmissionTest{
						Time:   execute.Response.TimeUsed,
						Memory: execute.Response.MemoryUsed,

						Message:  execute.Response.Message,
						ExitCode: execute.Response.ExitCode,

						SubmissionId: uint64(execute.Request.SubmissionId),
						TestId:       uint64(execute.Request.TestId),
					}
				} else {
					submissionTest = &models.SubmissionTest{
						Time:   execute.Response.TimeUsed,
						Memory: execute.Response.MemoryUsed,

						Message:  execute.Response.Message,
						ExitCode: execute.Response.ExitCode,

						SubmissionId: uint64(execute.Request.SubmissionId),
						TestId:       uint64(execute.Request.TestId),
					}
				}

				if err := handler.services.SubmissionTestService.Create(context.Background(), submissionTest); err != nil {
					handler.logger.Println(err)
					return
				}

			}(test)
		}

		wg.Wait()

		if next != nil {
			next <- submission
		}
	}
}

func NewExecuteHandler(config *models.Config, logger *log.Logger, channel chan *models.Submission,
	services *eval.ExecuteServices, sandboxManager eval.SandboxManager, debug bool) *ExecuteHandler {
	return &ExecuteHandler{
		config: config,
		logger: logger,

		submissions: channel,
		semaphore:   semaphore.NewWeighted(int64(config.MaxExecute)),

		services:       services,
		sandboxManager: sandboxManager,

		debug: debug,
	}
}
