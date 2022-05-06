package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/z9fr/greensforum-backend/internal/user"
)

func SendRefreshToken(user user.User) (string, int64, error) {
	var mySigningKey = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expireTime := time.Now().Add(time.Hour * 168).Unix() // 7 days

	claims["user_id"] = user.ID
	claims["user_email"] = user.Email
	claims["token_version"] = user.TokenVersion

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", 0, nil
	}

	return tokenString, expireTime, nil
}

func ValidateRefreshToken(tokenString string) (interface{}, interface{}, interface{}, error) {

	var hmacSampleSecret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != nil {
		return 0, 0, 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		user_id := claims["user_id"]
		user_email := claims["user_email"]
		token_version := claims["token_version"]

		return user_id, user_email, token_version, nil
	} else {
		return 0, 0, 0, err
	}
}
