package feed

import (
	"github.com/z9fr/greensforum-backend/internal/question"
	"github.com/z9fr/greensforum-backend/internal/user"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

type UserFeedService interface {
	GetUserInterestedQuestionsEngagement(u user.User) []question.Question
}

// NewService - create a instance of this service and return
// a pointer to the servie
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
