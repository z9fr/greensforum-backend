package collective

import (
	"errors"

	"github.com/z9fr/greensforum-backend/internal/user"
)

func (s *Service) CreatePostinCollective(post Post, user user.User, collective_slug string) (Collective, error) {
	collective := s.GetCollectiveBySlug(collective_slug)

	if collective.Slug == "" || collective.Name == "" {
		return Collective{}, errors.New("not a valid collective")
	}

	if !s.IsCollectiveMember(collective, user) {
		return Collective{}, errors.New("you need to join the collective first")
	}

	collective.Post = append(collective.Post, post)

	return *collective, nil
}
