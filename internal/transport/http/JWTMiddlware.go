package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"

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
		user, err := utils.VerifyToken(authToken)

		if err != nil {
			h.sendErrorResponse(w, "Unable to Verify JWT Token", err, 401)
			return
		}

		// create new context from `r` request context, and assign key `"user"`
		// to value of `"123"`
		ctx := context.WithValue(r.Context(), "user", user)

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		//
		// note: context.Context values are nested, so any previously set
		// values will be accessible as well, and the new `"user"` key
		// will be accessible from this point forward.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
