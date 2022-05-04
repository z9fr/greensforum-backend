package http

import "net/http"

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
	questions := h.QuestionService.GetAllQuestions()

	for _, question := range questions {
		defer h.TopWordsService.Reset()
		question.Related = append(question.Related, h.TopWordsService.TopTenWords(question.Title)...)
		// question.RelatedTopics = append(question.RelatedTopics, string(h.TopWordsService.TopTenWords(question.Body)))
	}

	h.sendOkResponse(w, questions)
	return
}

// =========== find posts pagination test =========================
func (h *Handler) ListArticles(w http.ResponseWriter, r *http.Request) {
	pageID := r.Context().Value(PageIDKey).(int)
	questions := h.QuestionService.GetQuestionsPaginate(pageID)
	h.sendOkResponse(w, questions)
}
