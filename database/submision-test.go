package database

import (
	"context"
	"fmt"
	"log"
	"strings"

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
	err := s.db.Select(&submissionTests, s.db.Rebind(query), submissionId)

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

func (s *SubmissionTestService) GetBySubmissionAndTestId(ctx context.Context, submissionId, testId uint64) (*models.SubmissionTest, error) {
	var submissionTest models.SubmissionTest
	err := s.db.GetContext(ctx, &submissionTest, s.db.Rebind("SELECT * FROM submission_tests WHERE submission_id = ? AND test_id = ? LIMIT 1"), submissionId, testId)

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

func (s *SubmissionTestService) Update(ctx context.Context, submissionId, testId uint64, update *models.UpdateSubmissionTestRequest) error {
	queryList, args := s.updateQueryMaker(update)

	if len(queryList) == 0 {
		return nil
	}

	args = append(args, submissionId)
	args = append(args, testId)

	query := s.db.Rebind(fmt.Sprintf("UPDATE submission_tests SET %s WHERE submission_id = ? AND test_id = ?", strings.Join(queryList, ", ")))

	_, err := s.db.ExecContext(ctx, query, args...)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *SubmissionTestService) updateQueryMaker(r *models.UpdateSubmissionTestRequest) ([]string, []interface{}) {
	var query []string
	var args []interface{}

	if v := r.ExitCode; v != -1 {
		query, args = append(query, "exit_code = ?"), append(args, v)
	}

	if v := r.Memory; v != -1 {
		query, args = append(query, "memory = ?"), append(args, v)
	}

	if v := r.Time; v != -1 {
		query, args = append(query, "time = ?"), append(args, v)
	}

	if v := r.Score; v != -1 {
		query, args = append(query, "score = ?"), append(args, v)
	}

	if v := r.Message; v != "" {
		query, args = append(query, "message = ?"), append(args, v)
	}

	return query, args
}

func NewSubmissionTestService(db *sqlx.DB, logger *log.Logger) *SubmissionTestService {
	return &SubmissionTestService{
		db:     db,
		logger: logger,
	}
}
