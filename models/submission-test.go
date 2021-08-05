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

func NewSubmissionTest(score int, time float64, memory int, submissionId, userId, testId uint64) *SubmissionTest {
	return &SubmissionTest{
		Score:  score,
		Time:   time,
		Memory: memory,

		SubmissionId: submissionId,
		TestId:       testId,
	}
}
