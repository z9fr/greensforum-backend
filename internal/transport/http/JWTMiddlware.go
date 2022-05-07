package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

func (h *Handler) JWTMiddlewhare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authheader := r.Header["Authorization"]

		if len(authheader) == 0 {
			h.sendErrorResponse(w, "Missing Authorization Header", fmt.Errorf("Authorization is required Header"), 401)
			return
		}

		authToken := strings.Split(authheader[0], " ")[1]
		decodedUser, err := utils.VerifyToken(authToken)

		user, err := h.UserService.GetUserByEmail(decodedUser.Email)

		if err != nil {
			LogWarningsWithRequestInfo(r, err)
			h.sendErrorResponse(w, "Unable to Verify JWT Token", err, 401)
			return
		}

		// create new context from `r` request context, and assign key `"user"`
		// to value of `"123"`
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) HighPrivilagesMiddlewhare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var u user.User
		u = r.Context().Value("user").(user.User)

		if !h.UserService.IsHighPriv(u) {
			LogWarningsWithRequestInfo(r, "low-priv user trying to create a collection : user -> "+u.Email)
			h.sendErrorResponse(w, "unable to create a collection", errors.New("not enough privilages"), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
