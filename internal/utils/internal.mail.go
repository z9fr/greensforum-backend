package utils

import "regexp"

func IsInternalMail(email string) bool {
	match, _ := regexp.MatchString("^[A-Za-z0-9._%+-]+@nsbm.ac.lk$", email)
	return match
}
