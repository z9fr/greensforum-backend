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
	Tags        pq.StringArray `gorm:"type:varchar(64)[]" json:"tags" swaggertype:"string"`
	CreatedBy   uint           `gorm:"creted_user" json:"created_user"`
	Admins      []user.User    `gorm:"many2many:collective_admins;"`
	Post        []Post         `gorm:"many2many:collective_posts;"`
}

type Post struct {
	types.Model
	Title         string         `gorm:"column:title" json:"title"`
	Slug          string         `gorm:"column:slug" json:"slug"`
	Body          string         `gorm:"column:body" json:"body"`
	IsAccepted    bool           `gorm:"column:is_accepted" json:"is_accepted"`
	DownVoteCount int            `gorm:"column:down_vote_count default:0" json:"down_vote_count"`
	UpVoteCount   int            `grom:"column:up_vote_count default:0" json:"up_vote_count"`
	CreatedBy     uint           `gorm:"creted_user" json:"created_user"`
	Tags          pq.StringArray `gorm:"type:varchar(64)[]" json:"tags" swaggertype:"array"`
	Comments      []Comments     `gorm:"many2many:post_comments;" json:"comments"`
}

type Comments struct {
	types.Model
	CreatedBy uint   `gorm:"creted_user" json:"created_user"`
	Content   string `gorm:"column:body" json:"content"`
}

type CollectiveService interface {
	// create
	CreateNewCollective(collective Collective) (Collective, error)

	// fetch
	IsUniqueSlug(slug string) bool
	GetAllCollectives() []*Collective
	GetCollectiveBySlug(slug string) *Collective

	// post
	CreatePostinCollective(post Post, u user.User, collective_slug string) (Collective, []user.User, user.Nofication, error, bool)

	// privilages
	IsCollectiveAdmin(collective *Collective, user user.User) bool
	IsCollectiveMember(collective *Collective, user user.User) bool
}

// NewService - create a instance of this service and return
// a pointer to the servie
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
