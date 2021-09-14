package util

import (
	"context"

	"github.com/marius004/phoenix/internal/models"
)

type ContextType string

const (
	UserContextKey     = ContextType("user")
	ProblemContextKey  = ContextType("problem")
	TestContextKey     = ContextType("test")
	BlogPostContextKey = ContextType("blog-post")

	SubmissionContextKey     = ContextType("submission")
	SubmissionTestContextKey = ContextType("submissionTest")
)

// UserFromRequestContext returns a pointer to the user from request context
func UserFromRequestContext(context context.Context) *models.User {
	switch usr := context.Value(UserContextKey).(type) {
	case models.User:
		return &usr
	case *models.User:
		return usr
	default:
		return nil
	}
}

// ProblemFromRequestContext returns a pointer to the problem from request context
func ProblemFromRequestContext(context context.Context) *models.Problem {
	switch prb := context.Value(ProblemContextKey).(type) {
	case models.Problem:
		return &prb
	case *models.Problem:
		return prb
	default:
		return nil
	}
}

// TestFromRequestContext returns a pointer to the test from request context
func TestFromRequestContext(context context.Context) *models.Test {
	switch test := context.Value(TestContextKey).(type) {
	case models.Test:
		return &test
	case *models.Test:
		return test
	default:
		return nil
	}
}

// SubmissionFromRequestContext returns a pointer to the file-managers from request context
func SubmissionFromRequestContext(context context.Context) *models.Submission {
	switch submission := context.Value(SubmissionContextKey).(type) {
	case models.Submission:
		return &submission
	case *models.Submission:
		return submission
	default:
		return nil
	}
}

func SubmissionTestFromRequestContext(context context.Context) *models.SubmissionTest {
	switch submissionTest := context.Value(SubmissionTestContextKey).(type) {
	case models.SubmissionTest:
		return &submissionTest
	case *models.SubmissionTest:
		return submissionTest
	default:
		return nil
	}
}
