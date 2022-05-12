package events

import (
	"github.com/z9fr/greensforum-backend/internal/types"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

type Event struct {
	types.Model
	Name        string `gorm:"column:name" json:"name"`
	Slug        string `gorm:"column:slug" json:"slug"`
	Description string `gorm:"description" json:"description"`
	CreatedBy   uint   `gorm:"creted_user" json:"created_user"`
}

type CollectiveService interface {
	CreateEvent(event Event) (Event, error)
	GetAllEvents() []Event
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
