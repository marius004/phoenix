package eval

import (
	"bytes"
	"context"
	"io"
	"strconv"

	"github.com/marius004/phoenix/internal"
	"github.com/marius004/phoenix/internal/models"
)

func CompileFile(ctx context.Context, sandbox internal.Sandbox, sourceCode []byte, lang internal.Language) (string, error) {
	if err := sandbox.WriteToFile(lang.SourceFile, sourceCode, 0644); err != nil {
		return "", err
	}

	var runConf internal.RunConfig

	out := &bytes.Buffer{}

	runConf.Stdout = out
	runConf.Stderr = out
	runConf.MaxProcesses = 5

	if _, err := sandbox.ExecuteCommand(ctx, lang.Compile, &runConf); err != nil {
		return out.String(), err
	}

	return out.String(), nil
}

func CopyFromSandbox(sandbox internal.Sandbox, path string, w io.Writer) error {
	content, err := sandbox.ReadFile(path)

	if err != nil {
		return err
	}

	if _, err := w.Write(content); err != nil {
		return err
	}

	return nil
}

func CopyInSandbox(sandbox internal.Sandbox, path string, data []byte) error {
	return sandbox.WriteToFile(path, data, 7777)
}

func ExecuteFile(ctx context.Context, sandbox internal.Sandbox, lang internal.Language, problemName string, limit internal.Limit, console bool) (*internal.RunStatus, error) {
	var runConf internal.RunConfig

	// limit stuff
	runConf.MaxProcesses = 10
	runConf.MemoryLimit = limit.Memory
	runConf.TimeLimit = limit.Time
	runConf.StackLimit = limit.Stack
	runConf.WallTimeLimit = 2 * limit.Time // should be enough for now

	runConf.InputPath = problemName + ".in"
	runConf.OutputPath = problemName + ".out"

	return sandbox.ExecuteCommand(ctx, lang.Execute, &runConf)
}

func GetBinaryName(config *internal.Config, submissionId int) string {
	return config.CompilePath + "/" + strconv.Itoa(submissionId) + ".bin"
}

func GetOutputFileName(config *internal.Config, submission *models.Submission, test *models.Test) string {
	return config.OutputPath + "/s" + strconv.Itoa(int(submission.Id)) + "t" + strconv.Itoa(int(test.Id)) + ".out"
}

func CompiledSourceCode(sandbox internal.Sandbox, fileName string) bool {
	return sandbox.FileExists(fileName)
}
