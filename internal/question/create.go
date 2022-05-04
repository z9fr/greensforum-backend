package question

import (
	"errors"
	"strings"

	"github.com/z9fr/greensforum-backend/internal/types"
	utils "github.com/z9fr/greensforum-backend/internal/utils"
)

// create a new question
func (s *Service) CreateNewQuestion(question Question, topwords []types.TopWord) (Question, error) {

	if s.IsTitleExist(question.Title) {
		return Question{}, errors.New(`Title already exist. please try something else.`)
	}

	// generate slug and append related topics
	question.Related = append(question.Related, topwords...)
	question.Slug = s.GenerateSlug(question.Title)

	if result := s.DB.Debug().Save(&question); result.Error != nil {
		utils.LogWarn(result.Error)
		return Question{}, result.Error
	}

	return question, nil
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

// generate a unique slug for a post
// @TODO
// do more validations and checking
func (s *Service) GenerateSlug(title string) string {
	title = utils.FirstN(title, 80)
	title = strings.ToLower(title)
	title = strings.ReplaceAll(title, " ", "-")

	return title
}
