package tasks

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/models"
)

type CompileTask struct {
	Config *models.Config
	Logger *log.Logger

	Request  *eval.CompileRequest
	Response *eval.CompileResponse
}

func (task *CompileTask) Run(ctx context.Context, sandbox eval.Sandbox) error {
	task.Logger.Printf("Compiling using sandbox %d\n", sandbox.GetID())

	lang, ok := task.Config.Languages[task.Request.Lang]

	if !ok {
		task.Logger.Printf("Invalid language %s\n", task.Request.Lang)
		return errors.New("no language found")
	}

	binaryPath := path.Join(task.Config.CompilePath, fmt.Sprintf("%d.bin", task.Request.ID))
	task.Response.Success = true

	if lang.IsCompiled {
		output, err := eval.CompileFile(ctx, sandbox, task.Request.Code, lang)
		task.Response.Output = output

		if err != nil {
			task.Response.Success = false
			task.Response.Output = err.Error()
			task.Logger.Printf("Could not compile %s\n", err.Error())
			return err
		} else if output != "" {
			task.Response.Success = false
			task.Response.Output = output
			return nil
		}

		file, err := os.OpenFile(binaryPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)

		if err != nil {
			task.Response.Output = err.Error()
			task.Response.Success = false
			task.Logger.Printf("Could not create the binary file %s\n", err.Error())
			return err
		}

		if err := eval.CopyFromSandbox(sandbox, lang.Executable, file); err != nil {
			task.Response.Output = err.Error()
			task.Response.Success = false
			task.Logger.Printf("Could not copy the binary file from sandbox %d %s\n", sandbox.GetID(), err.Error())
			return err
		}

		if err := file.Close(); err != nil {
			task.Response.Output = err.Error()
			task.Response.Success = false
			task.Logger.Printf("Could not close the binary file %s\n", err.Error())
			return err
		}

		return nil
	}

	// TODO add dynamic typed programming languages
	return nil
}
