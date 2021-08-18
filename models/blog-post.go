package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	postTitleValidation   = []validation.Rule{validation.Required, validation.Length(4, 0)}
	postMessageValidation = []validation.Rule{validation.Required, validation.Length(4, 0)}
)

type BlogPost struct {
	Id        uint64     `json:"id"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`

	Title   string `json:"title"`
	Message string `json:"message"`

	AuthorId uint64 `json:"authorId" db:"author_id"`
}

type UpdateBlogPost struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func (data UpdateBlogPost) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Message, postMessageValidation...),
		validation.Field(&data.Title, postTitleValidation...),
	)
}

type CreateBlogPost struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func (data CreateBlogPost) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Message, postMessageValidation...),
		validation.Field(&data.Title, postTitleValidation...),
	)
}

func NewBlogPost(req *CreateBlogPost, authorId uint64) *BlogPost {
	return &BlogPost{
		Title:    req.Title,
		Message:  req.Message,
		AuthorId: authorId,
	}
}
