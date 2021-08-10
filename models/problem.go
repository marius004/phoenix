package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ProblemDifficulty string

const (
	Easy    ProblemDifficulty = "easy"
	Medium  ProblemDifficulty = "medium"
	Hard    ProblemDifficulty = "hard"
	Contest ProblemDifficulty = "contest"
)

// ex. clasa a 9, clasa a 10 etc
type ProblemGrade string

const (
	NINE   ProblemGrade = "9"
	TEN    ProblemGrade = "10"
	ELEVEN ProblemGrade = "11"
)

type StreamFlag string

const (
	CONSOLE StreamFlag = "console"
	FILE    StreamFlag = "file"
)

var (
	problemNameValidation             = []validation.Rule{validation.Required, validation.Length(3, 20)}
	problemDescriptionValidation      = []validation.Rule{validation.Required, validation.Length(10, 1<<20)}
	problemShortDescriptionValidation = []validation.Rule{validation.Required, validation.Length(10, 2000)}
	problemVisibilityValidation       = []validation.Rule{validation.Required}
	problemDifficultyValidation       = []validation.Rule{validation.Required, validation.In(Easy, Medium, Hard, Contest)}
	problemGradeValidation            = []validation.Rule{validation.Required, validation.In(NINE, TEN, ELEVEN)}
	problemTimeLimitValidation        = []validation.Rule{validation.Required, validation.Min(0.0), validation.Max(2.0)}
	problemMemoryLimitValidation      = []validation.Rule{validation.Required, validation.Min(0), validation.Max(65537)}
	problemStackLimitValidation       = []validation.Rule{validation.Required, validation.Min(0), validation.Max(16384)}
	problemSourceSizeValidation       = []validation.Rule{validation.Required, validation.Min(0), validation.Max(10000)}
	problemStreamValidation           = []validation.Rule{validation.In(CONSOLE, FILE)}
	problemCreditsValidation          = []validation.Rule{validation.Length(0, 100)}
)

type Problem struct {
	Id        uint64     `json:"id"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`

	Name             string            `json:"name" db:"name"`
	Description      string            `json:"description" db:"description"`
	ShortDescription string            `json:"shortDescription" db:"short_description"`
	AuthorId         uint64            `json:"authorId" db:"author_id"`
	Visible          bool              `json:"visible" db:"visible"`
	Difficulty       ProblemDifficulty `json:"difficulty" db:"difficulty"`
	Grade            ProblemGrade      `json:"grade" db:"grade"`
	Credits          string            `json:"credits" db:"credits"`
	Stream           StreamFlag        `json:"stream" db:"stream"`

	TimeLimit   float64 `json:"timeLimit" db:"time_limit"`
	MemoryLimit int     `json:"memoryLimit" db:"memory_limit"`
	StackLimit  int     `json:"stackLimit" db:"stack_limit"`
	SourceSize  int     `json:"sourceSize" db:"source_size"`
}

func (p Problem) IsConsoleProblem() bool {
	return p.Stream == CONSOLE
}

type ProblemFilter struct {
	AuthorsId    []int
	Difficulties []string
	Credits      []string
	Stream       []string
	Grades       []string
}

type UpdateProblemRequest struct {
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	ShortDescription string            `json:"shortDescription"`
	AuthorId         uint64            `json:"authorId"`
	Visible          bool              `json:"visible"`
	Difficulty       ProblemDifficulty `json:"difficulty"`
	Grade            ProblemGrade      `json:"grade"`
	Credits          string            `json:"credits"`
	Stream           StreamFlag        `json:"stream"` // TODO the validation is broken

	TimeLimit   float64 `json:"timeLimit"`
	MemoryLimit int     `json:"memoryLimit"`
	StackLimit  int     `json:"stackLimit"`
	SourceSize  int     `json:"sourceSize"`
}

type CreateProblemRequest struct {
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	ShortDescription string            `json:"shortDescription"`
	Visible          bool              `json:"visible"`
	Difficulty       ProblemDifficulty `json:"difficulty"`
	Grade            ProblemGrade      `json:"grade"`
	Credits          string            `json:"credits"`
	Stream           StreamFlag        `json:"stream"`

	TimeLimit   float64 `json:"timeLimit"`
	MemoryLimit int     `json:"memoryLimit"`
	StackLimit  int     `json:"stackLimit"`
	SourceSize  int     `json:"sourceSize"`
}

func (data CreateProblemRequest) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Name, problemNameValidation...),
		validation.Field(&data.Description, problemDescriptionValidation...),
		validation.Field(&data.ShortDescription, problemShortDescriptionValidation...),
		validation.Field(&data.Visible, problemVisibilityValidation...),
		validation.Field(&data.Difficulty, problemDifficultyValidation...),
		validation.Field(&data.Grade, problemGradeValidation...),
		validation.Field(&data.TimeLimit, problemTimeLimitValidation...),
		validation.Field(&data.MemoryLimit, problemMemoryLimitValidation...),
		validation.Field(&data.SourceSize, problemSourceSizeValidation...),
		validation.Field(&data.StackLimit, problemStackLimitValidation...),
		validation.Field(&data.Credits, problemCreditsValidation...),
		validation.Field(&data.Stream, problemStreamValidation...),
	)
}

func NewProblem(request CreateProblemRequest) *Problem {
	return &Problem{
		Name:             request.Name,
		Description:      request.Description,
		ShortDescription: request.ShortDescription,
		Visible:          request.Visible,
		Difficulty:       request.Difficulty,
		Grade:            request.Grade,
		Credits:          request.Credits,
		Stream:           request.Stream,

		TimeLimit:   request.TimeLimit,
		MemoryLimit: request.MemoryLimit,
		StackLimit:  request.StackLimit,
		SourceSize:  request.SourceSize,
	}
}
