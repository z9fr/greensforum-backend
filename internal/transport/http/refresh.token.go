package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/utils"
)

type RefreshReq struct {
	Token string `json:"refresh_token"`
}

type RefreshResponse struct {
	Auth       string `json:"auth"`
	Authexp    int64  `json:"auth_expire"`
	Refresh    string `json:"refresh"`
	Refreshexp int64  `json:"refresh_expire"`
	UserRole   int    `json:"user_type"`
}

// @Summary Refresh Tokens
// @Description refresh users token based on a given refresh token
// @Accept  json
// @Produce  json
// @Param payload body RefreshReq true "payload"
// @Success 200 {object} RefreshResponse
// @Router /user/refresh [POST]
// @Tags Authentication
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	var rinfo RefreshReq

	if err := json.NewDecoder(r.Body).Decode(&rinfo); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	user_id, user_email, token_version, err := utils.ValidateRefreshToken(rinfo.Token)

	if err != nil {
		h.sendErrorResponse(w, "Unable to validate refresh token", err, 401)
		return
	}

	u, err := h.UserService.GetUserbyID(uint(user_id.(float64)), user_email.(string))

	if err != nil {
		h.sendErrorResponse(w, "Unable to validate refresh token", err, 401)
		return
	}

	if int(u.TokenVersion) != int(token_version.(float64)) {
		h.sendErrorResponse(w, "Old refresh token", errors.New("This refresh token is revoked by the user"), http.StatusUnauthorized)
		return
	}

	refreshtoken, refreshexp, err := utils.SendRefreshToken(u)
	authtoken, authexp, err := utils.GenerateJWT(u)

	if err != nil {
		h.sendErrorResponse(w, "Unable to generate jwt", err, http.StatusInternalServerError)
		return
	}

	resp := &RefreshResponse{
		Auth:       authtoken,
		Authexp:    authexp,
		Refresh:    refreshtoken,
		Refreshexp: refreshexp,
		UserRole:   u.UserType,
	}

	h.sendOkResponse(w, resp)
}
