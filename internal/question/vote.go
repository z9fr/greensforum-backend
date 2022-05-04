package question

import (
	"github.com/z9fr/greensforum-backend/internal/user"
)

// upvote a post
func (s *Service) UpVotePost(user *user.User, question *Question) {

	if question.Title == "" {
		return
	}
	// check if user is already upvoted
	voted := s.isUpvoted(user, question)
	if voted {
		return
	} else {
		question.UpVoteCount++
		upuser := UpVotedBy{
			UserId: uint(user.ID),
		}
		question.UpvotedUsers = append(question.UpvotedUsers, upuser)
		s.DB.Debug().Save(&question)
	}

}

func (s *Service) isUpvoted(user *user.User, question *Question) bool {
	for _, u := range question.UpvotedUsers {
		if u.UserId == u.ID {
			return true
		}
	}

	return false
}
