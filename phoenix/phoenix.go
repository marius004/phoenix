package phoenix

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/eval/container"
	"github.com/marius004/phoenix/eval/tasks"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/marius004/phoenix/api"
	"github.com/marius004/phoenix/database"
	"github.com/marius004/phoenix/models"
)

// Phoenix is a struct that wraps the functionality of the entire app
type Phoenix struct {
	db     *database.DB
	logger *log.Logger
	api    *api.API
	config *models.Config
}

// New returns an instance of Phoenix
func New() *Phoenix {
	logger := newLogger("logs.txt")
	config := newConfig("./config.json")
	db := newDatabase(config, logger)
	api := api.New(db, config, logger)

	return &Phoenix{
		logger: logger,
		config: config,
		db:     db,
		api:    api,
	}
}

func (p *Phoenix) Run() {
	r := chi.NewRouter()

	manager := container.NewManager(10, p.config, p.logger)

	// compile
	code := `
	#include <stdio.h>
	int main() {
   	// printf() displays the string inside quotation
	int i;
	for(i = 1;i <= 10;++i) {
	   	printf("Hello, World!");
	}
   	return 0;
	}`

	compile := &tasks.CompileTask{
		Config: p.config,
		Logger: p.logger,
		Request: &eval.CompileRequest{
			ID:   1,
			Code: []byte(code),
			Lang: "c",
		},
		Response: &eval.CompileResponse{},
	}

	err := manager.RunTask(context.Background(), compile)
	if err != nil {
		fmt.Println(err)
		fmt.Println(compile.Response)
	}

	// execute
	execute := &tasks.ExecuteTask{
		Config: p.config,
		Logger: p.logger,
		Request: &eval.ExecuteRequest{
			ID: 1,
			Limit: eval.Limit{
				Time:   1,
				Memory: 64000,
				Stack:  32000,
			},
			Lang:         "c",
			ProblemName:  "marsx",
			Input:        []byte{49, 48, 32, 50, 48},
			BinaryPath:   "/tmp/pn-compile",
			SubmissionId: 1,
			TestId:       2,
		},
		Response: &eval.ExecuteResponse{},
	}

	err = manager.RunTask(context.Background(), execute)
	if err != nil {
		fmt.Println(err)
		fmt.Println(execute.Response)
	}

	//change cors config according to my needs
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // https://stackoverflow.com/questions/24687313/what-exactly-does-the-access-control-allow-credentials-header-do
		MaxAge:           7200, // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age
	})

	r.Use(corsConfig.Handler)

	r.Mount("/api", p.api.Routes())

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", p.config.Server.Host, p.config.Server.Port),
		Handler: r,
	}

	fmt.Printf("Phoenix running on %s:%s\n", p.config.Server.Host, p.config.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
