package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

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
	isvalid := utils.CheckPasswordHash(loginreq.Password, user.Password)

	if !isvalid {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Invalid Username or password", errors.New("Invalid Username or Password"), 401)
		return
	}

	// check password lmao

	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Invalid Username or password", err, 401)
		return
	}

	jwt, expretime, err := utils.GenerateJWT(user)
	refreshtoken, refreshexp, err := utils.SendRefreshToken(user)

	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Unable to generate a JWT token", err, 500)
	}

	respwithtoken := struct {
		Token        string `json:"auth_token"`
		Expire       int64  `json:"expire"`
		RefreshToken string `json:"refresh_token"`
		RefreshExp   int64  `json:"refresh_expire"`
		UserRole     int    `json:"user_type"`
	}{
		Token:        jwt,
		Expire:       expretime,
		RefreshToken: refreshtoken,
		RefreshExp:   refreshexp,
		UserRole:     user.UserType,
	}

	h.sendOkResponse(w, respwithtoken)

}
