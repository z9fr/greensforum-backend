package question

import (
	"github.com/z9fr/greensforum-backend/internal/utils"
)

func (s *Service) IsTitleExist(title string) bool {
	var exists bool
	if err := s.DB.Debug().Model(&Question{}).
		Select("count(*) > 0").
		Where("title = ?", title).
		Find(&exists).
		Error; err != nil {
		utils.LogWarn(err)
	}

	return exists
}

func (s *Service) IsQuestionExist(id uint) bool {
	var exists bool
	if err := s.DB.Debug().Model(&Question{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error; err != nil {
		utils.LogWarn(err)
	}

	return exists
}
