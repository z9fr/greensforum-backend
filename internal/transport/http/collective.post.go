package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z9fr/greensforum-backend/internal/collective"
	"github.com/z9fr/greensforum-backend/internal/user"
)

/*
@TODO
maybe add the collective check to a middlewhare
*/

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

// @Summary view unaproved posts
// @Description list all unaproved posts in a collective
// @Accept  json
// @Produce  json
// @Param   collective   path  string  true  "collective slug"
// @Success 200 {object} collective.Post
// @Router /collectives/{collective}/unaproved [GET]
// @Security JWT
// @Tags Collectives
func (h *Handler) ViewUnaprovedPosts(w http.ResponseWriter, r *http.Request) {

	collective_slug := chi.URLParam(r, "collective")

	var u user.User
	u = r.Context().Value("user").(user.User)

	if !h.CollectiveService.IsUniqueSlug(collective_slug) {
		h.sendErrorResponse(w, "404 not found", errors.New("collective not found"), http.StatusNotFound)
		return
	}

	posts, err, success := h.CollectiveService.GetUnaprovtedPosts(collective_slug, u)

	if err != nil {
		h.sendErrorResponse(w, "Unable to get unaproved Posts", err, http.StatusInternalServerError)
		return
	}

	if !success {
		h.sendErrorResponse(w, "Unable to get unaproved Posts", errors.New("unknown errors occured"), http.StatusInternalServerError)
		return
	}

	h.sendOkResponse(w, posts)
}

// @Summary approve post
// @Description approve post
// @Accept  json
// @Produce  json
// @Param   collective   path  string  true  "collective slug"
// @Param   post   path  string  true  "post slug"
// @Success 200 {object} collective.Post
// @Router /collectives/{collective}/{post}/approve [POST]
// @Security JWT
// @Tags Collectives
func (h *Handler) ApprovePost(w http.ResponseWriter, r *http.Request) {

	collective_slug := chi.URLParam(r, "collective")
	post_slug := chi.URLParam(r, "post")

	var u user.User
	u = r.Context().Value("user").(user.User)

	if !h.CollectiveService.IsUniqueSlug(collective_slug) {
		h.sendErrorResponse(w, "404 not found", errors.New("collective "+collective_slug+" not found"), http.StatusNotFound)
		return
	}

	if !h.CollectiveService.IsPostSlugExist(post_slug) {
		h.sendErrorResponse(w, "404 not found", errors.New("cant find "+post_slug+" in collective "+collective_slug), http.StatusNotFound)
		return
	}

	post, err, _ := h.CollectiveService.ApprovePosts(post_slug, collective_slug, u)

	if err != nil {
		h.sendErrorResponse(w, "Unable to approve post", err, http.StatusInternalServerError)
		return
	}

	h.sendOkResponse(w, post)
}

// @Summary view post
// @Description view post details
// @Accept  json
// @Produce  json
// @Param   post   path  string  true  "post slug"
// @Success 200 {object} collective.Post
// @Router /post/{post} [GET]
// @Tags Collectives
func (h *Handler) GetPostbySlug(w http.ResponseWriter, r *http.Request) {

	post_slug := chi.URLParam(r, "post")

	if !h.CollectiveService.IsPostSlugExist(post_slug) {
		h.sendErrorResponse(w, "404 not found", errors.New("cant find "+post_slug), http.StatusNotFound)
		return
	}

	post := h.CollectiveService.GetPostBySlug(post_slug)

	h.sendOkResponse(w, post)
}

func (h *Handler) GetmyTobeApprovePosts(w http.ResponseWriter, r *http.Request) {

	var u user.User
	u = r.Context().Value("user").(user.User)
	posts := h.CollectiveService.GetMyUnaprovedPosts(uint(u.ID))
	h.sendOkResponse(w, posts)
}
