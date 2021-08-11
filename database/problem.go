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

// ProblemService implements services.ProblemService
type ProblemService struct {
	db     *sqlx.DB
	logger *log.Logger
}

const (
	createProblemQuery = `INSERT INTO problems(name, description, short_description, author_id, visible,
                     		difficulty, grade, time_limit, memory_limit, stack_limit, source_size, credits, stream) 
							VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?) RETURNING id`
)

func (s *ProblemService) GetByFilter(ctx context.Context, filter *models.ProblemFilter) ([]*models.Problem, error) {
	var problems []*models.Problem
	queryList, args := s.filterMaker(filter)

	if len(queryList) == 0 {
		if err := s.db.Select(&problems, "SELECT * FROM problems"); err != nil {
			return nil, err
		}
		return problems, nil
	}

	query := s.db.Rebind(fmt.Sprintf("SELECT * FROM problems WHERE %s", strings.Join(queryList, " AND ")))
	if err := s.db.Select(&problems, query, args...); err != nil {
		return nil, err
	}

	return problems, nil
}

func (s *ProblemService) GetById(ctx context.Context, id int) (*models.Problem, error) {
	var problem models.Problem
	err := s.db.GetContext(ctx, &problem, s.db.Rebind("SELECT * FROM problems where id = ? LIMIT 1"), id)

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return &problem, err
}

func (s *ProblemService) GetAll(ctx context.Context) ([]*models.Problem, error) {
	var problems []*models.Problem
	err := s.db.Select(&problems, "SELECT * FROM problems")

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return problems, err
}

func (s *ProblemService) GetByName(ctx context.Context, name string) (*models.Problem, error) {
	var problem models.Problem
	err := s.db.GetContext(ctx, &problem, s.db.Rebind("SELECT * FROM problems WHERE name = ? LIMIT 1"), name)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		s.logger.Println(err)
		return nil, err
	}

	return &problem, err
}

func (s *ProblemService) Create(ctx context.Context, problem *models.Problem, authorId int) error {
	var id uint64

	err := s.db.GetContext(ctx, &id, s.db.Rebind(createProblemQuery),
		problem.Name, problem.Description, problem.ShortDescription,
		authorId, problem.Visible, problem.Difficulty,
		problem.Grade, problem.TimeLimit, problem.MemoryLimit,
		problem.StackLimit, problem.SourceSize, problem.Credits, problem.Stream)

	if err == nil {
		problem.Id = id
	} else {
		s.logger.Println(err)
	}

	return err
}

func (s *ProblemService) Delete(ctx context.Context, id int) error {
	query := s.db.Rebind("DELETE FROM problems WHERE id = ?")
	_, err := s.db.ExecContext(ctx, query, id)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *ProblemService) Update(ctx context.Context, id int, updateRequest *models.UpdateProblemRequest) error {
	queryList, args := s.updateQueryMaker(updateRequest)

	if len(queryList) == 0 {
		return nil
	}

	args = append(args, id)
	query := s.db.Rebind(fmt.Sprintf("UPDATE problems SET %s WHERE id = ?", strings.Join(queryList, ", ")))

	_, err := s.db.ExecContext(ctx, query, args...)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *ProblemService) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int

	query := "SELECT COUNT(*) FROM problems WHERE name = ?"
	err := s.db.GetContext(ctx, &count, s.db.Rebind(query), name)

	if err != nil {
		s.logger.Println(err)
	}

	return count > 0, err
}

func (s *ProblemService) ExistsById(ctx context.Context, id int) (bool, error) {
	var count int

	query := "SELECT COUNT(*) FROM problems WHERE id = ?"
	err := s.db.GetContext(ctx, &count, s.db.Rebind(query), id)

	if err != nil {
		s.logger.Println(err)
	}

	return count > 0, err
}

// RESTRUCTURE THIS PIECE OF CODE ACCORDING TO DRY!!
func (s *ProblemService) filterMaker(filter *models.ProblemFilter) ([]string, []interface{}) {
	var query []string
	var args []interface{}

	if authors := filter.AuthorsId; len(authors) > 0 {

		str := ""

		for ind, val := range authors {
			str += "author_id = ?"
			args = append(args, val)

			if ind != len(authors)-1 {
				str += " OR "
			}
		}

		query = append(query, str)
	}

	if credits := filter.Credits; len(credits) > 0 {
		str := ""

		for ind, val := range credits {
			str += "credits = ?"
			args = append(args, val)

			if ind != len(credits)-1 {
				str += " OR "
			}
		}

		query = append(query, str)
	}

	if difficulties := filter.Difficulties; len(difficulties) > 0 {
		str := ""

		for ind, val := range difficulties {
			str += "difficulty = ?"
			args = append(args, val)

			if ind != len(difficulties)-1 {
				str += " OR "
			}
		}

		query = append(query, str)
	}

	if streams := filter.Stream; len(streams) > 0 {
		str := ""

		for ind, val := range streams {
			str += "stream = ?"
			args = append(args, val)

			if ind != len(streams)-1 {
				str += " OR "
			}
		}

		query = append(query, str)
	}

	if grades := filter.Grades; len(grades) > 0 {
		str := ""

		for ind, val := range grades {
			str += "grade = ?"
			args = append(args, val)

			if ind != len(grades)-1 {
				str += " OR "
			}
		}

		query = append(query, str)
	}

	return query, args
}

func (s *ProblemService) updateQueryMaker(r *models.UpdateProblemRequest) ([]string, []interface{}) {
	var query []string
	var args []interface{}

	if v := r.Name; v != "" {
		query, args = append(query, "name = ?"), append(args, v)
	}

	if v := r.Description; v != "" {
		query, args = append(query, "description = ?"), append(args, v)
	}

	if v := r.ShortDescription; v != "" {
		query, args = append(query, "short_description = ?"), append(args, v)
	}

	if v := r.AuthorId; v > 0 {
		query, args = append(query, "author_id = ?"), append(args, v)
	}

	if v := r.Visible; true {
		query, args = append(query, "visible = ?"), append(args, v)
	}

	if v := r.Difficulty; v != "" {
		query, args = append(query, "difficulty = ?"), append(args, v)
	}

	if v := r.Grade; v != "" {
		query, args = append(query, "grade = ?"), append(args, v)
	}

	if v := r.TimeLimit; v > 0 {
		query, args = append(query, "time_limit = ?"), append(args, v)
	}

	if v := r.MemoryLimit; v > 0 {
		query, args = append(query, "memory_limit = ?"), append(args, v)
	}

	if v := r.StackLimit; v > 0 {
		query, args = append(query, "stack_limit = ?"), append(args, v)
	}

	if v := r.SourceSize; v > 0 {
		query, args = append(query, "source_size = ?"), append(args, v)
	}

	if v := r.Stream; v != "" {
		query, args = append(query, "stream = ?"), append(args, v)
	}

	if v := r.Credits; v != "" {
		query, args = append(query, "credits = ?"), append(args, v)
	}

	return query, args
}

func NewProblemService(db *sqlx.DB, logger *log.Logger) *ProblemService {
	return &ProblemService{db, logger}
}
