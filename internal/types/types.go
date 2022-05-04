package types

import "time"

// Model definition same as gorm.Model, but including column and json tags
type Model struct {
	ID        uint       `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details"`
}

type AuthRequest struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type TopWord struct {
	Model
	Word  string `json:"word"`
	Count int    `json:"count"`
}
