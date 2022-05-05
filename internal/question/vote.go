package question

import (
	"fmt"

	"github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

// upvote a posttopwords
func (s *Service) UpVotePost(user *user.User, question *Question) (bool, []types.TopWord) {

	if question.Title == "" {
		return false, []types.TopWord{}
	}
	// check if user is already upvoted
	voted := s.isUpvoted(user, question)
	if voted {
		utils.LogWarn("user already upvoted")
		return false, []types.TopWord{}
	} else {
		question.UpVoteCount++
		upuser := UpVotedBy{
			UserId: uint(user.ID),
		}
		question.UpvotedUsers = append(question.UpvotedUsers, upuser)
		s.DB.Debug().Save(&question)
	}

	return true, question.Related
}

func (s *Service) isUpvoted(user *user.User, question *Question) bool {
	for _, upvoteduser := range question.UpvotedUsers {

		if upvoteduser.UserId == uint(user.ID) {
			fmt.Println(upvoteduser.UserId, user.ID)
			return true
		}
	}
	return false
}
