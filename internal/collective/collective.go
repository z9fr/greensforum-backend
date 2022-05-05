package collective

import (
	"github.com/lib/pq"
	"github.com/z9fr/greensforum-backend/internal/types"
	"github.com/z9fr/greensforum-backend/internal/user"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

// https://stackoverflow.com/a/62049071/17126147
type Collective struct {
	types.Model
	Name        string         `gorm:"column:name" json:"name"`
	Slug        string         `gorm:"column:slug" json:"slug"`
	Logo        string         `gorm:"column:logo_url" json:"logo_url"`
	Description string         `gorm:"description" json:"description"`
	Website     string         `gorm:"website" json:"website"`
	Github      string         `gorm:"gh" json:"gh"`
	Twitter     string         `gorm:"twitter" json:"twitter"`
	Facebook    string         `gorm:"fb" json:"fb"`
	Members     []user.User    `gorm:"many2many:user_collective;"`
	Tags        pq.StringArray `gorm:"type:varchar(64)[]" json:"tags"`
	CreatedBy   uint           `gorm:"creted_user" json:"created_user"`
	Admins      []user.User    `gorm:"many2many:collective_admins;"`
}

type CollectiveService interface {
	// create
	CreateNewCollective(collective Collective) (Collective, error)

	// fetch
	GetAllCollectives() []Collective
}

// NewService - create a instance of this service and return
// a pointer to the servie
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
