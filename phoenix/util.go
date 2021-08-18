package phoenix

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/marius004/phoenix/database"
	"github.com/marius004/phoenix/models"
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

func newConfig(path string) *models.Config {
	config, err := models.NewConfig(path)

	if err != nil {
		panic(err)
	}

	return config
}

func newDatabase(config *models.Config, logger *log.Logger) *database.DB {
	db, err := database.DefaultDatabase(config)

	if err != nil {
		panic(err)
	}

	db.RunMigrations(logger)
	return db
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(w, nil)
}

func staticServer(path string) http.Handler {
	return http.FileServer(http.Dir(path))
}
