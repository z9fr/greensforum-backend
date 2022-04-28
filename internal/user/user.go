package user

import "gorm.io/gorm"

type Interests struct {
	Tags []string `json:"tags"`
}

type User struct {
	ID       uint64
	Username string  `gorm:"column:username" json:"username"`
	Email    string  `gorm:"column:email" json:"email"`
	Password string  `gorm:"column:password" json:"password"`
	UserType string  `json:"user_type"`
	UserAcc  Account `json:"user_acc" gorm:"foreignKey:user_id;id"`
}

type Account struct {
	AccountID    uint64 `gorm:"primaryKey" json:"account_id"`
	UserID       int    `json:"user_id"`
	Location     string `json:"location"`
	WebsiteURL   string `json:"website_url"`
	Link         string `json:"link"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Description  string `json:"description"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	IsEmployee   bool   `json:"is_employee"`
	Reputation   int    `json:"reputation"`
}

type Service struct {
	DB *gorm.DB
}

type UserService interface {
	CreateUser(user User) (User, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) CreateUser(user User) (User, error) {
	// not implemented
	return user, nil
}
