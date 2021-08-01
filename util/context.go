package util

import (
	"net/http"

	"github.com/marius004/phoenix/models"
)

type ContextType string

const (
	UserContextKey       = ContextType("user")
	ProblemContextKey    = ContextType("problem")
	TestContextKey       = ContextType("test")
	SubmissionContextKey = ContextType("file-managers")
)

// UserFromRequestContext returns a pointer to the user from request context
func UserFromRequestContext(r *http.Request) *models.User {
	switch usr := r.Context().Value(UserContextKey).(type) {
	case models.User:
		return &usr
	case *models.User:
		return usr
	default:
		return nil
	}
}

// ProblemFromRequestContext returns a pointer to the problem from request context
func ProblemFromRequestContext(r *http.Request) *models.Problem {
	switch prb := r.Context().Value(ProblemContextKey).(type) {
	case models.Problem:
		return &prb
	case *models.Problem:
		return prb
	default:
		return nil
	}
}

// TestFromRequestContext returns a pointer to the test from request context
func TestFromRequestContext(r *http.Request) *models.Test {
	switch test := r.Context().Value(TestContextKey).(type) {
	case models.Test:
		return &test
	case *models.Test:
		return test
	default:
		return nil
	}
}

// SubmissionFromRequestContext returns a pointer to the file-managers from request context
func SubmissionFromRequestContext(r *http.Request) *models.Submission {
	switch submission := r.Context().Value(SubmissionContextKey).(type) {
	case models.Submission:
		return &submission
	case *models.Submission:
		return submission
	default:
		return nil
	}
}
