package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/marius004/phoenix/models"
)

type BlogPostService struct {
	db     *sqlx.DB
	logger *log.Logger
}

const createBlogPostQuery = `INSERT INTO blog_posts(title, message, author_id) VALUES (?,?,?) RETURNING id`

func (s *BlogPostService) GetById(context context.Context, id int) (*models.BlogPost, error) {
	var post models.BlogPost
	err := s.db.GetContext(context, &post, s.db.Rebind("SELECT * FROM blog_posts where id = ? LIMIT 1"), id)

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return &post, err
}

func (s *BlogPostService) Create(context context.Context, post *models.BlogPost) error {
	var id uint64

	err := s.db.GetContext(context, &id, s.db.Rebind(createBlogPostQuery),
		post.Title, post.Message, post.AuthorId)

	if err == nil {
		post.Id = id
	} else {
		s.logger.Println(err)
	}

	return err
}

func (s *BlogPostService) UpdateById(context context.Context, id int, update *models.UpdateBlogPost) error {
	queryList, args := s.updateQueryMaker(update)

	if len(queryList) == 0 {
		return nil
	}

	args = append(args, id)
	query := s.db.Rebind(fmt.Sprintf("UPDATE blog_posts SET %s WHERE id = ?", strings.Join(queryList, ", ")))

	_, err := s.db.ExecContext(context, query, args...)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *BlogPostService) DeleteById(context context.Context, id int) error {
	query := s.db.Rebind("DELETE FROM blog_posts WHERE id = ?")
	_, err := s.db.ExecContext(context, query, id)

	if err != nil {
		s.logger.Println(err)
	}

	return err
}

func (s *BlogPostService) updateQueryMaker(r *models.UpdateBlogPost) ([]string, []interface{}) {
	var query []string
	var args []interface{}

	if v := r.Message; v != "" {
		query, args = append(query, "message = ?"), append(args, v)
	}

	if v := r.Title; v != "" {
		query, args = append(query, "title = ?"), append(args, v)
	}

	return query, args
}

func NewBlogPostService(db *sqlx.DB, logger *log.Logger) *BlogPostService {
	return &BlogPostService{db, logger}
}
