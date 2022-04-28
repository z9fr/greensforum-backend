package question

import "fmt"

func (s *Service) IsTitleExist(title string) bool {
	var exists bool
	if err := s.DB.Debug().Model(&Question{}).
		Select("count(*) > 0").
		Where("title = ?", title).
		Find(&exists).
		Error; err != nil {
		fmt.Println(err)
	}

	return exists
}
