package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/z9fr/greensforum-backend/internal/user"
)

//   https://www.bacancytechnology.com/blog/golang-jwt

func GenerateJWT(user user.User) (string, int64, error) {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expireTime := time.Now().Add(time.Minute * 30).Unix()

	claims["user"] = user
	claims["exp"] = expireTime

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", 0, nil
	}

	return tokenString, expireTime, nil
}
