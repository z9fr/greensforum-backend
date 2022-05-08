package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z9fr/greensforum-backend/internal/collective"
	"github.com/z9fr/greensforum-backend/internal/question"
)

// @Summary get posts by tags
// @Description find posts by using tag
// @Accept  json
// @Produce  json
// @Param    tag   path      string  true  "Tag Name"
// @Success 200 {array} question.Question
// @Router  /view/questions/{tag} [GET]
// @Tags Question
func (h *Handler) FindPostsByTag(w http.ResponseWriter, r *http.Request) {
	tag := chi.URLParam(r, "tag")
	//questions := h.QuestionService.SearchQuestionsByTags(tag)
	questions := h.QuestionService.SearchQuestionsByTagsv2(tag)
	posts := h.CollectiveService.GetPostsByTags([]string{tag})

	//	h.sendOkResponse(w, )
	h.sendOkResponse(w, struct {
		Questions []*question.Question
		Posts     []collective.Post `json:"posts"`
	}{
		questions,
		posts,
	})
}

// @Summary search for posts
// @Description Search Posts based on a keyword
// @Accept  json
// @Produce  json
// @Param   q   query     string  true  "Search Query"
// @Success 200 {array} question.Question
// @Router /view/search [GET]
// @Tags Question
func (h *Handler) SearchPost(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		h.sendErrorResponse(w, "Search Parameter Missing", fmt.Errorf("search parameter `q` is missing"), 401)
		return
	}

	results := h.QuestionService.SearchPosts(q)
	h.sendOkResponse(w, results)
}

func (h *Handler) GetQuestionBasedonSlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	question := h.QuestionService.GetQuestionBasedonSlug(slug)

	h.sendOkResponse(w, question)
}
