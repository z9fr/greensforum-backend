package http

import (
	"encoding/json"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/helper"
	"github.com/z9fr/greensforum-backend/internal/user"
)

// @Summary Register
// @Description register a new user
// @Accept  json
// @Produce  json
// @Param payload body user.CreateUserRequest true "payload"
// @Success 200 {object} user.User
// @Router /user/join [POST]
// @Tags User
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	u, validationerror := helper.RequestToUserWithValidations(user)
	if validationerror != nil {
		LogWarningsWithRequestInfo(r, validationerror)
		h.sendErrorResponse(w, "Validation Error Occured", validationerror, 500)
		return
	}
	createdUser, err := h.UserService.CreateUser(u)
	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Unable to create a user please try again", err, 500)
		return
	}

	h.sendOkResponse(w, createdUser)
	return
}

func (h *Handler) CreateAdminUser(w http.ResponseWriter, r *http.Request) {
	var user user.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	u, validationerror := helper.RequestToUserWithValidations(user)
	if validationerror != nil {
		LogWarningsWithRequestInfo(r, validationerror)
		h.sendErrorResponse(w, "Validation Error Occured", validationerror, 500)
		return
	}
	createdUser, err := h.UserService.CreateAdminUser(u)
	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Unable to create a user please try again", err, 500)
		return
	}

	h.sendOkResponse(w, createdUser)
	return
}
