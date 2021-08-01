package services

import (
	"context"

	"github.com/marius004/phoenix/models"
)

// SubmissionService is a service for managing submissions
type SubmissionService interface {
	// GetByFilter returns a list of submissions after applying the filter provided
	GetByFilter(ctx context.Context, filter *models.SubmissionFilter) ([]*models.Submission, error)

	// GetById looks up a file-managers by id.
	GetById(ctx context.Context, id int) (*models.Submission, error)

	// GetByUserName looks up a file-managers by its name
	GetByUserName(ctx context.Context, userId int) ([]*models.Submission, error)

	// GetAll retrieves all submissions
	GetAll(ctx context.Context) ([]*models.Submission, error)

	// Create creates a new file-managers
	Create(ctx context.Context, submission *models.Submission) error

	// Delete deletes the file-managers with the given id
	Delete(ctx context.Context, id int) error

	// Update updates the file-managers with given id
	Update(ctx context.Context, id int, updateRequest *models.UpdateSubmissionRequest) error
}
