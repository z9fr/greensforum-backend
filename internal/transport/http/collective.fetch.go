package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// @TODO - support pagination

// @Summary Get all Collectives
// @Description get all collectives
// @in header
// @Accept  json
// @Produce  json
// @Success 200 {array} collective.Collective
// @Router /collectives [GET]
// @Tags Collective
func (h *Handler) FetchCollectives(w http.ResponseWriter, r *http.Request) {
	colletives := h.CollectiveService.GetAllCollectives()
	h.sendOkResponse(w, colletives)
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

	if !h.CollectiveService.IsUniqueSlug(collective_slug) {
		h.sendErrorResponse(w, "404 not found", errors.New("collective not found"), http.StatusNotFound)
		return
	}

	c := h.CollectiveService.GetCollectiveBySlug(collective_slug)
	h.sendOkResponse(w, c)

}
