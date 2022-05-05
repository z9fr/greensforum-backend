package collective

import (
	"errors"

	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

//@TODO
// can optimize this code by removing these return values
func (s *Service) CreatePostinCollective(post Post, u user.User, collective_slug string) (Collective, []user.User, user.Nofication, error, bool) {
	collective := s.GetCollectiveBySlug(collective_slug)

	post.Slug = utils.GenerateSlug(post.Title)
	// check if post slug is alreay taken

	if s.IsPostSlugExist(post.Slug) {
		return Collective{}, []user.User{}, user.Nofication{}, errors.New("this post title with the same slug already exist. "), false
	}

	if post.Slug == "" || post.Title == "" {
		return Collective{}, []user.User{}, user.Nofication{}, errors.New("please fill all the required values"), false
	}

	if !s.IsCollectiveMember(collective, u) {
		return Collective{}, []user.User{}, user.Nofication{}, errors.New("you need to join the collective first"), false
	}

	var nf user.Nofication

	nf.Message = "new post is waiting for aproval view post at <a href=/posts/unaproved/" + post.Slug + "> here </a>"

	collective.Post = append(collective.Post, post)
	s.DB.Debug().Save(&collective)

	return *collective, collective.Admins, nf, nil, true
}
