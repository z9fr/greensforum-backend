package question

// update question views
func (s *Service) UpdateQuestionViews(question Question) {
	question.ViewCount++
	s.DB.Debug().Save(&question)
}
