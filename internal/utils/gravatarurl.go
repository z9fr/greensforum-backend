package utils

import (
	"crypto/md5"
	"fmt"
)

func GenerateGavatarUrl(email string) string {
	hash := md5.Sum([]byte(email))
	return "https://www.gravatar.com/avatar/" + fmt.Sprintf("%x", hash) + "?s=200"
}
