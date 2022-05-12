package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/events"
	"github.com/z9fr/greensforum-backend/internal/user"
)

func (h *Handler) CreateNewEvent(w http.ResponseWriter, r *http.Request) {
	var eventcontent events.Event

	var u user.User
	u = r.Context().Value("user").(user.User)

	complete_user, _ := h.UserService.GetUserByEmail(u.Email)

	if err := json.NewDecoder(r.Body).Decode(&eventcontent); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	eventcontent.CreatedBy = uint(complete_user.ID)

	if complete_user.UserType >= 1 {
		event, err := h.EventServie.CreateEvent(eventcontent)
		if err != nil {
			h.sendErrorResponse(w, "Unable to create a event", err, 500)
			return
		}

		h.sendOkResponse(w, event)
		return
	}

	h.sendErrorResponse(w, "Not enough privilages", errors.New("Not enough privilages to create a new event"), http.StatusUnauthorized)
	h.sendOkResponse(w, eventcontent)

}
