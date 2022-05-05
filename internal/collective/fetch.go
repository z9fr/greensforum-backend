package collective

// @TODO
// support pagincation
func (s *Service) GetAllCollectives() []Collective {
	var collectives []Collective
	s.DB.Debug().Find(&collectives)
	return collectives
}
