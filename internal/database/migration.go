package database

import (
	"github.com/z9fr/greensforum-backend/internal/collective"
	"github.com/z9fr/greensforum-backend/internal/question"
	"github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}, &user.Account{},
		&question.Question{}, &question.Answer{}, &question.Tag{},
		&question.UpVotedBy{}, &types.TopWord{}, &collective.Post{},
		&collective.Comments{}, &collective.Collective{}); err != nil {
		utils.LogFatal(err)
		return err
	}
	return nil
}
