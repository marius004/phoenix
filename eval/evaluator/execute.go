package evaluator

import (
	"fmt"

	"github.com/marius004/phoenix/models"
)

type ExecuteHandler struct {
}

func (e *ExecuteHandler) Handle(ch chan *models.Submission) {
	for submission := range ch {
		fmt.Println("GONNA EXECUTE ", submission.Id)
	}
}
