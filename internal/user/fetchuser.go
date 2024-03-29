package user

import "fmt"

func (s *Service) FetchallUsers() []User {
	var users []User
	s.DB.Debug().Preload("UserAcc").Find(&users)
	return users
}

func (s *Service) GetUserByEmail(email string) (User, error) {
	var user User

	if !s.IsEmailExists(email) {
		return User{}, fmt.Errorf("No user with that email")
	}

	if result := s.DB.Debug().Preload("UserAcc").Preload("Nofications").First(&user, "email = ?", email); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (s *Service) GetUserbyID(id uint, email string) (User, error) {
	var user User

	if !s.IsEmailExists(email) {
		return User{}, fmt.Errorf("No user with that email")
	}

	// do not preload noficiations
	if result := s.DB.Debug().Preload("UserAcc").First(&user, "id = ?", id); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (s *Service) GetUserbyOnlyID(id uint) (User, error) {
	var user User

	// do not preload noficiations
	if result := s.DB.Debug().Preload("UserAcc").Preload("Nofications").First(&user, "id = ?", id); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
