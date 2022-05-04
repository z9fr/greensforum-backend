package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/z9fr/greensforum-backend/internal/helper"
	"github.com/z9fr/greensforum-backend/internal/question"
	"github.com/z9fr/greensforum-backend/internal/user"
)

// @Summary Create a new Question
// @Description create a new question
// @in header
// @Accept  json
// @Produce  json
// @Param payload body question.QuestionCreateRequest true "payload"
// @Success 200 {object} question.Question
// @Router /question/create [POST]
// @Security JWT
// @Tags Question
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

// @Summary fetch all posts
// @Description Get all the posts
// @in header
// @Accept  json
// @Produce  json
// @Param      next_post   query     int  true  "Next Post"
// @Success 200 {array} question.Question
// @Router /view/questions [GET]
// @Tags Question
func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	questions := h.QuestionService.GetAllPosts()
	h.sendOkResponse(w, questions)
	return
}

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
	h.sendOkResponse(w, questions)
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

// @Summary Write Answer
// @Description Answer to a question
// @Accept  json
// @Produce  json
// @Param payload body question.AnswerRequest true "payload"
// @Param   qid   path  uint  true  "Question ID"
// @Success 200 {array} question.Answer
// @Router /question/{qid}/answer/create [POST]
// @Tags Answer
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

	// check if question_id is actually a int
	qid, err := strconv.ParseUint(question_id, 10, 32)

	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	q, err := h.QuestionService.CreateAnswer(question, uint(qid))

	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Unable to Create a Post", err, 500)
		return
	}

	h.sendOkResponse(w, q)
}

// =========== find posts pagination test =========================
func (h *Handler) ListArticles(w http.ResponseWriter, r *http.Request) {
	pageID := r.Context().Value(PageIDKey).(int)
	questions := h.QuestionService.GetQuestionsPaginate(pageID)
	h.sendOkResponse(w, questions)
}
