package helper

import (
	"fmt"

	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

// RequestToUser
// take the user request and convert it to User Object
// this function will do validations, password hashing
func RequestToUserWithValidations(req user.CreateUserRequest) (user.User, error) {

	if req.Username == "" || req.Email == "" || req.Password == "" || req.Account.Name == "" {
		return user.User{}, fmt.Errorf("Missing Fields user")
	}

	if !utils.IsEmailValid(req.Email) {
		utils.LogWarn("Invalid Email" + req.Email + " from -> " + req.Username)
		return user.User{}, fmt.Errorf("Invalid Email")
	}
	req.Account.ProfileImage = utils.GenerateGavatarUrl(req.Email)
	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return user.User{}, err
	}

	user := user.User{
		Username:     req.Username,
		Email:        req.Email,
		Password:     hashedPassword,
		UserType:     0,
		TokenVersion: 0,
		IsVerified:   false,
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
