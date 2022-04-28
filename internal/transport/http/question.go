package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z9fr/greensforum-backend/internal/helper"
	"github.com/z9fr/greensforum-backend/internal/question"
	"github.com/z9fr/greensforum-backend/internal/user"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var questionreq question.QuestionCreateRequest
	var u user.User
	u = r.Context().Value("user").(user.User)

	if err := json.NewDecoder(r.Body).Decode(&questionreq); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	question, err := helper.RequestPostWithValidation(questionreq, u)

	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Post Validation Failed", err, 500)
		return
	}

	q, err := h.QuestionService.CreateNewQuestion(question)

	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Unable to Create a Post", err, 500)
		return
	}

	h.sendOkResponse(w, q)
}

//@TODO
// Add pagination
func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	questions := h.QuestionService.GetAllQuestions()
	h.sendOkResponse(w, questions)
	return
}

func (h *Handler) FindPostsByTag(w http.ResponseWriter, r *http.Request) {
	tag := chi.URLParam(r, "tag")
	questions := h.QuestionService.SearchQuestionsByTags(tag)
	h.sendOkResponse(w, questions)
}

func (h *Handler) SearchPost(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		h.sendErrorResponse(w, "Search Parameter Missing", fmt.Errorf("search parameter `q` is missing"), 401)
		return
	}

	results := h.QuestionService.SearchQuestions(q)
	h.sendOkResponse(w, results)
}

func (h *Handler) WriteAnswer(w http.ResponseWriter, r *http.Request) {
	var reqanswer question.AnswerRequest
	question_id := chi.URLParam(r, "qid")

	var u user.User
	u = r.Context().Value("user").(user.User)

	if err := json.NewDecoder(r.Body).Decode(&reqanswer); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	// validate the request
	question, err := helper.RequstAnswerValidation(reqanswer, u, question_id)
	q, err := h.QuestionService.CreateAnswer(question, question_id)

	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Unable to Create a Post", err, 500)
		return
	}

	h.sendOkResponse(w, q)
}
