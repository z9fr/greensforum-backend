package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

const (
	// PageIDKey refers to the context key that stores the next page id
	PageIDKey string = "page_id"
)

// Pagination middleware is used to extract the next page id from the url query
func (h *Handler) Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		PageID := r.URL.Query().Get(string(PageIDKey))
		intPageID := 0
		var err error
		if PageID != "" {
			intPageID, err = strconv.Atoi(PageID)
			if err != nil {
				h.sendErrorResponse(w, "No Page Found. ", fmt.Errorf("couldn't read %s: %w", PageIDKey, err), 500)
				return
			}
		}
		ctx := context.WithValue(r.Context(), PageIDKey, intPageID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
