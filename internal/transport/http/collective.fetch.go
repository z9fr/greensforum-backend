package http

import "net/http"

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

// view unaproved posts in a collection
