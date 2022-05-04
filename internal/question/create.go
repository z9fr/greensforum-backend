package question

import (
	"errors"

	utils "github.com/z9fr/greensforum-backend/internal/utils"
)

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
