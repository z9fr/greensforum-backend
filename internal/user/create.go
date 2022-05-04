package user

import "fmt"

func (s *Service) CreateUser(user User) (User, error) {

	if s.IsEmailExists(user.Email) {
		return User{}, fmt.Errorf("Email is Already Taken")
	}

	if s.IsUserNameExists(user.Username) {
		return User{}, fmt.Errorf("Email is Already Taken")
	}

	if result := s.DB.Debug().Save(&user); result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
