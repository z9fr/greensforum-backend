package user

import "fmt"

func (s *Service) IsEmailExists(email string) bool {
	var exists bool
	if err := s.DB.Debug().Model(User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		fmt.Println(err)
	}

	return exists
}

func (s *Service) IsUserNameExists(username string) bool {
	var exists bool
	if err := s.DB.Debug().Model(User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		fmt.Println(err)
	}

	return exists
}
