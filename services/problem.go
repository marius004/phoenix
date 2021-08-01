package services

import (
	"context"

	"github.com/marius004/phoenix/models"
)

type ProblemService interface {
	// GetByFilter returns a list of problems after applying the filter provided
	GetByFilter(ctx context.Context, filter *models.ProblemFilter) ([]*models.Problem, error)

	// GetById looks up a problem by id.
	GetById(ctx context.Context, id int) (*models.Problem, error)

	// GetByName looks up a problem by its name
	GetByName(ctx context.Context, name string) (*models.Problem, error)

	// GetAll retrieves all problems
	GetAll(ctx context.Context) ([]*models.Problem, error)

	// Create creates a new Problem
	Create(ctx context.Context, problem *models.Problem, authorId int) error

	// Update updates the problem with the given id.s
	Update(ctx context.Context, id int, updateRequest *models.UpdateProblemRequest) error

	// Delete deletes the problem with the given id
	Delete(ctx context.Context, id int) error

	// ExistsByName checks if there is a problem with the given name
	ExistsByName(ctx context.Context, name string) (bool, error)

	// ExistsById checks if there is a problem with the given id
	ExistsById(ctx context.Context, id int) (bool, error)
}
