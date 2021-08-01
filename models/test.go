package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	testScoreValidation     = []validation.Rule{validation.Required, validation.Max(100)}
	testProblemIdValidation = []validation.Rule{validation.Required, validation.Min(1)}
	testInputValidation     = []validation.Rule{validation.Required}
	testOutputValidation    = []validation.Rule{validation.Required}
)

type Test struct {
	Id        uint64     `json:"id" db:"id"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`

	ProblemId int `json:"problemId" db:"problem_id"`
	Score     int `json:"score" db:"score"`
}

type CreateTestRequest struct {
	ProblemId int    `json:"problemId"`
	Score     int    `json:"score"`
	Input     []byte `json:"input"`
	Output    []byte `json:"output"`
}

func (data CreateTestRequest) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Score, testScoreValidation...),
		validation.Field(&data.ProblemId, testProblemIdValidation...),
		validation.Field(&data.Input, testInputValidation...),
		validation.Field(&data.Output, testOutputValidation...),
	)
}

func NewTest(request CreateTestRequest) *Test {
	return &Test{
		Score:     request.Score,
		ProblemId: request.ProblemId,
	}
}

// FullTest is an extended Test that has input and output fields.
type FullTest struct {
	Test  		  `json:"test"`
	Input  string `json:"input"`
	Output string `json:"output"`
}

func NewFullTest(test Test, inputUri, outputUri string) *FullTest {
	return &FullTest{
		test,
		inputUri,
		outputUri,
	}
}

// TODO add input and output validation
type UpdateTestRequest struct {
	Score  int    `json:"score"`
	Input  []byte `json:"input"`
	Output []byte `json:"output"`
}

func UpdateTest(test *Test, request *UpdateTestRequest) {
	test.Score = request.Score
}

func (data UpdateTestRequest) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Score, testScoreValidation...),
	)
}
