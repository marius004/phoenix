package models

import "time"

type SubmissionTest struct {
	Id          uint64     `json:"id"`
	EvaluatedAt *time.Time `json:"evaluatedAt" db:"evaluated_at"`

	Score  int     `json:"score"`
	Time   float64 `json:"time"`
	Memory int     `json:"memory"` // in kb

	Message  string `json:"message"`
	ExitCode int    `json:"exitCode" db:"exit_code"`

	SubmissionId uint64 `json:"submissionId" db:"submission_id"`
	TestId       uint64 `json:"testId" db:"test_id"`
}

type UpdateSubmissionTestRequest struct {
	Score  int     `json:"score"`
	Time   float64 `json:"time"`
	Memory int     `json:"memory"`

	Message  string `json:"message"`
	ExitCode int    `json:"exitCode" db:"exit_code"` // -1 == do not update
}

func NewSubmissionTest(submissionId, testId uint64) *SubmissionTest {
	return &SubmissionTest{
		SubmissionId: submissionId,
		TestId:       testId,
	}
}
