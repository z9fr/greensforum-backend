package question

// fetch all posts avaible
func (s *Service) GetAllPosts() []Question {
	var questions []Question
	s.DB.Debug().Preload("Tags").Preload("Answers").Preload("UpvotedUsers").Find(&questions)

	return questions
}

// GetArticles returns all articles from the database
func (s *Service) GetQuestionsPaginate(pageID int) PaginatedQuestions {
	var response PaginatedQuestions
	var questions []Question
	s.DB.Where("id >= ? ", pageID).Order("id").Limit(pageSize + 1).Find(&questions)

	if len(questions) == pageSize+1 {
		next_id := questions[len(response.Items)-1].ID
		uq := questions[:pageSize]
		response.Items = append(response.Items, uq...)
		response.Next = int(next_id)
	}

	return response
}
