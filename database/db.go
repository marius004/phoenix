package database

import (
	"context"
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"

	"github.com/marius004/phoenix/models"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/marius004/phoenix/services"
)

//go:embed psql_schema
var psqlSchemaDirectory embed.FS

type DB struct {
	Conn *sqlx.DB
}

// NewPSQL returns a new psql database connection
func NewPSQL(dsn string) (*DB, error) {
	conn, err := sqlx.Connect("pgx", dsn)

	if err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}

// DefaultDatabase returns an instance of a psql database
func DefaultDatabase(config *models.Config) (*DB, error) {
	var (
		host   = config.Database.Host
		port   = config.Database.Port
		user   = config.Database.User
		pass   = config.Database.Password
		dbname = config.Database.Name
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, pass, dbname)
	db, err := NewPSQL(dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Conn.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// RunMigrations creates the database schema for the database
func (db *DB) RunMigrations(logger *log.Logger) {
	fs.WalkDir(psqlSchemaDirectory, "psql_schema", func(path string, d fs.DirEntry, err error) error {
		content, err := fs.ReadFile(psqlSchemaDirectory, path)

		if err != nil {
			logger.Println(err)
		}

		sqlQuery := string(content)
		_, err = db.Conn.ExecContext(context.Background(), sqlQuery)

		if err != nil {
			logger.Println(err)
		}

		return nil
	})
}

func (db *DB) UserService(logger *log.Logger) services.UserService {
	return NewUserService(db.Conn, logger)
}

func (db *DB) ProblemService(logger *log.Logger) services.ProblemService {
	return NewProblemService(db.Conn, logger)
}

func (db *DB) TestService(logger *log.Logger) services.TestService {
	return NewTestService(db.Conn, logger)
}

func (db *DB) SubmissionService(logger *log.Logger) services.SubmissionService {
	return NewSubmissionService(db.Conn, logger)
}

func (db *DB) SubmissionTestService(logger *log.Logger) services.SubmissionTestService {
	return NewSubmissionTestService(db.Conn, logger)
}
