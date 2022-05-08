package collective

// @TODO
// support pagincation
func (s *Service) GetAllCollectives() []*Collective {
	var collectives []*Collective
	s.DB.Debug().Preload("Admins").Preload("Members").Order("created_at DESC").Find(&collectives)

	/*
		for _, collective := range collectives {
			s.DB.Debug().Preload("UserAcc").Find(&collective.Admins)
			s.DB.Debug().Preload("UserAcc").Find(&collective.Members)
		}
	*/

	return collectives
}

func (s *Service) GetCollectiveBySlug(slug string) *Collective {
	var collective *Collective
	// only get the accepted posts
	s.DB.Debug().Preload("Admins").Preload("Members").Preload("Post", "is_accepted = ?", true).Order("created_at DESC").Where("slug = ?", slug).First(&collective)

	/*
		s.DB.Debug().Preload("UserAcc").Find(&collective.Members)
		s.DB.Debug().Preload("UserAcc").Find(&collective.Admins)
		s.DB.Debug().Preload("Comments").Find(&collective.Post)

	*/

	return collective
}
