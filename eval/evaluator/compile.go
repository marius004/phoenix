package evaluator

import (
	"context"
	"log"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/eval/tasks"
	"github.com/marius004/phoenix/models"
	"golang.org/x/sync/semaphore"
)

type CompileHandler struct {
	config *models.Config
	logger *log.Logger

	// submissions is the channel we receive new submissions from.
	semaphore   *semaphore.Weighted
	submissions chan *models.Submission

	services       *eval.CompileServices
	sandboxManager eval.SandboxManager
}

func (handler *CompileHandler) Handle(next chan *models.Submission) {
	for submission := range handler.submissions {

		if err := handler.semaphore.Acquire(context.Background(), 1); err != nil {
			handler.logger.Println(err)
			continue
		}

		handler.logger.Printf("Compiling submission %d\n", submission.Id)

		go func(submission *models.Submission) {
			defer handler.semaphore.Release(1)

			compile := &tasks.CompileTask{
				Config: handler.config,
				Logger: handler.logger,
				Request: &eval.CompileRequest{
					ID:   int(submission.Id),
					Lang: string(submission.Lang),
					Code: []byte(submission.SourceCode),
				},
				Response: &eval.CompileResponse{},
			}

			// try to compile
			if err := handler.sandboxManager.RunTask(context.Background(), compile); err != nil {
				handler.logger.Println(err)
				if err := handler.services.SubmissionService.Update(context.Background(), int(submission.Id), UpdateSubmissionInternalErr(err.Error())); err != nil {
					handler.logger.Println(err)
					return
				}
			}

			// compilation fail
			if !compile.Response.Success {
				if err := handler.services.SubmissionService.Update(context.Background(), int(submission.Id), UpdateSubmissionErr(compile.Response.Output)); err != nil {
					handler.logger.Println(err)
					return
				}
			}

			// succesful compilation
			// this means that we will send the submission to the next handler(if one is present)
			// (in this case send the submission to be executed)
			if next != nil {
				next <- submission
			}

		}(submission)
	}
}

func NewCompileHandler(config *models.Config, logger *log.Logger, channel chan *models.Submission,
	services *eval.CompileServices, sandboxManager eval.SandboxManager) *CompileHandler {
	return &CompileHandler{
		config: config,
		logger: logger,

		semaphore:   semaphore.NewWeighted(int64(config.MaxSandboxes)),
		submissions: channel,

		services:       services,
		sandboxManager: sandboxManager,
	}
}
