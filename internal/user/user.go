package user

type User struct {
	AccountID    int    `json:"account_id"`
	IsEmployee   bool   `json:"is_employee"`
	Reputation   int    `json:"reputation"`
	UserType     string `json:"user_type"`
	UserID       int    `json:"user_id"`
	Location     string `json:"location"`
	WebsiteURL   string `json:"website_url"`
	Link         string `json:"link"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Role         string `json:"role"`
	Description  string `json:"description"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
}

type Interests struct {
	Tags []string `json:"tags"`
}
