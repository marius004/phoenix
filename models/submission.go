package models

import (
	"net/http"
	"strconv"
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
}

func NewSubmission(request CreateSubmissionRequest, userId int) *Submission {
	return &Submission{
		Lang:       request.Lang,
		ProblemId:  request.ProblemId,
		UserId:     userId,
		SourceCode: string(request.SourceCode),
	}
}

func ParseSubmissionFilter(r *http.Request) *SubmissionFilter {
	ret := SubmissionFilter{}

	if v, ok := r.URL.Query()["userId"]; ok {
		last := len(v) - 1
		if val, err := strconv.Atoi(v[last]); err == nil {
			ret.UserId = val
		}
	}

	if v, ok := r.URL.Query()["lang"]; ok {
		ret.Langs = convertUrlValuesToLangArr(v)
	}

	if v, ok := r.URL.Query()["problemId"]; ok {
		last := len(v) - 1
		if val, err := strconv.Atoi(v[last]); err == nil {
			ret.ProblemId = val
		}
	}

	if v, ok := r.URL.Query()["score"]; ok {
		last := len(v) - 1
		if val, err := strconv.Atoi(v[last]); err == nil {
			ret.Score = val
		}
	} else {
		ret.Score = -1
	}

	if v, ok := r.URL.Query()["status"]; ok {
		ret.Statuses = convertUrlValuesToStatusArr(v)
	}

	if v, ok := r.URL.Query()["compileError"]; ok {
		last := len(v) - 1
		if val, err := strconv.ParseBool(v[last]); err == nil {
			ret.CompileError = &val
		}
	}

	return &ret
}
