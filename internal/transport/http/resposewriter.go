package http

import (
	"encoding/json"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/types"
)

// handle success response
func (h *Handler) sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

// handle error responses
func (h *Handler) sendErrorResponse(w http.ResponseWriter, message string, err error, errorcode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorcode)

	if err := json.NewEncoder(w).Encode(types.ErrorResponse{
		Error:   message,
		Details: err.Error(),
	}); err != nil {
		panic(err)
	}
}
