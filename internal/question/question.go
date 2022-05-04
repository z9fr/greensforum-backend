package question

import (
	models "github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/user"
	"gorm.io/gorm"

	pq "github.com/lib/pq"
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
	Title         string         `gorm:"column:title" json:"title"`
	Body          string         `gorm:"column:body" json:"body"`
	IsAnswered    bool           `gorm:"column:is_answered default:false" json:"is_answered"`
	ViewCount     int            `gorm:"column:view_count default:0" json:"view_count"`
	DownVoteCount int            `gorm:"column:down_vote_count default:0" json:"down_vote_count"`
	UpVoteCount   int            `grom:"column:up_vote_count default:0" json:"up_vote_count"`
	AnswerCount   int            `gorm:"column:answer_count default:0" json:"answer_count"`
	Score         int            `gorm:"column:score default:0" json:"score"`
	CreatedBy     uint64         `gorm:"column:created_by" json:"created_by" `
	Slug          string         `gorm:"column:slug" json:"slug"`
	Answers       []Answer       `gorm:"foreignKey:question_id;id" json:"answers"`
	Tags          []Tag          `gorm:"many2many:question_tags" json:"tags"`
	UpvotedUsers  []UpVotedBy    `gorm:"many2many:question_id;id" json:"upvotedUsers"`
	RelatedWorks  pq.StringArray `gorm:"type:varchar(64)[]" json:"relatedworks"`
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
	models.Model
	UserId uint `json:"uuid"`
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
	GetAllQuestions() []Question
	CreateAnswer(answer Answer, question_id string) (Question, error)
	SearchQuestions(q string) []Question
	SearchQuestionsByTags(tag string) []Question

	// pagination
	GetQuestionsPaginate(pageID int) []Question

	//utils
	GetQuestionByID(id uint) Question
	UpdateQuestionViews(id uint)

	// vote
	UpVotePost(user *user.User, question *Question)
	isUpvoted(user *user.User, question *Question) bool
}

// NewService - create a instance of this service and return
// a pointer to the servie
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
<<<<<<< HEAD

// CreateNewQuestion - Create a new question
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

// GetAllQuestions - get all posts
// and return a []Questions
func (s *Service) GetAllQuestions() []Question {
	var questions []Question
	s.DB.Debug().Preload("Tags").Preload("Answers").Find(&questions)

	return questions
}

// SearchQuestions - search posts on the database
// and return a Array of questions []Question
func (s *Service) SearchQuestions(q string) []Question {
	var questions []Question
	s.DB.Debug().Where("title LIKE ?", "%"+q+"%").Preload("Tags").Preload("Answers").Find(&questions)
	return questions
}

// GetQuestionByID - get question by id and return the question
func (s *Service) GetQuestionByID(id string) Question {
	var question Question
	s.DB.Debug().Preload("Tags").Preload("Answers").First(&question).Where("id = ?", id)

	return question
}

// SearchQuestionsByTags - search questions by Tag
// and returna Array of []Questions
func (s *Service) SearchQuestionsByTags(tag string) []Question {
	var questions []Question

	// finally get questions
	// SELECT * FROM questions WHERE id IN (SELECT question_id FROM question_tags WHERE tag_id IN (SELECT id FROM tags WHERE name= 'tag'));

	// this won't retun relationsips

	// Get question title and tag id
	// select questions.title, question_tags.tag_id  from questions inner join question_tags ON id=question_id;

	// combine these two tables
	// select questions.title, tags.name from questions inner join question_tags ON id=question_id inner join tags on question_tags.tag_id=tags.id;

	// and final query
	// select questions.title, tags.name from questions inner join question_tags ON id=question_id inner join tags on question_tags.tag_id=tags.id
	// where questions.id IN (SELECT question_id FROM question_tags WHERE tag_id IN (SELECT id FROM tags WHERE name= 'tag'));

	//	s.DB.Debug().Preload("Tags").Raw("SELECT * FROM questions WHERE id IN
	//(SELECT question_id FROM question_tags WHERE tag_id IN (SELECT id FROM tags WHERE name= ?))", tag).Scan(&questions)

	// this do return the correct values but gorm relationship wont work soo gonna ignore for now

	// postgres=# select questions.id, questions.title, questions.body ,tags.name  from questions inner join question_tags
	// ON id=question_id inner join tags on question_tags.tag_id=tags.id where questions.id IN
	// (SELECT question_id FROM question_tags WHERE tag_id IN (SELECT id FROM tags WHERE name= 'tag'))

	//	 id |   title    | body |    name
	//	----+------------+------+-------------
	//	  1 | test       | hehe | tag
	//	  1 | test       | hehe | programming
	//	  2 | New Test 2 | hehe | tag
	//	  2 | New Test 2 | hehe | programming
	//	(4 rows)

	s.DB.Debug().Raw("select * from questions inner join question_tags ON id=question_id inner join tags on question_tags.tag_id=tags.id where questions.id IN (SELECT question_id FROM question_tags WHERE tag_id IN (SELECT id FROM tags WHERE name= ?))", tag).Scan(&questions)

	if questions == nil {
		return []Question{}
	}

	return questions

}

// @UTILS
// CreateAnswer - write a new answer to a question
func (s *Service) CreateAnswer(answer Answer, question_id string) (Question, error) {

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
=======
>>>>>>> tmp
