package http

import (
	"fmt"
	"net/http"
	"os"

	b64 "encoding/base64"

	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

func (h *Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	secret_encoded := r.URL.Query().Get("secret")

	secret, err := b64.StdEncoding.DecodeString(secret_encoded)

	if err != nil {
		h.sendErrorResponse(w, "Unable to decode the secret", err, 500)
		return
	}

	tokeninfo := h.VerificationServie.GetTokenInfo(token)
	user, err := h.UserService.GetUserbyOnlyID(uint(tokeninfo.UserId))

	if err != nil {
		h.sendErrorResponse(w, "Unable to find user for that token", err, 401)
		return
	}

	err, success := h.VerificationServie.RequestVerification(user, tokeninfo, string(secret))

	if err != nil {
		h.sendErrorResponse(w, "Unable to complete email verification", err, 500)
		return
	}

	fmt.Println(token, string(secret))
	h.sendOkResponse(w, struct {
		Success bool `json:"success"`
	}{
		Success: success,
	})

}

func (h *Handler) RequestVerification(w http.ResponseWriter, r *http.Request) {
	var u user.User
	var apirUrl = []byte(os.Getenv("API_URL"))

	u = r.Context().Value("user").(user.User)

	err, token, secret := h.VerificationServie.GenerateVerification(u)

	if err != nil {
		h.sendErrorResponse(w, "Unable to generate the secret", err, 400)
		return
	}

	// SendEmailWithToken
	submitUrl := string(apirUrl) + "/user/confirm?token=" + token + "&secret=" + b64.StdEncoding.EncodeToString([]byte(secret))
	utils.SendEmailWithToken(u.Email, submitUrl)

	h.sendOkResponse(w, struct {
		Token  string `json:"token"`
		Secret string `json:"secret"`
	}{
		Token:  token,
		Secret: b64.StdEncoding.EncodeToString([]byte(secret)),
	})
}
