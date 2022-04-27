package question

import (
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
	Tags          []string `gorm:"column:tags" json:"tags"`
}

// QuestionService - interface for Question Service
type QuestionService interface {
}

// NewService - create a instance of this service and return
// a pointer to the servie
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
