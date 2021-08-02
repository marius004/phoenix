package api

import (
	"log"
	"net/http"

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

	userService       services.UserService
	problemService    services.ProblemService
	testService       services.TestService
	submissionService services.SubmissionService

	testManager managers.TestManager
}

// New declares a new API instance
func New(db *database.DB, config *models.Config, logger *log.Logger) *API {
	var (
		userService       = db.UserService(logger)
		problemService    = db.ProblemService(logger)
		testService       = db.TestService(logger)
		submissionService = db.SubmissionService(logger)

		testManager = disk.NewTestManager("tests")
	)

	return &API{
		db:     db.Conn,
		config: config,
		logger: logger,

		userService:       userService,
		problemService:    problemService,
		testService:       testService,
		submissionService: submissionService,

		testManager: testManager,
	}
}

// Routes returns the handler that will be used for the route /api
func (s *API) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(s.UserCtx)

	r.Route("/users", func(r chi.Router) {
		r.With(s.MustBeAdmin).Get("/", s.GetAllUsers)
		r.Get("/{userName}", s.GetUserByUserName)
	})

	r.Route("/auth", func(r chi.Router) {
		r.With(s.MustNotBeAuthed).Post("/signup", s.Signup)
		r.With(s.MustNotBeAuthed).Post("/login", s.Login)
		r.With(s.MustBeAuthed).Post("/logout", s.Logout)
	})

	r.Route("/problems", func(r chi.Router) {
		r.Get("/", s.GetProblems)
		r.With(s.MustBeProposer).Post("/", s.CreateProblem)

		r.Route("/{problemName}", func(r chi.Router) {
			r.Use(s.ProblemCtx)
			r.Get("/", s.GetProblemByName)
			r.With(s.MustBeProposer).Put("/", s.UpdateProblemByName)
			r.With(s.MustBeProposer).Delete("/", s.DeleteByName)
		})
	})

	r.Route("/tests", func(r chi.Router) {
		r.With(s.MustBeAdmin).Get("/", s.GetAllTests)
		r.With(s.MustBeProposer).Post("/", s.CreateTest)

		r.Route("/{testId}", func(r chi.Router) {
			r.Use(s.TestCtx)
			r.With(s.MustBeProposer).Get("/", s.GetTestById)
			r.With(s.MustBeProposer).Put("/", s.UpdateTestById)
			r.With(s.MustBeProposer).Delete("/", s.DeleteTestById)
		})

		r.With(s.ProblemCtx, s.MustBeProposer).Get("/{problemName}", s.GetProblemTests)
	})

	r.Route("/submissions", func(r chi.Router) {
		r.Get("/", s.GetSubmissions)
		r.With(s.SubmissionCtx).Get("/{submissionId}", s.GetSubmissionById)
		r.With(s.MustBeAuthed).Post("/", s.CreateSubmission)
	})

	return r
}
