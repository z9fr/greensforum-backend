package collective

import (
	"errors"

	"github.com/z9fr/greensforum-backend/internal/user"
)

func (s *Service) CreatePostinCollective(post Post, u user.User, collective_slug string) (Collective, []user.User, user.Nofication, error, bool) {
	collective := s.GetCollectiveBySlug(collective_slug)

	if collective.Slug == "" || collective.Name == "" {
		return Collective{}, []user.User{}, user.Nofication{}, errors.New("not a valid collective"), false
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
