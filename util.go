package main

import (
	"log"
	"os"

	"github.com/marius004/phoenix/database"
	"github.com/marius004/phoenix/internal"
)

func newLogger(path string) *log.Logger {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	logger := &log.Logger{}
	logger.SetFlags(log.LstdFlags | log.Ldate | log.Llongfile)
	logger.SetOutput(file)

	return logger
}

func newConfig(path string) *internal.Config {
	config, err := internal.NewConfig(path)

	if err != nil {
		panic(err)
	}

	return config
}

func newDatabase(config *internal.Config, logger *log.Logger) *database.DB {
	db, err := database.DefaultDatabase(config)

	if err != nil {
		panic(err)
	}

	db.RunMigrations(logger)
	return db
}
