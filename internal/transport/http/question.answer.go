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

// @Summary Write Answer
// @Description Answer to a question
// @Accept  json
// @Produce  json
// @Param payload body question.AnswerRequest true "payload"
// @Param   qid   path  uint  true  "Question ID"
// @Success 200 {array} question.Answer
// @Router /question/{qid}/answer/create [POST]
// @Security JWT
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
	u64, err := strconv.ParseUint(question_id, 10, 32)
	ogquestion := h.QuestionService.GetQuestionByID(uint(u64))
	question_ownser, err := h.UserService.GetUserbyOnlyID(uint(ogquestion.CreatedBy))

	if err != nil {
		fmt.Println(err)
	}

	var notifiaction user.Nofication
	notifiaction.Message = "New answer has been posted to your question. view <a href=/q/" + question_id + "/> here </a>"
	h.UserService.SendNofications(question_ownser, notifiaction)

	h.sendOkResponse(w, q)
}
