package models

import (
	"os"

	"encoding/json"
)

type Command struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type Eval struct {
	IsolatePath string `json:"isolatePath"`

	// Concurrency stuff
	MaxSandboxes int `json:"maxSandboxes"`
	MaxCompile   int `json:"maxCompile"`
	MaxExecute   int `json:"maxExecute"`
	MaxCheck     int `json:"maxCheck"`

	CompilePath string `json:"compilePath"`
	OutputPath  string `json:"outputPath"`
	LoggerPath  string `json:"loggerPath"`
}

type Language struct {
	Extension  string `json:"extension"`
	IsCompiled bool   `json:"isCompiled"`

	Compile []string `json:"compile"`
	Execute []string `json:"execute"`

	SourceFile string `json:"sourceFile"`
	Executable string `json:"executable"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Api struct {
	Jwt                string `json:"jwt"`
	AuthCookieLifeTime int    `json:"authCookieLifeTime"`
}

// Config is the struct behind the config.json file.
// It holds basic configurations for the server to run
type Config struct {
	Database `json:"database"`
	Server   `json:"server"`

	Api  `json:"api"`
	Eval `json:"eval"`

	Languages map[string]Language `json:"languages"`
}

// JwtSecret is a utility function that returns config.Api.Jwt
// Added to ease the manipulation of the config struct.
func (c *Config) JwtSecret() string {
	return c.Api.Jwt
}

// NewConfig returns a new Config object that will be read from the configPath provided
func NewConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
