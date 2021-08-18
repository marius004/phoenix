package evaluator

import (
	"context"
	"log"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/eval/tasks"
	"github.com/marius004/phoenix/models"
)

type CompileHandler struct {
	config *models.Config
	logger *log.Logger

	// submissions is the channel we receive new submissions from.
	submissions chan *models.Submission

	services       *eval.CompileServices
	sandboxManager eval.SandboxManager

	debug bool
}

func (handler *CompileHandler) Handle(next chan *models.Submission) {
	for submission := range handler.submissions {
		handler.logger.Printf("Compiling submission %d\n", submission.Id)

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

			updateSubmission := &models.UpdateSubmissionRequest{
				Message: submission.Message,
			}

			if err := handler.services.SubmissionService.Update(context.Background(), int(submission.Id), updateSubmission); err != nil {
				handler.logger.Println(err)
				return
			}
		}

		// compilation fail
		if !compile.Response.Success {
			updateSubmission := models.NewUpdateSubmissionRequest(0, models.Finished, compile.Response.Message, &hasError)
			if err := handler.services.SubmissionService.Update(context.Background(), int(submission.Id), updateSubmission); err != nil {
				handler.logger.Println(err)
			}
			return
		}

		if compile.Response.Message != "" {
			update := &models.UpdateSubmissionRequest{Message: compile.Response.Message}
			if err := handler.services.SubmissionService.Update(context.Background(), int(submission.Id), update); err != nil {
				handler.logger.Println(err)
			}
		}

		// succesful compilation
		// this means that we will send the submission to the next handler(if one is present)
		// (in this case send the submission to be executed)
		if next != nil {
			next <- submission
		}
	}
}

func NewCompileHandler(config *models.Config, logger *log.Logger, channel chan *models.Submission,
	services *eval.CompileServices, sandboxManager eval.SandboxManager, debug bool) *CompileHandler {
	return &CompileHandler{
		config: config,
		logger: logger,

		submissions: channel,

		services:       services,
		sandboxManager: sandboxManager,

		debug: debug,
	}
}
