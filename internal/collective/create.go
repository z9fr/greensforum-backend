package collective

import "errors"

//@Description
// create a new collection based on given details
func (s *Service) CreateNewCollective(collective Collective) (Collective, error) {
	if s.IsCollectiveNameExist(collective.Name) {
		return Collective{}, errors.New("Collective name is already taken")

	}

	if s.IsCollectiveSlugExist(collective.Slug) {
		return Collective{}, errors.New("Collective slug (url) is already taken.")
	}

	return collective, nil
}
