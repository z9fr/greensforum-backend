package collective

// @TODO
// support pagincation
func (s *Service) GetAllCollectives() []*Collective {
	var collectives []*Collective
	s.DB.Debug().Preload("Admins").Preload("Members").Find(&collectives)

	// fetch user accounts and stuff too
	for _, collective := range collectives {
		s.DB.Debug().Preload("UserAcc").Find(&collective.Admins)
		s.DB.Debug().Preload("UserAcc").Find(&collective.Members)
	}

	return collectives
}

func (s *Service) GetCollectiveBySlug(slug string) *Collective {
	var collective *Collective
	s.DB.Debug().Preload("Admins").Preload("Members").Find(&collective).Where("slug = ", slug)
	s.DB.Debug().Preload("UserAcc").Find(&collective.Members)
	s.DB.Debug().Preload("UserAcc").Find(&collective.Admins)

	return collective
}
