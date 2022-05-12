package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) ViewallEvents(w http.ResponseWriter, r *http.Request) {
	events := h.EventServie.GetAllEvents()
	h.sendOkResponse(w, events)
}

func (h *Handler) GeteventsBySlug(w http.ResponseWriter, r *http.Request) {
	event_slug := chi.URLParam(r, "event")

	isexist := h.EventServie.IsEventSlugExist(event_slug)

	if !isexist {
		h.sendErrorResponse(w, "event not found", errors.New("No events found for this slug"), http.StatusOK)
		return
	}

	event := h.EventServie.GeteventbySlug(event_slug)
	h.sendOkResponse(w, event)
}
