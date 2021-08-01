package eval

import (
	"context"
	"io"
)

type Sandbox interface {
	GetID() int
	GetPath(path string) string

	CreateDirectory(path string) error
	DeleteDirectory(path string) error

	CreateFile(path string) error
	WriteToFile(path string, data []byte) error
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

	TimeLimit 	  float64
	WallTimeLimit float64

	MaxProcesses int
}

// RunStatus is what is returned after a task has been run.
type RunStatus struct {
	Memory int `json:"memory"`

	ExitCode   int  `json:"exitCode"`
	ExitSignal int  `json:"exitSignal"`
	Killed 	   bool `json:"killed"`

	Message string `json:"message"`
	Status  string `json:"status"`

	Time 	 float64 `json:"time"`
	WallTime float64 `json:"wallTime"`
}

type CompileRequest struct {
	ID   int  // ID is the file-managers ID.
	Code []byte
	Lang string 
}

type CompileResponse struct {
	Output string
	Other string

	Success bool 
}

type Limit struct {
	Time 	float64
	Memory  int
	Stack 	int
}

type ExecuteRequest struct {
	ID int

	SubmissionId int
	TestId 		 int

	Limit

	Lang 		string
	ProblemName string

	IsConsole 	bool
	Input 		[]byte

	BinaryPath  string
}

type ExecuteResponse struct {
	TimeUsed 	float64
	MemoryUsed  int

	ExitCode int
	Message string
}