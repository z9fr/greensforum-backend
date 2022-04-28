package helper

import (
	"fmt"
	"strconv"

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

func RequstAnswerValidation(answerreq question.AnswerRequest, user user.User, questionid_string string) (question.Answer, error) {

	if answerreq.Title == "" || answerreq.Body == "" {
		return question.Answer{}, nil
	}

	question_id, err := strconv.Atoi(questionid_string)

	if err != nil {
		return question.Answer{}, err
	}

	return question.Answer{
		Title:      answerreq.Title,
		Body:       answerreq.Body,
		QuestionID: int64(question_id),
	}, nil

}
