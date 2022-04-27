package server

import (
	log "github.com/sirupsen/logrus"
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
