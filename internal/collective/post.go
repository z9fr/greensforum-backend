package collective

import (
	"errors"

	"github.com/z9fr/greensforum-backend/internal/user"
)

func (s *Service) GetUnaprovtedPosts(collective_slug string, u user.User) ([]Post, error, bool) {
	var collective *Collective
	s.DB.Debug().Preload("Admins").Preload("Members").Preload("Post", "is_accepted = ?", false).Where("slug = ?", collective_slug).First(&collective)

	// hide the unaproved posts for
	// normal users
	if !s.IsCollectiveAdmin(collective, u) {
		return []Post{}, errors.New("only collection admins can view unaproved posts"), false
	}

	return collective.Post, nil, true
}
