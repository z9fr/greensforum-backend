package http

import (
	"encoding/json"
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
	if err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "Unable to find user with that email", err, 500)
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
	}{
		Token:        jwt,
		Expire:       expretime,
		RefreshToken: refreshtoken,
		RefreshExp:   refreshexp,
	}

	h.sendOkResponse(w, respwithtoken)

}
