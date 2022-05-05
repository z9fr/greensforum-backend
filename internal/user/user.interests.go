package user

import (
	"fmt"

	"github.com/z9fr/greensforum-backend/internal/types"
)

func (s *Service) AppendUserInterests(user User, interests []types.TopWord) bool {

	if !s.IsUserNameExists(user.Username) {
		return false
	}

	for _, interest := range interests {
		user.Interests = append(user.Interests, interest)
	}

	fmt.Println(user)
	s.DB.Debug().Save(&user)

	return true
}
