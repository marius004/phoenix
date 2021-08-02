package tasks

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/models"
)

type ExecuteTask struct {
	Config *models.Config
	Logger *log.Logger

	Request  *eval.ExecuteRequest
	Response *eval.ExecuteResponse
}

func (t *ExecuteTask) Run(ctx context.Context, sandbox eval.Sandbox) error {
	t.Logger.Printf("Executing using sandbox %d\n", sandbox.GetID())
	lang, ok := t.Config.Languages[t.Request.Lang]

	if !ok {
		t.Logger.Printf("Invalid language %s\n", t.Request.Lang)
		return errors.New("no language found")
	}

	if err := sandbox.WriteToFile("box/"+t.Request.ProblemName+".in", t.Request.Input, 0644); err != nil {
		t.Logger.Println("Can't write the input file in the sandbox", err)
		t.Response.Message = "Sandbox error: Cannot copy input file to the sandbox"
		return err
	}

	if err := sandbox.CreateFile("box/"+t.Request.ProblemName+".out", 0644); err != nil {
		t.Logger.Println("Can't write the output file in the sandbox", err)
		t.Response.Message = "Sandbox error: Cannot copy output file to the sandbox"
		return err
	}

	binaryPath := path.Join(t.Config.CompilePath, fmt.Sprintf("%d.bin", t.Request.ID))
	binaryFile, err := os.OpenFile(binaryPath, os.O_RDONLY, 0644)

	if err != nil {
		t.Logger.Println("Could not open the binary file", err)
		t.Response.Message = "Sandbox error: Could not open the binary file"
		return err
	}

	bin, err := ioutil.ReadAll(binaryFile)

	if err != nil {
		t.Logger.Println("Could not open the read the binary file", err)
		t.Response.Message = "Sandbox error: Could not read the binary file"
		return err
	}

	if err := eval.CopyInSandbox(sandbox, lang.Executable, bin); err != nil {
		t.Logger.Println(fmt.Sprintf("Could not copy the binary file in the sandbox %d", sandbox.GetID()), err)
		t.Response.Message = fmt.Sprintf("Could not copy the binary file in the sandbox %d", sandbox.GetID())
		return err
	}

	limit := eval.Limit{
		Time: t.Request.Time,

		Memory: t.Request.Memory,
		Stack:  t.Request.Stack,
	}

	metaFile, err := eval.ExecuteFile(ctx, sandbox, lang, t.Request.ProblemName, limit, t.Request.IsConsole)
	if err != nil {
		t.Logger.Println("could not execute the program", err)
		t.Response.Message = fmt.Sprintf("Could not execute the program %s", err.Error())
		return err
	}

	t.Response.TimeUsed = metaFile.Time
	t.Response.MemoryUsed = metaFile.Memory

	switch metaFile.Status {
	case "TO":
		t.Response.Message = "TLE: " + metaFile.Message
	case "RE":
		t.Response.Message = "Runtime Error: " + metaFile.Message
	case "SG":
		t.Response.Message = metaFile.Message
	case "XX":
		t.Response.Message = "Sandbox error: " + metaFile.Message
	}

	if t.Response.ExitCode == 0 {
		path := fmt.Sprintf("%s/s%dt%d.out", t.Config.OutputPath, t.Request.ID, t.Request.TestId)
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)

		if err != nil {
			t.Logger.Println(err)
			return err
		}

		if err := eval.CopyFromSandbox(sandbox, "box/"+t.Request.ProblemName+".out", file); err != nil {
			t.Logger.Println("Could not copy the output file " + err.Error())
			t.Response.Message = "Could not copy the output file " + err.Error()
			return err
		}
	}

	return nil
}
