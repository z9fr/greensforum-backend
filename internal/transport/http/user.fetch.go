package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
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

func (h Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "id")

	u64, err := strconv.ParseUint(user_id, 10, 32)
	if err != nil {
		h.sendErrorResponse(w, "Invalid parameter", err, 500)
		utils.LogWarn(err)
		return
	}
	user, err := h.UserService.GetUserbyOnlyID(uint(u64))

	if err != nil {
		h.sendErrorResponse(w, "Unable to find the user", err, 401)
		return
	}

	h.sendOkResponse(w, user)
}
