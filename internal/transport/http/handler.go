package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	// service and router
	Router *chi.Mux
}

// NewHandler -  construcutre to create and return a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRotues() {
	h.Router = chi.NewRouter()
	h.Router.Use(middleware.Logger)

	apiRouter := chi.NewRouter()
	apiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	h.Router.Mount("/api/v1", apiRouter)
}
