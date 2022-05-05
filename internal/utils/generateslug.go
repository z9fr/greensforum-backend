package utils

import "strings"

// generate a unique slug for a post
// @TODO
// do more validations and checking
func GenerateSlug(title string) string {
	title = FirstN(title, 80)
	title = strings.ToLower(title)
	title = strings.ReplaceAll(title, " ", "-")

	return title
}
