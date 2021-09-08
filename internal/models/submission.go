package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type SubmissionLang string

const C = SubmissionLang("c")

type SubmissionStatus string

const (
	Waiting  = SubmissionStatus("waiting")
	Working  = SubmissionStatus("working")
	Finished = SubmissionStatus("finished")
)

type Submission struct {
	Id        uint64     `json:"id"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`

	Score           int              `json:"score"`
	Lang            SubmissionLang   `json:"lang"`
	Status          SubmissionStatus `json:"status"`
	Message         string           `json:"message"`
	HasCompileError *bool            `json:"hasCompileError" db:"has_compile_error"`

	ProblemId  int    `json:"problemId" db:"problem_id"`
	UserId     int    `json:"userId" db:"user_id"`
	SourceCode string `json:"sourceCode" db:"source_code"`

	// I know it is not a good practice to do something like this
	// but I hope I will restructure this in the future
	Username    string `json:"username"`
	EmailHash   string `json:"emailHash"`
	ProblemName string `json:"problemName"`
}

var (
	submissionLangValidation       = []validation.Rule{validation.Required, validation.In(C)}
	submissionProblemIdValidation  = []validation.Rule{validation.Required, validation.Min(0)}
	submissionSourceCodeValidation = []validation.Rule{validation.Required, validation.Length(1, 0)}
)

type UpdateSubmissionRequest struct {
	Score           int
	Status          SubmissionStatus
	Message         string
	HasCompileError *bool
}

func NewUpdateSubmissionRequest(score int, status SubmissionStatus, message string, compileError *bool) *UpdateSubmissionRequest {
	return &UpdateSubmissionRequest{
		Score:           score,
		Status:          status,
		Message:         message,
		HasCompileError: compileError,
	}
}

type CreateSubmissionRequest struct {
	Lang       SubmissionLang `json:"lang"`
	ProblemId  int            `json:"problemId"`
	SourceCode string         `json:"sourceCode"`
}

func (data CreateSubmissionRequest) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Lang, submissionLangValidation...),
		validation.Field(&data.ProblemId, submissionProblemIdValidation...),
		validation.Field(&data.SourceCode, submissionSourceCodeValidation...),
	)
}

type SubmissionFilter struct {
	UserId       int
	ProblemId    int
	Score        int
	Langs        []SubmissionLang
	Statuses     []SubmissionStatus
	CompileError *bool // nil == skip(does not matter), false no compile error, true compile errors

	Limit  int
	Offset int
}

func NewSubmission(request CreateSubmissionRequest, userId int) *Submission {
	return &Submission{
		Lang:       request.Lang,
		ProblemId:  request.ProblemId,
		UserId:     userId,
		SourceCode: string(request.SourceCode),
	}
}
