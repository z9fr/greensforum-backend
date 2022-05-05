package collective

import (
	"errors"
	"fmt"
)

//@Description
// create a new collection based on given details
func (s *Service) CreateNewCollective(collective Collective) (Collective, error) {

	if collective.Name == "" || collective.Slug == "" {
		return Collective{}, errors.New("Slug or Name can't be empty")
	}

	if s.IsCollectiveNameExist(collective.Name) {
		fmt.Println("collective name exist")
		return Collective{}, errors.New("Collective name is already taken")
	}

	if s.IsCollectiveSlugExist(collective.Slug) {
		fmt.Println("slug")
		return Collective{}, errors.New("Collective slug (url) is already taken.")
	}

	s.DB.Debug().Save(&collective)
	return collective, nil
}
