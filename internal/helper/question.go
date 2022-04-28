package helper

import (
	"fmt"

	"github.com/z9fr/greensforum-backend/internal/question"
	"github.com/z9fr/greensforum-backend/internal/user"
)

func RequestPostWithValidation(questionreq question.QuestionCreateRequest, user user.User) (question.Question, error) {

	if questionreq.Title == "" || questionreq.Body == "" {
		return question.Question{}, fmt.Errorf("Missing Fields")
	}

	return question.Question{
		Title:     questionreq.Title,
		Body:      questionreq.Body,
		Tags:      questionreq.Tags,
		CreatedBy: user.ID,
	}, nil
}
