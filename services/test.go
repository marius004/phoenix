package services

import (
	"context"

	"github.com/marius004/phoenix/models"
)

// TestService is a service for managing tests.
// It does not actually store the real tests(that is the job of TestManger).
// It stores basic information about a test such as when it was created and its score.
type TestService interface {

	// Create creates a new test.
	Create(ctx context.Context, test *models.Test) error

	// GetById retrieves the test matching the specified id.
	GetById(ctx context.Context, id int) (*models.Test, error)

	// GetAllProblemTests retrieves all the tests for the specified problem
	GetAllProblemTests(ctx context.Context, problemId int) ([]*models.Test, error)

	// Update updates the test for a given problem
	Update(ctx context.Context, testId, problemId int, test *models.Test) error

	// Delete deletes the test for a given problem
	Delete(ctx context.Context, testId, problemId int) error

	// DeleteAllProblemTests deletes all the tests for the specified problem
	DeleteAllProblemTests(ctx context.Context, problemId int) error
}
