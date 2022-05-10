package user

import (
	"github.com/z9fr/greensforum-backend/internal/types"
	"gorm.io/gorm"
)

type Interests struct {
	Tags []string `json:"tags"`
}

/*
User Types
==========

0 -> user
1 -> mod
2 -> admin

*/

type User struct {
	ID           uint64
	Username     string          `gorm:"column:username" json:"username"`
	Email        string          `gorm:"column:email" json:"email"`
	Password     string          `gorm:"column:password" json:"-"`
	UserType     int             `gorm:"column:user_type" json:"user_type"`
	UserAcc      Account         `json:"account" gorm:"foreignKey:user_id;id"`
	Nofications  []Nofication    `gorm:"many2many:user_nofication" json:"-"`
	Interests    []types.TopWord `gorm:"many2many:user_interests" json:"interests"`
	TokenVersion int             `gorm:"column:tokenversion" json:"tokenversion"`
	IsVerified   bool            `gorm:"column:verified" json:"verified"`
}

type Account struct {
	AccountID    uint64 `gorm:"primaryKey" json:"account_id"`
	UserID       int    `gorm:"column:user_id" json:"user_id"`
	Location     string `gorm:"column:location" json:"location"`
	WebsiteURL   string `gorm:"column:website_url" json:"website_url"`
	ProfileImage string `gorm:"column:profile_image" json:"profile_image"`
	DisplayName  string `gorm:"column:display_name" json:"display_name"`
	Description  string `gorm:"column:description" json:"description"`
	Name         string `gorm:"column:name" json:"name"`
	IsEmployee   bool   `gorm:"column:is_employee default:false" json:"is_employee"`
	Reputation   int    `gorm:"column:reputation default:0" json:"reputation"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Account  struct {
		DisplayName  string `json:"display_name"`
		Description  string `json:"description"`
		Name         string `json:"name"`
		Location     string `json:"location"`
		WebsiteURL   string `json:"website_url"`
		ProfileImage string `json:"profile_image"`
	} `json:"account"`
}

type Nofication struct {
	types.Model
	Message string `gorm:"column:message" json:"message"`
	Read    bool   `gorm:"read" json:"read"`
}

type Service struct {
	DB *gorm.DB
}

type UserService interface {
	CreateUser(user User) (User, error)
	FetchallUsers() []User
	GetUserByEmail(email string) (User, error)
	GetUserbyID(id uint, email string) (User, error)

	// roles
	IsAdmin(user User) bool
	IsHighPriv(user User)

	// nofications
	SendNofications(user User, nofication Nofication)
	GetNofications(user User) []Nofication

	// interests
	AppendUserInterests(user User, interests types.TopWord) bool
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
