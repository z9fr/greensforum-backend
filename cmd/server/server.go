package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/z9fr/greensforum-backend/internal/database"

	transportHttp "github.com/z9fr/greensforum-backend/internal/transport/http"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	//log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up Application")

	_, err := database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHttp.NewHandler()
	handler.SetupRotues()

	if err := http.ListenAndServe(":4000", handler.Router); err != nil {
		return err
	}

	return nil
}

func Start() {
	app := App{
		Name:    "api-greenforum-staging.dasith.works",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error(err)
	}
}
