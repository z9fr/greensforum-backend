package http

import (
	"encoding/json"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/helper"
	"github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
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

// @Summary Authenticate User
// @Description login as a existing user
// @Accept  json
// @Produce  json
// @Param payload body types.Login true "payload"
// @Success 200 {object} types.AuthRequest
// @Router /user/login [POST]
// @Tags Authentication
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginreq types.Login

	if err := json.NewDecoder(r.Body).Decode(&loginreq); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	user, err := h.UserService.GetUserByEmail(loginreq.Email)
	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Unable to find user with that email", err, 500)
		return
	}

	jwt, expretime, err := utils.GenerateJWT(user)

	if err != nil {
		LogWarningsWithRequestInfo(r, err)
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
