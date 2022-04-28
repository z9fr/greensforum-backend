package database

import (
	"github.com/z9fr/greensforum-backend/internal/question"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&question.Question{}); err != nil {
		return err
	}
	return nil
}
