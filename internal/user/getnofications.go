package user

// get nofication sof a given user
func (s *Service) GetNofications(user User) []Nofication {
	var u User
	// maybe we can add something like read or not here but later
	s.DB.Debug().Preload("Nofications").Where("id =? ", user.ID).Find(&u)
	return u.Nofications
}
