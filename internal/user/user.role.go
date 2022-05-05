package user

func (s *Service) IsHighPriv(user User) bool {

	/*
		    u, err := s.GetUserByEmail(user.Email)

			if err != nil {
				fmt.Println(err)
			}
	*/

	if user.UserType > 0 {
		return true
	}
	return false
}

func (s *Service) IsAdmin(user User) bool {

	/*
		u, err := s.GetUserByEmail(user.Email)

		if err != nil {
			fmt.Println(err)
		}
	*/
	if user.UserType == 2 {
		return true
	}
	return false
}
