package http

import (
	"encoding/json"
	"net/http"

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
