package http

import (
	"errors"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

func (h *Handler) UserApplyforMod(w http.ResponseWriter, r *http.Request) {
	var u user.User
	u = r.Context().Value("user").(user.User)

	user, err := h.UserService.GetUserByEmail(u.Email)

	if err != nil {
		h.sendErrorResponse(w, "Unable to find that user", err, http.StatusBadRequest)
		return
	}

	if !user.IsVerified {
		h.sendErrorResponse(w, "Please verify your email first", errors.New("E-Mail verification failed"), http.StatusBadRequest)
		return
	}

	if user.UserType > 0 {
		h.sendErrorResponse(w, "User already has these privilages", errors.New("User already has admin privilages"), http.StatusBadRequest)
		return
	}

	// check if email is a internal mail
	if !utils.IsInternalMail(user.Email) {
		h.sendErrorResponse(w, "Not a internal email", errors.New("You need to have a internal email to get this privilages. please contact admins"), http.StatusBadRequest)
		return
	}
	user.UserType = 1
	h.UserService.DB.Debug().Save(&user)

	h.sendOkResponse(w, user)
}
