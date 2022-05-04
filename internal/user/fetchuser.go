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

	if result := s.DB.Debug().Preload("UserAcc").First(&user, "email = ?", email); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
