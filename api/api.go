package api

import (
	"log"
	"net/http"
	"time"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/eval/evaluator"
	"github.com/marius004/phoenix/managers"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/marius004/phoenix/database"
	"github.com/marius004/phoenix/disk"
	"github.com/marius004/phoenix/models"
	"github.com/marius004/phoenix/services"
)

type API struct {
	db     *sqlx.DB
	config *models.Config
	logger *log.Logger

	userService    services.UserService
	problemService services.ProblemService
	testService    services.TestService

	submissionService     services.SubmissionService
	submissionTestService services.SubmissionTestService

	evaluator   *evaluator.Evaluator
	testManager managers.TestManager
}

// New declares a new API instance
func New(db *database.DB, config *models.Config, logger *log.Logger) *API {
	var (
		userService    = db.UserService(logger)
		problemService = db.ProblemService(logger)
		testService    = db.TestService(logger)

		submissionService     = db.SubmissionService(logger)
		submissionTestService = db.SubmissionTestService(logger)

		testManager = disk.NewTestManager("tests")
		evaluator   = evaluator.New(100*time.Millisecond, evaluatorServices(problemService, submissionService, submissionTestService, testService, testManager), config, true)
	)

	return &API{
		db:     db.Conn,
		config: config,
		logger: logger,

		userService:    userService,
		problemService: problemService,
		testService:    testService,

		submissionService:     submissionService,
		submissionTestService: submissionTestService,

		testManager: testManager,
		evaluator:   evaluator,
	}
}

// Routes returns the handler that will be used for the route /api
func (s *API) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(s.UserCtx)

	r.Route("/users", func(r chi.Router) {
		r.With(s.MustBeAdmin).Get("/", s.GetAllUsers)
		r.Get("/{userName}", s.GetUserByUserName)
		r.Get("/{userId}/gravatar", s.GetUserGravatar)
	})

	r.Route("/auth", func(r chi.Router) {
		r.With(s.MustNotBeAuthed).Post("/signup", s.Signup)
		r.With(s.MustNotBeAuthed).Post("/login", s.Login)
		r.With(s.MustBeAuthed).Post("/logout", s.Logout)
	})

	r.Route("/problems", func(r chi.Router) {
		r.Get("/", s.GetProblems)
		r.With(s.MustBeProposer).Post("/", s.CreateProblem)

		r.With(s.ProblemCtx).Route("/{problemName}", func(r chi.Router) {
			r.Get("/", s.GetProblemByName)

			r.With(s.MustBeProposer).Put("/", s.UpdateProblemByName)
			r.With(s.MustBeProposer).Delete("/", s.DeleteByName)

			r.Route("/tests", func(r chi.Router) {

				r.With(s.MustBeProposer).Get("/", s.GetProblemTests)
				r.With(s.MustBeProposer).Post("/", s.CreateTest)

				r.With(s.TestCtx).Route("/{testId}", func(r chi.Router) {
					r.With(s.MustBeProposer).Get("/", s.GetTestById)
					r.With(s.MustBeProposer).Put("/", s.UpdateTestById)
					r.With(s.MustBeProposer).Delete("/", s.DeleteTestById)
				})
			})
		})
	})

	go s.evaluator.Serve()

	r.Route("/submissions", func(r chi.Router) {
		r.Get("/", s.GetSubmissions)
		r.With(s.MustBeAuthed).Post("/", s.CreateSubmission)

		r.With(s.SubmissionCtx).Route("/{submissionId}", func(r chi.Router) {
			r.With(s.SubmissionCtx).Get("/", s.GetSubmissionById)
			r.Get("/tests", s.GetSubmissionTests)
		})
	})

	return r
}

func evaluatorServices(problemService services.ProblemService, submissionService services.SubmissionService, submissionTestService services.SubmissionTestService, testService services.TestService, testManager managers.TestManager) *eval.EvaluatorServices {
	return &eval.EvaluatorServices{
		ProblemService: problemService,

		SubmissionService:     submissionService,
		SubmissionTestService: submissionTestService,

		TestService: testService,
		TestManager: testManager,
	}
}
