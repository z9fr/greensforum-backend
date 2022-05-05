package collective

import "github.com/z9fr/greensforum-backend/internal/user"

// return true if a user is a member of a collective
func (s *Service) IsCollectiveMember(collective *Collective, user user.User) bool {
	isexist := false

	for _, u := range collective.Members {
		if u.Username == user.Username {
			isexist = true
		}
	}

	return isexist
}

// return true if a user is a admin of a collective
func (s *Service) IsCollectiveAdmin(collective *Collective, user user.User) bool {
	isexist := false

	for _, u := range collective.Admins {
		if u.Username == user.Username {
			isexist = true
		}
	}

	return isexist
}
