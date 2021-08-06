package evaluator

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/eval/container"
	"github.com/marius004/phoenix/models"
)

type Evaluator struct {
	compileHandler eval.Handler
	compileChannel chan *models.Submission

	executeHandler eval.Handler
	executeChannel chan *models.Submission

	checkerHandler eval.Handler
	checkerChannel chan *models.Submission

	finalizerHandler eval.Handler
	finalizerChannel chan *models.Submission

	dbIterationTimeout time.Duration
	services           *eval.EvaluatorServices
	sandboxManager     eval.SandboxManager

	config *models.Config
	logger *log.Logger
}

func (e *Evaluator) Serve() {
	ticker := time.NewTicker(e.dbIterationTimeout)

	go e.compileHandler.Handle(e.executeChannel)
	go e.executeHandler.Handle(e.checkerChannel)
	go e.checkerHandler.Handle(e.finalizerChannel)
	go e.finalizerHandler.Handle(nil)

	for range ticker.C {
		submissions, err := e.services.SubmissionService.GetByFilter(context.Background(), waitingSubmissionFilter)

		if err != nil || len(submissions) == 0 {
			continue
		}

		for _, submission := range submissions {
			if err := e.services.SubmissionService.Update(context.Background(), int(submission.Id), workingSubmissionUpdate); err != nil {
				e.logger.Println(err)
				continue
			}

			e.compileChannel <- submission
		}
	}
}

func New(dbIterationTimeout time.Duration, services *eval.EvaluatorServices, config *models.Config) *Evaluator {
	os.Mkdir(config.CompilePath, 0775)
	os.Mkdir(config.OutputPath, 0755)

	logger := newLogger(config.Eval.LoggerPath)
	sandboxManager := container.NewManager(config, logger)

	var (
		compileChannel   = make(chan *models.Submission, config.MaxCompile)
		executeChannel   = make(chan *models.Submission, config.MaxExecute)
		checkerChannel   = make(chan *models.Submission, config.MaxCheck)
		finalizerChannel = make(chan *models.Submission, config.MaxExecute)
	)

	return &Evaluator{
		compileChannel: compileChannel,
		compileHandler: NewCompileHandler(config, logger, compileChannel, services.CompileServices(), sandboxManager),

		executeChannel: executeChannel,
		executeHandler: NewExecuteHandler(config, logger, executeChannel, services.ExecuteServices(), sandboxManager),

		checkerHandler: NewCheckerHandler(config, logger, checkerChannel, services.CheckerServices()),
		checkerChannel: checkerChannel,

		finalizerHandler: NewFinalizeHandler(config, logger, finalizerChannel, services),
		finalizerChannel: finalizerChannel,

		dbIterationTimeout: dbIterationTimeout,

		services:       services,
		sandboxManager: sandboxManager,

		logger: newLogger(config.Eval.LoggerPath),
		config: config,
	}
}
