package collective

import (
	"errors"
	"fmt"

	"github.com/z9fr/greensforum-backend/internal/utils"
)

//@Description
// create a new collection based on given details
func (s *Service) CreateNewCollective(collective Collective) (Collective, error) {

	collective.Slug = utils.GenerateSlug(collective.Slug)
	// make sure slug is unique

	if collective.Name == "" || collective.Slug == "" {
		return Collective{}, errors.New("Slug or Name can't be empty")
	}

	if s.IsCollectiveNameExist(collective.Name) {
		fmt.Println("collective name exist")
		return Collective{}, errors.New("Collective name is already taken")
	}

	if s.IsCollectiveSlugExist(collective.Slug) {
		return Collective{}, errors.New("Collective slug (url) is already taken.")
	}

	s.DB.Debug().Save(&collective)
	return collective, nil
}

func (s *Service) IsUniqueSlug(slug string) bool {
	var exists bool
	if err := s.DB.Debug().Model(&Collective{}).
		Select("count(*) > 0").
		Where("slug = ?", slug).
		Find(&exists).
		Error; err != nil {
		utils.LogWarn(err)
	}

	return exists
}
