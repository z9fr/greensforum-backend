package collective

import "github.com/z9fr/greensforum-backend/internal/utils"

// @Desctiption
// check if a collection slug is
// already taken
func (s *Service) IsCollectiveSlugExist(slug string) bool {
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

// @Desctiption
// check if a collective name is already taken
// already taken
func (s *Service) IsCollectiveNameExist(title string) bool {
	var exists bool
	if err := s.DB.Debug().Model(&Collective{}).
		Select("count(*) > 0").
		Where("name = ?", title).
		Find(&exists).
		Error; err != nil {
		utils.LogWarn(err)
	}

	return exists
}

// @Desctiption
// check if a post slug is already taken
func (s *Service) IsPostSlugExist(slug string) bool {
	var exists bool
	if err := s.DB.Debug().Model(&Post{}).
		Select("count(*) > 0").
		Where("slug = ?", slug).
		Find(&exists).
		Error; err != nil {
		utils.LogWarn(err)
	}

	return exists
}
