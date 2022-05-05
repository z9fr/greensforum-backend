package http

import (
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/user"
)

// @Summary get notifications
// @Description Get nofications for the current user
// @Accept  json
// @Produce  json
// @Success 200 {array} user.Nofication
// @Router /user/nofications [GET]
// @Security JWT
// @Tags User
func (h *Handler) GetNofications(w http.ResponseWriter, r *http.Request) {
	var u user.User
	u = r.Context().Value("user").(user.User)
	nf := h.UserService.GetNofications(u)
	h.sendOkResponse(w, nf)
}
