package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z9fr/greensforum-backend/internal/user"
)

func (h *Handler) Joincollective(w http.ResponseWriter, r *http.Request) {
	collective_slug := chi.URLParam(r, "collective")

	if !h.CollectiveService.IsUniqueSlug(collective_slug) {
		h.sendErrorResponse(w, "404 not found", errors.New("collective not found"), http.StatusNotFound)
		return
	}

	var u user.User
	u = r.Context().Value("user").(user.User)

	complete_user, _ := h.UserService.GetUserByEmail(u.Email)
	status := h.CollectiveService.JoinMembertoConnective(complete_user, collective_slug)
	h.sendOkResponse(w, struct {
		Status bool `json:"status"`
	}{
		Status: status,
	})

}
