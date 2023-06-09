package helper

import "regexp"

const USERNAME_REGEX = "[a-z0-9]{2,20}"

func ValidateUsername(username string) bool {
	isValid, err := regexp.MatchString(USERNAME_REGEX, username)
	if isValid && err == nil {
		return true
	}
	return false
}
