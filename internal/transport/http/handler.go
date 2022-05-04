package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/z9fr/greensforum-backend/internal/question"
	topwords "github.com/z9fr/greensforum-backend/internal/top-words"
	"github.com/z9fr/greensforum-backend/internal/user"
)

type Handler struct {
	// service and router
	Router          *chi.Mux
	UserService     *user.Service
	QuestionService *question.Service
	TopWordsService *topwords.ITopTenWords
}

// NewHandler -  construcutre to create and return a pointer to a handler
func NewHandler(userService *user.Service, questionService *question.Service, topwordsservice *topwords.ITopTenWords) *Handler {
	return &Handler{
		UserService:     userService,
		QuestionService: questionService,
		TopWordsService: topwordsservice,
	}
}

func (h *Handler) SetupRotues() {
	h.Router = chi.NewRouter()

	// logs the start and end of each request, along with some useful data about what was requested,
	// what the response status was, and how long it took to return. When standard output is a TTY,
	// Logger will print in color, otherwise it will print in black and white. Logger prints a request ID if one is provided.
	h.Router.Use(middleware.Logger)

	// clean out double slash mistakes from a user's request path.
	// For example, if a user requests /users//1 or //users////1 will both be treated as: /users/1
	h.Router.Use(middleware.CleanPath)

	// automatically route undefined HEAD requests to GET handlers.
	h.Router.Use(middleware.GetHead)

	// recovers from panics, logs the panic (and a backtrace),
	// returns a HTTP 500 (Internal Server Error) status if possible. Recoverer prints a request ID if one is provided.
	h.Router.Use(middleware.Recoverer)

	h.Router.Route("/api/v1", func(r chi.Router) {
		r.Get("/", ListArticles)

		r.Route("/user", func(r chi.Router) {
			r.Post("/join", h.CreateUser)
			r.Post("/login", h.Login)
			r.Get("/all", h.GetAllUsers)
		})

		r.Route("/view", func(r chi.Router) {
			r.With(h.Pagination).Get("/p2", h.ListArticles)

			r.Get("/questions", h.GetAllPosts)
			r.Get("/search", h.SearchPost)
			r.Get("/questions/{tag}", h.FindPostsByTag)
		})

		r.Route("/question", func(r chi.Router) {
			r.Use(h.JWTMiddlewhare)
			r.Post("/create", h.CreatePost)
			r.Post("/{qid}/answer/create", h.WriteAnswer)
			r.Patch("/upvote", h.UpvotePost)
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
