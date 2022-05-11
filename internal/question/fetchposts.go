package question

import "fmt"

// fetch all posts avaible
func (s *Service) GetAllQuestions() []*Question {
	var questions []*Question
	s.DB.Debug().Preload("Tags").Preload("Answers").Preload("UpvotedUsers").Order("created_at DESC").Preload("Related").Find(&questions)
	return questions
}

func (s *Service) GetQuestionBasedonSlug(slug string) Question {
	var question Question
	s.DB.Debug().Preload("Tags").Preload("Answers").Preload("UpvotedUsers").Order("created_at DESC").Preload("Related").Where("slug = ?", slug).Find(&question)

	return question
}

func (s *Service) GetQuestionsBasedonTags(vals []string) []Question {
	var questions []Question
	s.DB.Debug().
		Raw("select * from questions where id in (select DISTINCT question_id from question_related where top_word_id in (select id from top_words where word in (?)))", vals).
		Scan(&questions)

	for _, q := range questions {
		var tags []Tag
		s.DB.Debug().Raw("select * from tags where id IN (select tag_id from question_tags where question_id = ?)", q.ID).Scan(&tags)
		q.Tags = tags
	}

	return questions
}

// GetArticles returns all articles from the database
func (s *Service) GetQuestionsPaginate(pageID int) PaginatedQuestions {
	var response PaginatedQuestions
	var questions []Question
	var last_id int

	fmt.Println("============== page id paginate =================")
	fmt.Println(pageID)
	s.DB.Debug().Where("id >= ? ", pageID).Order("id").Limit(pageSize + 1).Find(&questions)
	s.DB.Debug().Raw("SELECT id FROM questions ORDER BY id DESC LIMIT 1").Find(&last_id)

	// db.Model(&User{}).Group("name").Count(&count)
	/*
		if len(questions) == pageSize {
			next_id := questions[len(response.Items)-1].ID
			uq := questions[:pageSize]
			response.Items = append(response.Items, uq...)
			response.Next = int(next_id)
		}
	*/

	response.Items = questions
	response.Count = int64(last_id)
	response.Next = int(questions[len(questions)-1].ID)
	return response
}
