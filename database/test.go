package database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/marius004/phoenix/models"
	"github.com/jmoiron/sqlx"
)

// TestService implements services.TestService
type TestService struct {
	db     *sqlx.DB
	logger *log.Logger
}

func (s *TestService) Create(ctx context.Context, test *models.Test) error {
	var id uint64
	query := "INSERT INTO tests(problem_id, score) VALUES(?,?) RETURNING id"

	err := s.db.GetContext(ctx, &id, s.db.Rebind(query), test.ProblemId, test.Score)
	if err == nil {
		test.Id = id
	}

	return err
}

func (s *TestService) GetById(ctx context.Context, id int) (*models.Test, error) {
	var test models.Test
	err := s.db.GetContext(ctx, &test, s.db.Rebind("SELECT * FROM tests where id = ? LIMIT 1"), id)

	if err != nil {
		return nil, err
	}

	return &test, err
}

func (s *TestService) GetAllProblemTests(ctx context.Context, problemId int) ([]*models.Test, error) {
	var tests []*models.Test
	err := s.db.Select(&tests, s.db.Rebind("SELECT * FROM tests WHERE problem_id = ?"), problemId)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return tests, err
}

func (s *TestService) GetAllTests(ctx context.Context) ([]*models.Test, error) {
	var tests []*models.Test
	err := s.db.Select(&tests, "SELECT * FROM tests")

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return tests, err
}

func (s *TestService) Update(ctx context.Context, testId, problemId int, test *models.Test) error {
	query := s.db.Rebind("UPDATE tests SET score = ? WHERE id = ? AND problem_id = ?")

	_, err := s.db.ExecContext(ctx, query, test.Score, testId, problemId)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *TestService) Delete(ctx context.Context, testId, problemId int) error {
	query := s.db.Rebind("DELETE FROM tests WHERE id = ? AND problem_id = ?")
	_, err := s.db.ExecContext(ctx, query, testId, problemId)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *TestService) DeleteAllProblemTests(ctx context.Context, problemId int) error {
	query := s.db.Rebind("DELETE FROM tests WHERE problem_id = ?")
	_, err := s.db.ExecContext(ctx, query, problemId)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func NewTestService(db *sqlx.DB, logger *log.Logger) *TestService {
	return &TestService{db, logger}
}
