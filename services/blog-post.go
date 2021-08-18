package services

import (
	"context"

	"github.com/marius004/phoenix/models"
)

type BlogPostService interface {
	GetById(context context.Context, id int) (*models.BlogPost, error)

	Create(context context.Context, post *models.BlogPost) error

	UpdateById(context context.Context, id int, update *models.UpdateBlogPost) error

	DeleteById(context context.Context, id int) error
}
