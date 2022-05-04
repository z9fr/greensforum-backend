package http

import (
	"net/http"
)

//@TODO add pagiation
func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.UserService.FetchallUsers()
	h.sendOkResponse(w, users)
}
