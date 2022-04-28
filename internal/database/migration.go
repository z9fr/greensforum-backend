package database

import (
	"github.com/z9fr/greensforum-backend/internal/question"
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}, &user.Account{}, &question.Question{}, &question.Answer{}, &question.Tag{}); err != nil {
		utils.LogFatal(err)
		return err
	}
	return nil
}
