package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/z9fr/greensforum-backend/internal/user"
)

type Handler struct {
	// service and router
	Router      *chi.Mux
	UserService *user.Service
}

// NewHandler -  construcutre to create and return a pointer to a handler
func NewHandler(userService *user.Service) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (h *Handler) SetupRotues() {
	h.Router = chi.NewRouter()
	h.Router.Use(middleware.Logger)

	h.Router.Route("/api/v1", func(r chi.Router) {
		r.Get("/", ListArticles)
		r.Get("/test", TestRoute)

		r.Route("/user", func(r chi.Router) {
			r.Post("/join", h.CreateUser)
			r.Get("/all", h.GetAllUsers)
		})

		r.Route("/sub", func(r chi.Router) {
			r.Get("/", ListArticles)
			r.Get("/subtest", TestRoute)
		})

	})

}

func ListArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("articles"))
	return
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
	return
}
