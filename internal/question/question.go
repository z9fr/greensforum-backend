package question

import (
	"errors"

	models "github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/user"
	utils "github.com/z9fr/greensforum-backend/internal/utils"
	"gorm.io/gorm"
)

const (
	pageSize = 10
)

type Service struct {
	DB *gorm.DB
}

type Question struct {
	models.Model
	// QuestionID    int      `gorm:"column:question_id primaryKey" json:"question_id"`
	Title         string   `gorm:"column:title" json:"title"`
	Body          string   `gorm:"column:body" json:"body"`
	IsAnswered    bool     `gorm:"column:is_answered default:false" json:"is_answered"`
	ViewCount     int      `gorm:"column:view_count default:0" json:"view_count"`
	DownVoteCount int      `gorm:"column:down_vote_count default:0" json:"down_vote_count"`
	UpVoteCount   int      `grom:"column:up_vote_count default:0" json:"up_vote_count"`
	AnswerCount   int      `gorm:"column:answer_count default:0" json:"answer_count"`
	Score         int      `gorm:"column:score default:0" json:"score"`
	Answers       []Answer `gorm:"foreignKey:question_id;id" json:"answers"`
	Tags          []Tag    `gorm:"many2many:question_tags" json:"tags"`
	CreatedBy     uint64   `gorm:"column:created_by" json:"created_by" `
	Slug          string   `gorm:"column:slug" json:"slug"`
	// UpvotedUsers  []UpVotedBy `gorm:"foreignKey:question_id;id" json:"upvotedUsers"`
}

type Answer struct {
	models.Model
	Title         string `gorm:"column:title" json:"title"`
	Body          string `gorm:"column:body" json:"body"`
	QuestionID    int64  `gorm:"column:question_id" json:"question_id"`
	ViewCount     int    `gorm:"column:view_count default:0" json:"view_count"`
	DownVoteCount int    `gorm:"column:down_vote_count default:0" json:"down_vote_count"`
	UpVoteCount   int    `grom:"column:up_vote_count default:0" json:"up_vote_count"`
	Score         int    `gorm:"column:score default:0" json:"score"`
}

type Tag struct {
	models.Model
	Name string `json:"name"`
}

type UpVotedBy struct {
	UserId int `json:"upvoted"`
}

type QuestionCreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Tags  []Tag  `json:"tags"`
}

type AnswerRequest struct {
	Title string `gorm:"column:title" json:"title"`
	Body  string `gorm:"column:body" json:"body"`
}

type PaginatedQuestions struct {
	Items []Question `json:"items"`
	Next  int        `json:"next_page_id"`
}

// QuestionService - interface for Question Service
type QuestionService interface {
	CreateNewQuestion(question Question) (Question, error)
	GetAllPosts() []Question
	SearchQuestionsByTags(tag string) []Question
	SearchPosts(q string) []Question
	CreateAnswer(answer Answer, question_id string) (Question, error)

	// pagination
	GetQuestionsPaginate(pageID int) []Question

	//utils
	GetQuestionByID(id uint) Question
	UpdateQuestionViews(id uint)
}

// NewService - create a instance of this service and return
// a pointer to the servie
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// create a new question
func (s *Service) CreateNewQuestion(question Question) (Question, error) {

	if s.IsTitleExist(question.Title) {
		return Question{}, errors.New(`Title already exist. please try something else.`)
	}

	if result := s.DB.Debug().Save(&question); result.Error != nil {
		utils.LogWarn(result.Error)
		return Question{}, result.Error
	}

	return question, nil
}

// fetch all posts avaible
func (s *Service) GetAllPosts() []Question {
	var questions []Question
	s.DB.Debug().Preload("Tags").Preload("Answers").Find(&questions)

	return questions
}

// seach for a post based on a keyword
func (s *Service) SearchPosts(q string) []Question {
	var questions []Question
	s.DB.Debug().Where("title LIKE ?", "%"+q+"%").Preload("Tags").Preload("Answers").Find(&questions)
	return questions
}

// upvote a post

func (s *Service) UpVotePost(user user.User, question Question) {

}

func (s *Service) CreateAnswer(answer Answer, question_id uint) (Question, error) {

	if !s.IsQuestionExist(question_id) {
		return Question{}, errors.New("Question not found")
	}

	question := s.GetQuestionByID(question_id)
	question.Answers = append(question.Answers, answer)
	question.AnswerCount++
	question.IsAnswered = true

	s.DB.Debug().Save(&question)

	return question, nil

}

func (s *Service) GetQuestionByID(id uint) Question {
	var question Question
	s.DB.Debug().Preload("Tags").Preload("Answers").First(&question).Where("id = ?", id)

	return question
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

// get questions based on a tag
// @depriciated
func (s *Service) SearchQuestionsByTags(tag string) []Question {
	var questions []Question
	var questions_with_tags []Question

	// s.DB.Debug().Raw("select * from questions inner join question_tags ON id=question_id inner join tags on question_tags.tag_id=tags.id where questions.id IN (SELECT question_id FROM question_tags WHERE tag_id IN (SELECT id FROM tags WHERE name= ?))", tag).Scan(&questions)
	s.DB.Debug().Raw("select * from questions where id in (select question_id from question_tags where tag_id in (select id from tags where name= ?))", tag).Scan(&questions)

	for _, question := range questions {
		q := s.GetQuestionByID(question.ID)
		questions_with_tags = append(questions_with_tags, q)
	}

	if questions == nil {
		return []Question{}
	}

	return questions_with_tags

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

// update question views
func (s *Service) UpdateQuestionViews(question Question) {
	question.ViewCount++
	s.DB.Debug().Save(&question)
}
