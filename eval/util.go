package eval

import (
	"bytes"
	"context"
	"github.com/marius004/phoenix/models"
	"io"
	"math/rand"
	"strings"
)

const randomCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(size int) string {
	sb := strings.Builder{}
	sb.Grow(size)

	for ; size > 0; size-- {
		randIndex := rand.Intn(len(randomCharacters))
		sb.WriteByte(randomCharacters[randIndex])
	}

	return sb.String()
}

func CompileFile(ctx context.Context, sandbox Sandbox, sourceCode []byte, lang models.Language) (string, error) {
	if err := sandbox.WriteToFile(lang.SourceFile, sourceCode); err != nil {
		return "", err
	}

	var runConf RunConfig

	out := &bytes.Buffer{}

	runConf.Stdout = out
	runConf.Stderr = out
	runConf.MaxProcesses = 5

	if _,err := sandbox.ExecuteCommand(ctx, lang.Compile, &runConf); err != nil {
		return out.String(), err
	}

	return out.String(), nil
}

func CopyFromSandbox(sandbox Sandbox, path string, w io.Writer) error {
	content, err := sandbox.ReadFile(path)

	if err != nil {
		return err
	}

	if _, err := w.Write(content); err != nil {
		return err
	}

	return nil
}

func CopyInSandbox(sandbox Sandbox, path string, data []byte) error {
	return sandbox.WriteToFile(path, data)
}

func ExecuteFile(ctx context.Context, sandbox Sandbox, lang models.Language, problemName string, limit Limit, console bool) (*RunStatus, error) {
	var runConf RunConfig

	// limit stuff
	runConf.MaxProcesses  = 10
	runConf.MemoryLimit   = limit.Memory
	runConf.TimeLimit     = limit.Time
	runConf.StackLimit    = limit.Stack
	runConf.WallTimeLimit = 10 // should be enough for now

	runConf.InputPath = problemName+".in"
	runConf.OutputPath = problemName+".out"

	return sandbox.ExecuteCommand(ctx, lang.Execute, &runConf)
}