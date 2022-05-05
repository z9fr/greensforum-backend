package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/user"
)

type VoteStruct struct {
	ID uint `json:"id"`
}

type SuccessResultUpvote struct {
	Success bool `json:"success"`
}

// @Summary Upvote a question
// @Description upvote a question
// @Accept  json
// @Produce  json
// @Router /question/upvote [PATCH]
// @Tags Question
// @Param   payload   body VoteStruct true  "payload"
// @Success 200 {array} question.Question
// @Security JWT
func (h *Handler) UpvotePost(w http.ResponseWriter, r *http.Request) {
	var vreq VoteStruct
	var u user.User
	u = r.Context().Value("user").(user.User)

	if err := json.NewDecoder(r.Body).Decode(&vreq); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	if !h.QuestionService.IsQuestionExist(vreq.ID) {
		h.sendErrorResponse(w, "404 question not found", errors.New("question with that id not found"), 500)
		return
	}

	question := h.QuestionService.GetQuestionByID(vreq.ID)
	complete_user, err := h.UserService.GetUserByEmail(u.Email)

	if err != nil {
		h.sendErrorResponse(w, "unable to find a user with that email", err, 500)
		return
	}

	upvoted, interests := h.QuestionService.UpVotePost(u, question)

	if upvoted {
		h.UserService.AppendUserInterests(complete_user, interests)
	}

	h.sendOkResponse(w, &SuccessResultUpvote{
		Success: upvoted,
	})
}
