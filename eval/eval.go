package eval

import (
	"context"
	"io"
	"io/fs"

	"github.com/marius004/phoenix/managers"
	"github.com/marius004/phoenix/models"
	"github.com/marius004/phoenix/services"
)

type Sandbox interface {
	GetID() int
	GetPath(path string) string

	CreateDirectory(path string, perm fs.FileMode) error
	DeleteDirectory(path string) error

	CreateFile(path string, perm fs.FileMode) error
	WriteToFile(path string, data []byte, perm fs.FileMode) error
	ReadFile(path string) ([]byte, error)
	DeleteFile(path string) error

	ExecuteCommand(ctx context.Context, command []string, config *RunConfig) (*RunStatus, error)
	Cleanup() error
}

// Task represents a task to be executed in the sandbox.
type Task interface {
	Run(ctx context.Context, sandbox Sandbox) error
}

type SandboxManager interface {
	// RunTask runs a task within a sandbox
	RunTask(ctx context.Context, task Task) error

	// Stop waits for all sandboxes to stop running
	Stop(ctx context.Context) error
}

// RunConfig represents the configuration for a Task.
type RunConfig struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	MemoryLimit int
	StackLimit  int

	InputPath  string
	OutputPath string

	TimeLimit     float64
	WallTimeLimit float64

	MaxProcesses int
}

// RunStatus is what is returned after a task has been run.
type RunStatus struct {
	Memory int `json:"memory"`

	ExitCode   int  `json:"exitCode"`
	ExitSignal int  `json:"exitSignal"`
	Killed     bool `json:"killed"`

	Message string `json:"message"`
	Status  string `json:"status"`

	Time     float64 `json:"time"`
	WallTime float64 `json:"wallTime"`
}

type CompileRequest struct {
	ID   int // ID is the file-managers ID.
	Code []byte
	Lang string
}

type CompileResponse struct {
	Message string
	Success bool
}

type Limit struct {
	Time   float64
	Memory int
	Stack  int
}

type ExecuteRequest struct {
	ID int

	SubmissionId int
	TestId       int

	Limit

	Lang        string
	ProblemName string

	IsConsole bool
	Input     []byte

	BinaryPath string
}

type ExecuteResponse struct {
	TimeUsed   float64
	MemoryUsed int

	ExitCode int
	Message  string
}

type EvaluatorServices struct {
	ProblemService services.ProblemService
	TestService    services.TestService

	SubmissionService     services.SubmissionService
	SubmissionTestService services.SubmissionTestService

	TestManager managers.TestManager
}

type CompileServices struct {
	SubmissionService services.SubmissionService
}

type ExecuteServices struct {
	ProblemService services.ProblemService
	TestService    services.TestService

	SubmissionService     services.SubmissionService
	SubmissionTestService services.SubmissionTestService

	TestManager managers.TestManager
}

type CheckerServices struct {
	ProblemService services.ProblemService
	TestService    services.TestService

	SubmissionService     services.SubmissionService
	SubmissionTestService services.SubmissionTestService

	TestManager managers.TestManager
}

func (s *EvaluatorServices) CompileServices() *CompileServices {
	return &CompileServices{
		SubmissionService: s.SubmissionService,
	}
}

func (s *EvaluatorServices) ExecuteServices() *ExecuteServices {
	return &ExecuteServices{
		ProblemService: s.ProblemService,
		TestService:    s.TestService,

		SubmissionService:     s.SubmissionService,
		SubmissionTestService: s.SubmissionTestService,

		TestManager: s.TestManager,
	}
}

func (s *EvaluatorServices) CheckerServices() *CheckerServices {
	return &CheckerServices{
		ProblemService: s.ProblemService,
		TestService:    s.TestService,

		SubmissionService:     s.SubmissionService,
		SubmissionTestService: s.SubmissionTestService,

		TestManager: s.TestManager,
	}
}

type Handler interface {
	Handle(next chan *models.Submission)
}
