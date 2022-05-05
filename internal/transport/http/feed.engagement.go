package http

import (
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/user"
)

// @Summary user feed based on engadgement
// @Description fetch questions based on recent engadgement
// @in header
// @Accept  json
// @Produce  json
// @Success 200 {array} question.Quesion
// @Router /user/feed [GET]
// @Tags Feed
func (h *Handler) GetEngagementFeed(w http.ResponseWriter, r *http.Request) {
	var u user.User
	u = r.Context().Value("user").(user.User)
	// run algorithm to get data
	values := h.FeedService.GetUserInterestedQuestionsEngagement(u)
	h.sendOkResponse(w, values)

}