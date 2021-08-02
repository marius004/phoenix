package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/marius004/phoenix/util"
)

// MustNotBeAuthed is a middleware that makes sure that the user creating the request is not authenticated.
func (s *API) MustNotBeAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if util.IsRAuthed(r) {
			util.ErrorResponse(w, http.StatusUnauthorized, "You must not be logged in to do this", s.logger)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// MustBeAuthed is a middleware that makes sure that the user creating the request is authenticated.
func (s *API) MustBeAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !util.IsRAuthed(r) {
			util.ErrorResponse(w, http.StatusUnauthorized, "You must be logged in to do this", s.logger)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// MustBeAdmin is a middleware that makes sure that the user creating the request is an admin.
func (s *API) MustBeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !util.IsRAdmin(r) {
			util.ErrorResponse(w, http.StatusUnauthorized, "You must be an admin to do this", s.logger)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// MustBeProposer is a middleware that makes sure that the user creating the request is a proposer.
func (s *API) MustBeProposer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !util.IsRProposer(r) {
			util.ErrorResponse(w, http.StatusUnauthorized, "You must be a proposer to do this", s.logger)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// UserCtx is a middleware that attaches the user to the request context.
func (s *API) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie(HttpOnlyCookieAuthKey)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		token, err := util.VerifyToken(authCookie.Value, s.config.JwtSecret())

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId, err := strconv.Atoi(claims["iss"].(string))

			if err != nil {
				s.logger.Println(err)
				next.ServeHTTP(w, r)
				return
			}

			user, err := s.userService.GetById(r.Context(), userId)

			if err != nil {
				s.logger.Println(err)
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), util.UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))

			return
		}

		next.ServeHTTP(w, r)
	})
}

// ProblemCtx middleware is used to attach the problem from the URL parameters to the request context.
func (s *API) ProblemCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		problemName := chi.URLParam(r, "problemName")
		problem, err := s.problemService.GetByName(r.Context(), problemName)

		if err != nil || problem == nil {
			util.DataResponse(w, http.StatusNotFound, "Problem not found", s.logger)
			return
		}

		ctx := context.WithValue(r.Context(), util.ProblemContextKey, problem)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TestCtx middleware is used to attach the test from the URL parameters to the request context.
func (s *API) TestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testId, err := strconv.Atoi(chi.URLParam(r, "testId"))

		if err != nil {
			util.ErrorResponse(w, http.StatusBadRequest, "Invalid test id", s.logger)
			return
		}

		if testId < 0 {
			util.DataResponse(w, http.StatusNotFound, "Test not found", s.logger)
			return
		}

		test, err := s.testService.GetById(r.Context(), testId)

		if err != nil || test == nil {
			util.DataResponse(w, http.StatusNotFound, "Problem not found", s.logger)
			return
		}

		ctx := context.WithValue(r.Context(), util.TestContextKey, test)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// SubmissionCtx middleware is used to attach the file-managers from the URL parameters to the request context.
func (s *API) SubmissionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		submissionId, err := strconv.Atoi(chi.URLParam(r, "submissionId"))

		if err != nil {
			util.ErrorResponse(w, http.StatusBadRequest, "Invalid file-managers id", s.logger)
			return
		}

		if submissionId < 0 {
			util.DataResponse(w, http.StatusNotFound, "Submission not found", s.logger)
			return
		}

		submission, err := s.submissionService.GetById(r.Context(), submissionId)
		if err != nil || submission == nil {
			util.DataResponse(w, http.StatusNotFound, "Submission not found", s.logger)
			return
		}

		ctx := context.WithValue(r.Context(), util.SubmissionContextKey, submission)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
