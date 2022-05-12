package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/z9fr/greensforum-backend/internal/collective"
	"github.com/z9fr/greensforum-backend/internal/database"
	"github.com/z9fr/greensforum-backend/internal/events"
	"github.com/z9fr/greensforum-backend/internal/feed"
	"github.com/z9fr/greensforum-backend/internal/question"
	topwords "github.com/z9fr/greensforum-backend/internal/top-words"
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/verification"

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

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	if err := database.MigrateDB(db); err != nil {
		return err
	}

	userService := user.NewService(db)
	questionservice := question.NewService(db)
	topwordsService := topwords.InitTopTenWordsService()
	collectiveService := collective.NewService(db)
	feedservice := feed.NewService(db)
	verificationService := verification.NewService(db)
	eventservice := events.NewService(db)

	handler := transportHttp.NewHandler(userService, questionservice,
		topwordsService, collectiveService, feedservice, verificationService, eventservice)
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
