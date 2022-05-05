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

// mark the post as approved
func (s *Service) ApprovePosts(post_slug string, collective_slug string, u user.User) (Post, error, bool) {

	var c *Collective
	s.DB.Debug().Preload("Admins").First(&c)

	// make sure user is a admin

	if !s.IsCollectiveAdmin(c, u) {
		return Post{}, errors.New("Only Collective admins can accept a post"), false
	}

	post := s.GetPostBySlug(post_slug)

	if post.Title == "" || post.Body == "" {
		return Post{}, errors.New("Post not found"), false
	}

	// accpet the post
	post.IsAccepted = true
	s.DB.Debug().Save(&post)

	return post, nil, true
}

// fetch post by using slug
func (s *Service) GetPostBySlug(slug string) Post {
	var post Post
	s.DB.Debug().Preload("Comments").Find(&post).Where("slug = ?", slug)
	return post
}
