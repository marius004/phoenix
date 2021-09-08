package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/marius004/phoenix/database"
	"github.com/marius004/phoenix/eval/grader"
	"github.com/marius004/phoenix/internal"
	"github.com/marius004/phoenix/internal/services"
	"github.com/marius004/phoenix/managers"
)

type API struct {
	db     *sqlx.DB
	config *internal.Config
	logger *log.Logger

	userService    services.UserService
	problemService services.ProblemService
	testService    services.TestService

	submissionService     services.SubmissionService
	submissionTestService services.SubmissionTestService

	grader      *grader.Grader
	testManager *managers.TestManager
}

// New declares a new API instance
func New(db *database.DB, config *internal.Config, logger *log.Logger) *API {
	var (
		userService    = db.UserService(logger)
		problemService = db.ProblemService(logger)
		testService    = db.TestService(logger)

		submissionService     = db.SubmissionService(logger)
		submissionTestService = db.SubmissionTestService(logger)

		testManager = managers.NewTestManager("tests")

		graderServices = evaluatorServices(problemService, submissionService, submissionTestService, testService, *testManager)
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
		grader:      grader.NewGrader(100*time.Millisecond, graderServices, config, logger),
	}
}

// Routes returns the handler that will be used for the route /api
func (s *API) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(s.UserCtx)
	go s.grader.Handle()

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

func evaluatorServices(problemService services.ProblemService, submissionService services.SubmissionService,
	submissionTestService services.SubmissionTestService, testService services.TestService,
	testManager managers.TestManager) *internal.EvaluatorServices {
	return &internal.EvaluatorServices{
		ProblemService: problemService,

		SubmissionService:     submissionService,
		SubmissionTestService: submissionTestService,

		TestService: testService,
		TestManager: testManager,
	}
}
