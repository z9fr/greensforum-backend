package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/user"
)

type VoteStruct struct {
	ID uint `json:"id"`
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

	fmt.Println(u)

	if err := json.NewDecoder(r.Body).Decode(&vreq); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	question := h.QuestionService.GetQuestionByID(vreq.ID)

	if question.Title == "" {
		return
	}

	h.QuestionService.UpVotePost(&u, &question)
	h.sendOkResponse(w, map[string]interface{}{"success": true})
}
