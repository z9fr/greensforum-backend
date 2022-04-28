package helper

import "github.com/z9fr/greensforum-backend/internal/user"

// RequestToUser
// take the user request and convert it to User Object
// this function will do validations, password hashing
func RequestToUserWithValidations(req user.CreateUserRequest) user.User {
	// not implemented

	user := user.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		UserAcc: user.Account{
			DisplayName:  req.Account.DisplayName,
			Description:  req.Account.Description,
			Name:         req.Account.Name,
			Location:     req.Account.Location,
			WebsiteURL:   req.Account.WebsiteURL,
			ProfileImage: req.Account.ProfileImage,
		},
	}

	return user
}
