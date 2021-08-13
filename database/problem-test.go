package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/marius004/phoenix/models"
)

// TestService implements services.TestService
type TestService struct {
	db     *sqlx.DB
	logger *log.Logger
}

func (s *TestService) Create(ctx context.Context, test *models.Test) error {
	var id uint64
	query := "INSERT INTO problem_tests(problem_id, score) VALUES(?,?) RETURNING id"

	err := s.db.GetContext(ctx, &id, s.db.Rebind(query), test.ProblemId, test.Score)
	if err == nil {
		test.Id = id
	}

	return err
}

func (s *TestService) GetById(ctx context.Context, id int) (*models.Test, error) {
	var test models.Test
	err := s.db.GetContext(ctx, &test, s.db.Rebind("SELECT * FROM problem_tests where id = ? LIMIT 1"), id)

	if err != nil {
		return nil, err
	}

	return &test, err
}

func (s *TestService) GetAllProblemTests(ctx context.Context, problemId int) ([]*models.Test, error) {
	var tests []*models.Test
	err := s.db.Select(&tests, s.db.Rebind("SELECT * FROM problem_tests WHERE problem_id = ?"), problemId)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return tests, err
}

func (s *TestService) GetAllTests(ctx context.Context, filter *models.TestFilter) ([]*models.Test, error) {
	var tests []*models.Test
	var query string

	queryList, args := s.filterMaker(filter)

	if len(queryList) == 0 {
		query = "SELECT * FROM problem_tests"
	} else {
		query = s.db.Rebind(fmt.Sprintf("SELECT * FROM problem_tests WHERE %s ORDER BY id DESC", strings.Join(queryList, " AND ")))
	}

	if err := s.db.Select(&tests, query, args...); err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return tests, nil
}

func (s *TestService) Update(ctx context.Context, testId, problemId int, test *models.Test) error {
	query := s.db.Rebind("UPDATE problem_tests SET score = ? WHERE id = ? AND problem_id = ?")

	_, err := s.db.ExecContext(ctx, query, test.Score, testId, problemId)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *TestService) Delete(ctx context.Context, testId, problemId int) error {
	query := s.db.Rebind("DELETE FROM problem_tests WHERE id = ? AND problem_id = ?")
	_, err := s.db.ExecContext(ctx, query, testId, problemId)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *TestService) DeleteAllProblemTests(ctx context.Context, problemId int) error {
	query := s.db.Rebind("DELETE FROM problem_tests WHERE problem_id = ?")
	_, err := s.db.ExecContext(ctx, query, problemId)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func NewTestService(db *sqlx.DB, logger *log.Logger) *TestService {
	return &TestService{db, logger}
}

func (s *TestService) filterMaker(filter *models.TestFilter) ([]string, []interface{}) {
	var query []string
	var args []interface{}

	if v := filter.ProblemId; v > 0 {
		query, args = append(query, "problem_id = ?"), append(args, v)
	}

	if v := filter.Score; v > 0 {
		query, args = append(query, "score = ?"), append(args, v)
	}

	return query, args
}
