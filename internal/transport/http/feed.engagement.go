package http

import (
	"errors"
	"net/http"

	_ "github.com/z9fr/greensforum-backend/internal/question"
	"github.com/z9fr/greensforum-backend/internal/user"
)

// @Summary user feed based on engadgement
// @Description fetch questions based on recent engadgement
// @in header
// @Accept  json
// @Produce  json
// @Success 200 {array}     question.Question

// @Router /user/feed [GET]
// @Tags Feed
func (h *Handler) GetEngagementFeed(w http.ResponseWriter, r *http.Request) {
	var u user.User
	u = r.Context().Value("user").(user.User)
	// run algorithm to get data
	values := h.FeedService.GetUserInterestedQuestionsEngagement(u)

	// just return a error if there's no posts
	if values == nil {
		h.sendErrorResponse(w, "unable to generate posts", errors.New("can't find posts. please upvote more posts"), http.StatusServiceUnavailable)
		return
	}

	h.sendOkResponse(w, values)

}
