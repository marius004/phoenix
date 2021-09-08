package internal

import (
	"context"
	"io"
	"io/fs"

	"github.com/marius004/phoenix/internal/models"
	"github.com/marius004/phoenix/internal/services"
	"github.com/marius004/phoenix/managers"
)

type Sandbox interface {
	GetID() int
	GetPath(path string) string

	CreateDirectory(path string, perm fs.FileMode) error
	DeleteDirectory(path string) error

	FileExists(path string) bool
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
	// RunTask runs a task within a sandbox. waits if there is no box available
	RunTask(ctx context.Context, task Task) error

	// Stop waits for all sandboxes to stop running
	Stop(ctx context.Context) error
}

// RunConfig represents the configuration needed for a task to be run.
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

// RunStatus is what is "the status" returned after a task has been run.
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

type Checker interface {
	Check(submission *models.Submission) error
}
