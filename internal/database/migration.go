package database

import (
	"github.com/z9fr/greensforum-backend/internal/user"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}, &user.Account{}); err != nil {
		return err
	}
	return nil
}
