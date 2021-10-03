package grader

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/eval/checker"
	"github.com/marius004/phoenix/eval/sandbox"
	"github.com/marius004/phoenix/eval/tasks"
	"github.com/marius004/phoenix/internal"
	"github.com/marius004/phoenix/internal/models"
)

type Grader struct {
	iterationInterval time.Duration

	services *internal.EvaluatorServices
	manager  internal.SandboxManager

	config *internal.Config
	logger *log.Logger
}

var ctx = context.Background()

func (g *Grader) Handle() {
	ticker := time.NewTicker(g.iterationInterval)

	for range ticker.C {

		submissions, err := g.services.SubmissionService.GetByFilter(ctx, waitingSubmissionFilter)

		// no submissions unevaluated
		if len(submissions) == 0 {
			continue
		}

		// something went wrong
		if err != nil {
			g.logger.Printf("grader internal errror: %e", err)
			return
		}

		for _, submission := range submissions {
			if err := g.services.SubmissionService.Update(ctx, int(submission.Id), workingSubmissionUpdate); err != nil {
				g.logger.Println(err)
				continue
			}

			g.handleSubmission(submission)
		}
	}
}

func (g *Grader) getAppropriateChecker() internal.Checker {
	// TO BE CONTINUED :)
	return checker.NewChecker(g.services, g.config, g.logger)
}

func (g *Grader) handleSubmission(submission *models.Submission) {
	if !g.compileSubmission(submission) {
		fmt.Printf("Could not compile submission %d\n", submission.Id)
		return
	}

	if !g.executeSubmission(submission) {
		fmt.Printf("Could not execute submission %d\n", submission.Id)
		return
	}

	checker := g.getAppropriateChecker()
	if err := checker.Check(submission); err != nil {
		fmt.Printf("Could not check the submission %d\n", submission.Id)
		return
	}
}

// compileSubmission compiles the given submission.
// returns true if the compilation was succesful, false otherwise
func (g *Grader) compileSubmission(submission *models.Submission) bool {
	g.logger.Printf("compiling submission %d\n", submission.Id)

	compile := &tasks.CompileTask{
		Config: g.config,
		Logger: g.logger,

		Request: &internal.CompileRequest{
			ID:   int(submission.Id),
			Lang: string(submission.Lang),
			Code: []byte(submission.SourceCode),
		},

		Response: &internal.CompileResponse{},
	}

	// try to compile
	if err := g.manager.RunTask(ctx, compile); err != nil {
		g.logger.Println(err)

		updateSubmission := &models.UpdateSubmissionRequest{
			Message: submission.Message,
		}

		if err := g.services.SubmissionService.Update(ctx, int(submission.Id), updateSubmission); err != nil {
			g.logger.Println(err)
		}

		return false
	}

	// compilation fail
	if !compile.Response.Success {
		updateSubmission := models.NewUpdateSubmissionRequest(0, models.Finished, compile.Response.Message, &hasError)

		if err := g.services.SubmissionService.Update(ctx, int(submission.Id), updateSubmission); err != nil {
			g.logger.Println(err)
		}

		return false
	}

	if compile.Response.Message != "" {
		update := &models.UpdateSubmissionRequest{Message: compile.Response.Message}

		if err := g.services.SubmissionService.Update(ctx, int(submission.Id), update); err != nil {
			g.logger.Println(err)
		}
	}

	return true
}

func (g *Grader) handleGraderInternalErr(submission *models.Submission, message string) {
	updateSubmission := models.NewUpdateSubmissionRequest(0, models.Finished, "grader internal error: "+message, &noError)

	if err := g.services.SubmissionService.Update(context.Background(), int(submission.Id), updateSubmission); err != nil {
		g.logger.Println(err)
	}
}

func (g *Grader) executeTest(submission *models.Submission, problem *models.Problem, test *models.Test) (*tasks.ExecuteTask, error) {
	input, err := g.services.TestManager.GetInputTest(uint(test.Id), problem.Name)
	fmt.Printf("---%s---", string(input))
	g.logger.Printf("---%s---", string(input))

	if err != nil {
		g.logger.Println(err)
		return nil, err
	}

	return &tasks.ExecuteTask{
		Config: g.config,
		Logger: g.logger,

		Request: &internal.ExecuteRequest{
			ID: int(submission.Id),

			SubmissionId: int(submission.Id),
			TestId:       int(test.Id),

			Limit: internal.Limit{
				Time:   problem.TimeLimit,
				Memory: problem.MemoryLimit,
				Stack:  problem.StackLimit,
			},

			Input:     input,
			IsConsole: problem.IsConsoleProblem(),

			ProblemName: problem.Name,
			Lang:        string(submission.Lang),
			BinaryPath:  eval.GetBinaryName(g.config, int(submission.Id)),
		},

		Response: &internal.ExecuteResponse{},
	}, nil
}

func (g *Grader) executeSubmission(submission *models.Submission) bool {
	problem, err := g.services.ProblemService.GetById(ctx, int(submission.ProblemId))

	if err != nil {
		g.logger.Println("could not fetch problem", err)
		g.handleGraderInternalErr(submission, "could not fetch problem")
		return false
	}

	tests, err := g.services.TestService.GetAllProblemTests(ctx, submission.ProblemId)
	if err != nil {
		g.logger.Println("could not fetch tests", err)
		g.handleGraderInternalErr(submission, "could not fetch tests")
		return false
	}

	var wg sync.WaitGroup
	for _, test := range tests {
		wg.Add(1)

		go func(test models.Test) {
			defer wg.Done()

			defer fmt.Println(test)

			execute, err := g.executeTest(submission, problem, &test)
			if err != nil {
				g.logger.Println(err)

				submissionTest := models.NewSubmissionTest(submission.Id, test.Id)
				submissionTest.Message = fmt.Sprintf("grader internal error: %e", err)

				if err := g.services.SubmissionTestService.Create(context.Background(), submissionTest); err != nil {
					g.logger.Println(err)
				}

				return
			}

			if err := g.manager.RunTask(ctx, execute); err != nil {
				g.logger.Println(err)

				submissionTest := models.NewSubmissionTest(submission.Id, test.Id)
				submissionTest.Message = "internal grader error: could not execute test"

				if err := g.services.SubmissionTestService.Create(ctx, submissionTest); err != nil {
					g.logger.Println(err)
				}

				return
			}

			submissionTest := &models.SubmissionTest{
				Time:   execute.Response.TimeUsed,
				Memory: execute.Response.MemoryUsed,

				Message:  execute.Response.Message,
				ExitCode: execute.Response.ExitCode,

				SubmissionId: uint64(execute.Request.SubmissionId),
				TestId:       uint64(execute.Request.TestId),
			}

			if g.services.SubmissionTestService.Create(ctx, submissionTest); err != nil {
				g.logger.Println(err)
				return
			}

		}(*test)
	}

	wg.Wait()

	return true
}

func NewGrader(interval time.Duration, services *internal.EvaluatorServices, config *internal.Config, logger *log.Logger) *Grader {
	manager := sandbox.NewManager(config, logger)

	os.Mkdir(config.CompilePath, 0777)
	os.Mkdir(config.OutputPath, 0777)

	return &Grader{
		iterationInterval: interval,

		services: services,
		manager:  manager,

		config: config,
		logger: logger,
	}
}
