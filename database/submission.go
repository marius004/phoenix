package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/marius004/phoenix/internal/models"
)

// SubmissionService implements services.SubmissionService
type SubmissionService struct {
	db     *sqlx.DB
	logger *log.Logger
}

const (
	createSubmissionQuery = `INSERT INTO submissions(lang, problem_id, user_id, source_code)
							 VALUES(?,?,?,?) RETURNING id`
)

func (s *SubmissionService) GetByFilter(ctx context.Context, filter *models.SubmissionFilter) ([]*models.Submission, error) {
	var submissions []*models.Submission
	var query string

	queryList, args := s.filterMaker(filter)

	if len(queryList) == 0 {
		query = s.db.Rebind(fmt.Sprintf("SELECT * FROM submissions ORDER BY id DESC LIMIT %d OFFSET %d", filter.Limit, filter.Offset))
	} else {
		query = s.db.Rebind(fmt.Sprintf("SELECT * FROM submissions WHERE %s ORDER BY id DESC LIMIT %d OFFSET %d", strings.Join(queryList, " AND "), filter.Limit, filter.Offset))
	}

	if err := s.db.Select(&submissions, query, args...); err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return submissions, nil
}

func (s *SubmissionService) GetById(ctx context.Context, id int) (*models.Submission, error) {
	var submission models.Submission
	err := s.db.GetContext(ctx, &submission, s.db.Rebind("SELECT * FROM submissions where id = ? LIMIT 1"), id)

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return &submission, err
}

func (s *SubmissionService) GetByUserName(ctx context.Context, userId int) ([]*models.Submission, error) {
	var submission []*models.Submission
	err := s.db.GetContext(ctx, &submission, s.db.Rebind("SELECT * FROM submissions where userId = ? LIMIT 1"), userId)

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return submission, err
}

func (s *SubmissionService) GetAll(ctx context.Context) ([]*models.Submission, error) {
	var submissions []*models.Submission
	err := s.db.Select(&submissions, "SELECT * FROM submissions")

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return submissions, err
}

func (s *SubmissionService) Create(ctx context.Context, submission *models.Submission) error {
	var id uint64

	err := s.db.GetContext(ctx, &id, s.db.Rebind(createSubmissionQuery),
		submission.Lang, submission.ProblemId, submission.UserId, submission.SourceCode)

	if err == nil {
		submission.Id = id
	} else {
		s.logger.Println(err)
	}

	return err
}

func (s *SubmissionService) Delete(ctx context.Context, id int) error {
	query := s.db.Rebind("DELETE FROM submissions WHERE id = ?")
	_, err := s.db.ExecContext(ctx, query, id)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *SubmissionService) Update(ctx context.Context, id int, updateRequest *models.UpdateSubmissionRequest) error {
	queryList, args := s.updateQueryMaker(updateRequest)

	if len(queryList) == 0 {
		return nil
	}

	args = append(args, id)
	query := s.db.Rebind(fmt.Sprintf("UPDATE submissions SET %s WHERE id = ?", strings.Join(queryList, ", ")))

	_, err := s.db.ExecContext(ctx, query, args...)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *SubmissionService) filterMaker(filter *models.SubmissionFilter) ([]string, []interface{}) {
	var query []string
	var args []interface{}

	if v := filter.UserId; v != 0 {
		query, args = append(query, "user_id = ?"), append(args, v)
	}

	if v := filter.ProblemId; v != 0 {
		query, args = append(query, "problem_id = ?"), append(args, v)
	}

	if v := filter.Score; v >= 0 {
		query, args = append(query, "score = ?"), append(args, v)
	}

	if langs := filter.Langs; len(langs) > 0 {
		str := ""

		for ind, val := range langs {
			str += "lang = ?"
			args = append(args, val)

			if ind != len(langs)-1 {
				str += " OR "
			}
		}

		query = append(query, str)
	}

	if statuses := filter.Statuses; len(statuses) > 0 {
		str := ""

		for ind, val := range statuses {
			str += "status = ?"
			args = append(args, val)

			if ind != len(statuses)-1 {
				str += " OR "
			}
		}

		query = append(query, str)
	}

	if v := filter.CompileError; v != nil {
		query, args = append(query, "has_compile_error = ?"), append(args, v)
	}

	if filter.Limit <= 0 {
		filter.Limit = (1 << 30)
	}

	if filter.Offset <= 0 {
		filter.Offset = 0
	}

	return query, args
}

func (s *SubmissionService) updateQueryMaker(r *models.UpdateSubmissionRequest) ([]string, []interface{}) {
	var query []string
	var args []interface{}

	if v := r.Score; v > 0 {
		query, args = append(query, "score = ?"), append(args, v)
	}

	if v := r.Status; v != "" {
		query, args = append(query, "status = ?"), append(args, v)
	}

	if v := r.Message; v != "" {
		query, args = append(query, "message = ?"), append(args, v)
	}

	if v := r.HasCompileError; v != nil {
		query, args = append(query, "has_compile_error = ?"), append(args, v)
	}

	return query, args
}

func NewSubmissionService(db *sqlx.DB, logger *log.Logger) *SubmissionService {
	return &SubmissionService{db, logger}
}
