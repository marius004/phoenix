package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/marius004/phoenix/api"
	"github.com/marius004/phoenix/database"
	"github.com/marius004/phoenix/internal"
)

// Phoenix is a struct that wraps the functionality of the entire app
type Phoenix struct {
	db  *database.DB
	api *api.API

	logger *log.Logger
	config *internal.Config
}

// NewServer returns an instance of Phoenix
func NewServer() *Phoenix {
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

func main() {
	NewServer().Run()
}
