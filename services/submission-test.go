package services

import (
	"context"

	"github.com/marius004/phoenix/models"
)

type SubmissionTestService interface {
	GetBySubmissionId(ctx context.Context, submissionId uint64) ([]*models.SubmissionTest, error)

	GetById(ctx context.Context, id uint64) (*models.SubmissionTest, error)

	Create(ctx context.Context, submissionTest *models.SubmissionTest) error

	Update(ctx context.Context, submissionId, testId uint64, submissionTest *models.UpdateSubmissionTestRequest) error
}
