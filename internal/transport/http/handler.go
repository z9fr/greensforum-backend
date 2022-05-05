package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/z9fr/greensforum-backend/internal/collective"
	"github.com/z9fr/greensforum-backend/internal/question"
	topwords "github.com/z9fr/greensforum-backend/internal/top-words"
	"github.com/z9fr/greensforum-backend/internal/user"
)

type Handler struct {
	// service and router
	Router            *chi.Mux
	UserService       *user.Service
	QuestionService   *question.Service
	TopWordsService   topwords.ITopTenWords
	CollectiveService *collective.Service
}

// NewHandler -  construcutre to create and return a pointer to a handler
func NewHandler(userService *user.Service,
	questionService *question.Service,
	topwordsservice topwords.ITopTenWords,
	collectiveService *collective.Service) *Handler {
	return &Handler{
		UserService:       userService,
		QuestionService:   questionService,
		TopWordsService:   topwordsservice,
		CollectiveService: collectiveService,
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

		r.Route("/collectives", func(r chi.Router) {
			r.Get("/", h.FetchCollectives)
			r.Get("/{collective}", h.GetCollectiveBySlug)

			r.Route("/create", func(r chi.Router) {
				r.Use(h.JWTMiddlewhare)
				r.Use(h.HighPrivilagesMiddlewhare)
				r.Post("/", h.CreateCollective)
			})

			r.Route("/{collective}/post/", func(r chi.Router) {
				r.Use(h.JWTMiddlewhare)
				r.Post("/write", h.WritePostinCollective)
			})
		})
		r.Route("/", func(r chi.Router) {
			r.Use(h.JWTMiddlewhare)
			r.Get("/test", h.TestRoute)
		})

	})

	h.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))

	/* handle errors */

	h.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "route not found"})
	})

	h.Router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "method is not valid"})
	})

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
