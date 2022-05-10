package verification

import (
	"errors"

	"github.com/google/uuid"
	"github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
	"gorm.io/gorm"
)

type EmailVerification struct {
	types.Model
	UserId  uint64 `gorm:"column:user_id"`
	Token   string `gorm:"column:token"`
	Secret  string `gorm:"secret"`
	IsValid bool   `gorm:"column:isvalid"`
}

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) RequestVerification(user user.User, token string) (error, bool) {
	vinfo := s.GetTokenInfo(token)

	if len(vinfo.Secret) > 0 {
		return errors.New("Invalid token"), false
	}

	// validate the user and kill the token
	user.IsVerified = true
	vinfo.IsValid = false
	// save them on db later
	return nil, true
}

func (s *Service) GetTokenInfo(token string) EmailVerification {
	var vinfo EmailVerification
	s.DB.Debug().Where("token = ? AND isvalid IS TRUE", token).Find(&vinfo)
	return vinfo
}

// return the token and hashed secret to the calling function
func (s *Service) GenerateVerification(user user.User) (error, string, string) {

	if user.IsVerified {
		return errors.New("already verified email"), "", ""
	}

	var emailverification EmailVerification

	emailverification.Secret = uuid.New().String()
	emailverification.Token = uuid.New().String()
	emailverification.IsValid = true
	emailverification.UserId = uint64(user.ID)

	s.DB.Save(&emailverification)
	v, err := utils.GenerateHashSecret(emailverification.Secret, emailverification.Token)

	if err != nil {
		return err, "", ""
	}

	return nil, emailverification.Token, v
}
