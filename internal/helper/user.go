package helper

import (
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

// RequestToUser
// take the user request and convert it to User Object
// this function will do validations, password hashing
func RequestToUserWithValidations(req user.CreateUserRequest) (user.User, error) {
	// not implemented

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return user.User{}, err
	}

	user := user.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		UserAcc: user.Account{
			DisplayName:  req.Account.DisplayName,
			Description:  req.Account.Description,
			Name:         req.Account.Name,
			Location:     req.Account.Location,
			WebsiteURL:   req.Account.WebsiteURL,
			ProfileImage: req.Account.ProfileImage,
		},
	}

	return user, nil
}
