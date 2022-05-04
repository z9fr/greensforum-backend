package question

import (
	models "github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/user"
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
	Title         string      `gorm:"column:title" json:"title"`
	Body          string      `gorm:"column:body" json:"body"`
	IsAnswered    bool        `gorm:"column:is_answered default:false" json:"is_answered"`
	ViewCount     int         `gorm:"column:view_count default:0" json:"view_count"`
	DownVoteCount int         `gorm:"column:down_vote_count default:0" json:"down_vote_count"`
	UpVoteCount   int         `grom:"column:up_vote_count default:0" json:"up_vote_count"`
	AnswerCount   int         `gorm:"column:answer_count default:0" json:"answer_count"`
	Score         int         `gorm:"column:score default:0" json:"score"`
	CreatedBy     uint64      `gorm:"column:created_by" json:"created_by" `
	Slug          string      `gorm:"column:slug" json:"slug"`
	Answers       []Answer    `gorm:"foreignKey:question_id;id" json:"answers"`
	Tags          []Tag       `gorm:"many2many:question_tags" json:"tags"`
	UpvotedUsers  []UpVotedBy `gorm:"many2many:question_id;id" json:"upvotedUsers"`
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
	GetAllPosts() []Question
	SearchQuestionsByTags(tag string) []Question
	SearchPosts(q string) []Question
	CreateAnswer(answer Answer, question_id string) (Question, error)

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
