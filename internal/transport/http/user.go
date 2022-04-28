package http

import (
	"encoding/json"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/helper"
	"github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	u, validationerror := helper.RequestToUserWithValidations(user)
	if validationerror != nil {
		h.sendErrorResponse(w, "Validation Error Occured", validationerror, 500)
		return
	}
	createdUser, err := h.UserService.CreateUser(u)
	if err != nil {
		h.sendErrorResponse(w, "Unable to create a user please try again", err, 500)
		return
	}

	h.sendOkResponse(w, createdUser)
	return
}

//@TODO add pagiation
func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.UserService.FetchallUsers()
	h.sendOkResponse(w, users)
}

// @Login
// Login Request
// users will send a payload of email and password
// validate the user details and generate a JWT Token
// and send the response
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginreq types.Login

	if err := json.NewDecoder(r.Body).Decode(&loginreq); err != nil {
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	user, err := h.UserService.GetUserByEmail(loginreq.Email)
	if err != nil {
		h.sendErrorResponse(w, "Unable to find user with that email", err, 500)
		return
	}

	jwt, expretime, err := utils.GenerateJWT(user)

	if err != nil {
		h.sendErrorResponse(w, "Unable to generate a JWT token", err, 500)
	}

	respwithtoken := struct {
		Token  string `json:"token"`
		Expire int64  `json:"expire"`
	}{
		Token:  jwt,
		Expire: expretime,
	}

	h.sendOkResponse(w, respwithtoken)

}
