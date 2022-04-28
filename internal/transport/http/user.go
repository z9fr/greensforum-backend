package http

import (
	"encoding/json"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/helper"
	"github.com/z9fr/greensforum-backend/internal/user"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	u, validationerror := helper.RequestToUserWithValidations(user)
	if validationerror != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(validationerror.Error()))
		return
	}
	createdUser, err := h.UserService.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdUser)
	return
}

// Fetch All Users
// @Summary Get All Users
// @Produce json
// @success 200
// @Router  /users/all [get]
func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.UserService.FetchallUsers()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
