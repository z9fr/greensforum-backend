package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
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

		r.Route("/user", func(r chi.Router) {
			r.Post("/join", h.CreateUser)
			r.Post("/login", h.Login)
			r.Get("/all", h.GetAllUsers)
		})

		r.Route("/", func(r chi.Router) {
			r.Use(h.JWTMiddlewhare)
			r.Get("/test", h.TestRoute)
		})

	})

	h.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))

}

func ListArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list articles"))
	return
}

func (h *Handler) TestRoute(w http.ResponseWriter, r *http.Request) {
	var u user.User
	u = r.Context().Value("user").(user.User)
	h.sendOkResponse(w, u.Email)
	return

}
