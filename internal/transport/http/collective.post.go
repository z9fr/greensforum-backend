package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z9fr/greensforum-backend/internal/collective"
	"github.com/z9fr/greensforum-backend/internal/user"
)

// @Summary write a post
// @Description create a post in collective
// @in header
// @Accept  json
// @Produce  json
// @Param   collective   path  string  true  "collective slug"
// @Success 200 {object} collective.Post
// @Router /collectives/{collective}/post/write [POST]
// @Tags Collectives
func (h *Handler) WritePostinCollective(w http.ResponseWriter, r *http.Request) {
	var postcontent collective.Post
	collective_slug := chi.URLParam(r, "collective")

	if !h.CollectiveService.IsUniqueSlug(collective_slug) {
		h.sendErrorResponse(w, "404 not found", errors.New("collective not found"), http.StatusNotFound)
		return
	}

	var u user.User
	u = r.Context().Value("user").(user.User)

	complete_user, _ := h.UserService.GetUserByEmail(u.Email)

	if err := json.NewDecoder(r.Body).Decode(&postcontent); err != nil {
		LogWarningsWithRequestInfo(r, err)
		h.sendErrorResponse(w, "unable to decode json body", err, 500)
		return
	}

	// set these values by default
	postcontent.UpVoteCount = 0
	postcontent.DownVoteCount = 0
	postcontent.IsAccepted = false

	collective, admins, nofication, err, success := h.CollectiveService.CreatePostinCollective(postcontent, complete_user, collective_slug)

	if err != nil {
		h.sendErrorResponse(w, "Unable to create a post", err, http.StatusInternalServerError)
		return
	}

	if !success {
		h.sendErrorResponse(w, "Unable to create a post", errors.New("Something went wrong while creating the post please contact admins"), http.StatusInternalServerError)
		return
	}

	// send nofications for admins
	for _, admin := range admins {
		h.UserService.SendNofications(admin, nofication)
	}

	h.sendOkResponse(w, collective)

}

// @Summary get collective by slug
// @Description get infromation about a collecting using slug
// @in header
// @Accept  json
// @Produce  json
// @Param   collective   path  string  true  "collective slug"
// @Success 200 {object} collective.Collective
// @Router /collectives/{collective} [GET]
// @Tags Collectives
func (h *Handler) GetCollectiveBySlug(w http.ResponseWriter, r *http.Request) {
	collective_slug := chi.URLParam(r, "collective")

	if !h.CollectiveService.IsUniqueSlug(collective_slug) {
		h.sendErrorResponse(w, "404 not found", errors.New("collective not found"), http.StatusNotFound)
		return
	}

	c := h.CollectiveService.GetCollectiveBySlug(collective_slug)
	h.sendOkResponse(w, c)

}
