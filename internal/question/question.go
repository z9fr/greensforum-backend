package question

import (
	"errors"

	models "github.com/z9fr/greensforum-backend/internal/types"
	"gorm.io/gorm"
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

type QuestionCreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Tags  []Tag  `json:"tags"`
}

// QuestionService - interface for Question Service
type QuestionService interface {
	CreateNewQuestion(question Question) (Question, error)
	GetAllPosts() []Question
	SearchQuestionsByTags(tag string) []Question
	SearchPosts(q string) []Question
}

// NewService - create a instance of this service and return
// a pointer to the servie
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) CreateNewQuestion(question Question) (Question, error) {

	if s.IsTitleExist(question.Title) {
		return Question{}, errors.New(`Title already exist. please try something else.`)
	}

	if result := s.DB.Debug().Save(&question); result.Error != nil {
		return Question{}, result.Error
	}

	return question, nil
}

func (s *Service) GetAllPosts() []Question {
	var questions []Question
	s.DB.Debug().Preload("Tags").Find(&questions)

	return questions
}

func (s *Service) SearchPosts(q string) []Question {
	var questions []Question
	s.DB.Debug().Where("title LIKE ?", "%"+q+"%").Preload("Tags").Find(&questions)
	return questions
}

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
