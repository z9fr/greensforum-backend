package http

import (
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/user"
)

//@TODO add pagiation
func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.UserService.FetchallUsers()
	h.sendOkResponse(w, users)
}

func (h *Handler) LoggedinUser(w http.ResponseWriter, r *http.Request) {
	var u user.User

	u = r.Context().Value("user").(user.User)
	h.sendOkResponse(w, u)

}
