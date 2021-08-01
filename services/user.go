package services

import (
	"context"

	"github.com/marius004/phoenix/models"
)

// UserService is a service for managing users
type UserService interface {

	// GetById looks up a user by id. Returns an error if the user is not found
	GetById(ctx context.Context, id int) (*models.User, error)

	// GetByEmail looks up a user by email.
	GetByEmail(ctx context.Context, email string) (*models.User, error)

	// GetByUsername looks up a user by username.
	GetByUsername(ctx context.Context, username string) (*models.User, error)

	// GetAll retrieves all users.
	GetAll(ctx context.Context) ([]*models.User, error)

	// Create creates a new user entity
	Create(ctx context.Context, user *models.User) error

	// ExistsById checks if the user with the specified id exits in the database
	ExistsById(ctx context.Context, id int) (bool, error)

	// ExistsByUsername checks if the user with the specified username exits in the database
	ExistsByUsername(ctx context.Context, username string) (bool, error)

	// ExistsByEmail checks if the user with the specified email exits in the database
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}
