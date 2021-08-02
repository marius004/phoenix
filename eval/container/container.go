package container

import (
	"bufio"
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/models"
)

// Container implements eval.Sandbox
type Container struct {
	path string
	id   int

	metaFile string

	config *models.Config
	logger *log.Logger
}

func (c *Container) GetPath(path string) string {
	if path == "" {
		return c.path
	}
	return c.path + "/" + path
}

func (c *Container) CreateDirectory(path string, perm fs.FileMode) error {
	fullPath := c.GetPath(path)
	return os.Mkdir(fullPath, perm)
}

func (c *Container) DeleteDirectory(path string) error {
	fullPath := c.GetPath(path)
	return os.RemoveAll(fullPath)
}

func (c *Container) WriteToFile(path string, data []byte, perm fs.FileMode) error {
	fullPath := c.GetPath(path)
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	return err
}

func (c *Container) CreateFile(path string, perm fs.FileMode) error {
	fullPath := c.GetPath(path)
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_RDONLY|os.O_TRUNC, perm)

	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}

func (c *Container) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(c.GetPath(path))
}

func (c *Container) DeleteFile(path string) error {
	return os.Remove(c.GetPath(path))
}

func (c *Container) ExecuteCommand(ctx context.Context, command []string, config *eval.RunConfig) (*eval.RunStatus, error) {
	metaFile := path.Join(os.TempDir(), "pn-"+eval.RandomString(24))
	c.metaFile = metaFile

	defer func() { c.metaFile = "" }()

	c.logger.Print("Command meta file: ", c.metaFile)
	params := append(c.buildRunFlags(config), command...)

	c.logger.Println("Command to be executed:", "isolate", params)
	cmd := exec.CommandContext(ctx, c.config.IsolatePath, params...)

	cmd.Stdin = config.Stdin
	cmd.Stdout = config.Stdout
	cmd.Stderr = config.Stderr

	err := cmd.Run()
	if _, ok := err.(*exec.ExitError); ok {
		metaData, metaFileErr := parseMetaFile(c.metaFile)
		if metaFileErr != nil {
			c.logger.Println(err)
			return nil, err
		} else { // this means that the program was stopped because of the time or memory constraints.
			return metaData, nil
		}
	}

	return parseMetaFile(c.metaFile)
}

func (c *Container) Cleanup() error {
	var params []string

	params = append(params, "--cg")
	params = append(params, "--box-id="+fmt.Sprintf("%d", c.id))
	params = append(params, "--cleanup")

	return exec.Command(c.config.IsolatePath, params...).Run()
}

func (c *Container) GetID() int {
	return c.id
}

func (c *Container) buildRunFlags(config *eval.RunConfig) (res []string) {
	res = append(res, "--box-id="+strconv.Itoa(c.id))
	res = append(res, "--cg", "--cg-timing")

	res = append(res, "--full-env")

	if config.TimeLimit != 0 {
		res = append(res, "--time="+strconv.FormatFloat(config.TimeLimit, 'f', -1, 64))
	}

	if config.WallTimeLimit != 0 {
		res = append(res, "--wall-time="+strconv.FormatFloat(config.WallTimeLimit, 'f', -1, 64))
	}

	if config.MemoryLimit != 0 {
		res = append(res, "--mem="+strconv.Itoa(config.MemoryLimit))
	}

	if config.StackLimit != 0 {
		res = append(res, "--stack="+strconv.Itoa(config.StackLimit))
	}

	if config.MaxProcesses == 0 {
		res = append(res, "--processes=1")
	} else {
		res = append(res, "--processes="+strconv.Itoa(config.MaxProcesses))
	}

	if config.InputPath != "" {
		res = append(res, "--stdin="+config.InputPath)
	}

	if config.OutputPath != "" {
		res = append(res, "--stdout="+config.OutputPath)
	}

	if c.metaFile != "" {
		res = append(res, "--meta="+c.metaFile)
	}

	res = append(res, "--silent", "--run", "--")
	return
}

func newContainer(id int, config *models.Config, logger *log.Logger) (*Container, error) {
	ret, err := exec.Command(config.IsolatePath, fmt.Sprintf("--box-id=%d", id), "--cg", "--init").CombinedOutput()

	if strings.HasPrefix(string(ret), "Box already exists") {
		exec.Command(config.IsolatePath, fmt.Sprintf("--box-id=%d", id), "--cg", "--cleanup").Run()
		return newContainer(id, config, logger)
	}

	if err != nil {
		logger.Print("Could not create sandbox", err)
		return nil, err
	}

	path := string(ret)
	return &Container{path: strings.TrimSpace(path), id: id, config: config, logger: logger}, nil
}

func parseMetaFile(path string) (*eval.RunStatus, error) {
	r, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}

	var ret = new(eval.RunStatus)
	s := bufio.NewScanner(r)

	for s.Scan() {
		if !strings.Contains(s.Text(), ":") {
			continue
		}
		l := strings.SplitN(s.Text(), ":", 2)
		switch l[0] {
		case "cg-mem":
			ret.Memory, _ = strconv.Atoi(l[1])
		case "exitcode":
			ret.ExitCode, _ = strconv.Atoi(l[1])
		case "exitsig":
			ret.ExitSignal, _ = strconv.Atoi(l[1])
		case "killed":
			ret.Killed = true
		case "message":
			ret.Message = l[1]
		case "status":
			ret.Status = l[1]
		case "time":
			ret.Time, _ = strconv.ParseFloat(l[1], 32)
		case "time-wall":
			ret.WallTime, _ = strconv.ParseFloat(l[1], 32)
		default:
			continue
		}
	}

	return ret, nil
}
