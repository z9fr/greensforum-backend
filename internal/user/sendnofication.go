package user

func (s *Service) SendNofications(user User, nofication Nofication) {
	nofication.Read = false
	user.Nofications = append(user.Nofications, nofication)
	s.DB.Debug().Save(&user)
}
