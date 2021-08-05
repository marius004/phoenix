package database

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/marius004/phoenix/models"
)

type SubmissionTestService struct {
	db     *sqlx.DB
	logger *log.Logger
}

func (s *SubmissionTestService) GetBySubmissionId(ctx context.Context, submissionId uint64) ([]*models.SubmissionTest, error) {
	var submissionTests []*models.SubmissionTest

	query := "SELECT * FROM submission_tests WHERE submission_id = ?"
	err := s.db.GetContext(ctx, &submissionTests, s.db.Rebind(query), submissionId)

	if err != nil {
		return nil, err
	}

	return submissionTests, err
}

func (s *SubmissionTestService) GetById(ctx context.Context, id uint64) (*models.SubmissionTest, error) {
	var submissionTest models.SubmissionTest
	err := s.db.GetContext(ctx, &submissionTest, s.db.Rebind("SELECT * FROM submission_tests where id = ? LIMIT 1"), id)

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return &submissionTest, err
}

func (s *SubmissionTestService) Create(ctx context.Context, tst *models.SubmissionTest) error {
	var id uint64

	query := "INSERT INTO submission_tests(score, time, memory, message, exit_code, submission_id, test_id) VALUES(?,?,?,?,?,?,?) RETURNING id"

	err := s.db.GetContext(ctx, &id, s.db.Rebind(query),
		tst.Score, tst.Time, tst.Memory,
		tst.Message, tst.ExitCode,
		tst.SubmissionId, tst.TestId)

	if err == nil {
		tst.Id = id
	}

	return err
}

func NewSubmissionTestService(db *sqlx.DB, logger *log.Logger) *SubmissionTestService {
	return &SubmissionTestService{
		db:     db,
		logger: logger,
	}
}
