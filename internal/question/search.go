package question

import "github.com/lib/pq"

func (s *Service) GetQuestionByID(id uint) Question {
	var question Question
	s.DB.Debug().Preload("Tags").Preload("Answers").Preload("UpvotedUsers").Preload("Related").Where("id = ?", id).Find(&question)

	return question
}

// seach for a post based on a keyword
func (s *Service) SearchPosts(q string) []Question {
	var questions []Question
	s.DB.Debug().Where("title LIKE ?", "%"+q+"%").Preload("Tags").Preload("Answers").Find(&questions)
	return questions
}

// get questions based on a tag
// @updates this is mostly a performance update
// someone from discord helped me to figure this out

/*
at first .Scan(dest interface{}) works like this .Scan(&struct.Field1, &struct.Field2)
Where in every field with their respective tags for the database counterparty, it automatically maps it to their location.
the reason why i switched to pointer type is that I can get their memory address
once they are parsed 1 by 1
or scanned*
that way on this snippet:

```go
for _, question := range questions {
        var tags []Tags
        s.DB.Debug().Raw("select * from tags where id IN (select tag_id from question_tags where question_id = ?)", question.Question_ID).Scan(&tags)
        question.Tags = tags
    }
```

I don't need to create a new variable to store questions with tags,
in which question in this case is *Question and I can directly modify it since it is already a pointer.
that's why I directly inserted to question.Tags = tags
yes and everytime you .Scan(&dest) it automatically populates the fields

take note that there is a difference between *[]Question and []*Question

```go
var questions *[]Question

// questions is nil by default
// whereas

var questions []*Question

// `questions` is a zero-empty array by default
```

*/

func (s *Service) SearchQuestionsByTagsv2(tag string) []*Question {
	var questions []*Question

	s.DB.Debug().Raw("select * from questions where id in (select question_id from question_tags where tag_id in (select id from tags where name= ?))", tag).Scan(&questions)

	for _, question := range questions {
		var tags []Tag
		s.DB.Debug().Raw("select * from tags where id IN (select tag_id from question_tags where question_id = ?)", question.ID).Scan(&tags)
		question.Tags = tags
	}

	if len(questions) == 0 {
		return []*Question{}
	}

	return questions
}

func (s *Service) SearchQuestionsByTagsv2PGStringArray(tags pq.StringArray) []*Question {
	var questions []*Question

	for _, tag := range tags {
		s.DB.Debug().Raw("select * from questions where id in (select question_id from question_tags where tag_id in (select id from tags where name= ?))", tag).Scan(&questions)
		for _, question := range questions {
			var tags []Tag
			s.DB.Debug().Raw("select * from tags where id IN (select tag_id from question_tags where question_id = ?)", question.ID).Scan(&tags)
			question.Tags = tags
		}
	}

	if len(questions) == 0 {
		return []*Question{}
	}

	return questions
}
