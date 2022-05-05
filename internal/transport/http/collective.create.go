package http

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/z9fr/greensforum-backend/internal/collective"
	"github.com/z9fr/greensforum-backend/internal/user"
)

// @Summary Create a new CreateCollective
// @Description creating a new collective
// @in header
// @Accept  json
// @Produce  json
// @Param payload body collective.Collective true "payload"
// @Success 200 {object} collective.Collective
// @Router /collective/create [POST]
// @Security JWT
// @Tags Collective
func (h *Handler) CreateCollective(w http.ResponseWriter, r *http.Request) {

	var data collective.Collective
	var u user.User
	u = r.Context().Value("user").(user.User)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, http.StatusInternalServerError)
		return
	}
	// add the created usr as a admin and as a member
	data.CreatedBy = uint(u.ID)
	data.Admins = append(data.Admins, u)
	data.Members = append(data.Members, u)

	c, err := h.CollectiveService.CreateNewCollective(data)

	if err != nil {
		h.sendErrorResponse(w, "unable to create collection", err, http.StatusInternalServerError)
		return
	}

	h.sendOkResponse(w, c)
}
