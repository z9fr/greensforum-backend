package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/z9fr/greensforum-backend/internal/collective"
)

func (h *Handler) WritePostinCollective(w http.ResponseWriter, r *http.Request) {

}

// @Summary get collective by slug
// @Description get infromation about a collecting using slug
// @in header
// @Accept  json
// @Produce  json
// @Param   collective   path  string  true  "collective slug"
// @Success 200 {object} collective.Collective
// @Router /collectives/{collective} [GET]
// @Tags Collectives
func (h *Handler) GetCollectiveBySlug(w http.ResponseWriter, r *http.Request) {
	collective_slug := chi.URLParam(r, "collective")
	c := h.CollectiveService.GetCollectiveBySlug(collective_slug)
	h.sendOkResponse(w, c)

}
