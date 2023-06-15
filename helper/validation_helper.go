package helper

import (
	"regexp"
	"strings"
)

// must be a combination of lowercase English alphabet or numeric, 2-20 characters
const USERNAME_REGEX = "[a-z0-9]{2,20}"

// is 0-7 characters, does not contain a single number, does not contain a single English alphabet
const INVALID_PASSWORD_REGEX = "^(.{0,7}|[^0-9]*|[^a-zA-Z]*)$"

func IsValidUsername(username string) bool {
	isValid, err := regexp.MatchString(USERNAME_REGEX, username)
	if isValid && err == nil {
		return true
	}
	return false
}

func IsValidPassword(password string) bool {
	isInvalid, err := regexp.MatchString(INVALID_PASSWORD_REGEX, password)
	if err == nil {
		return !isInvalid
	} else {
		return false
	}
}

func ValidateLength(minlen int, maxlen int, text string) bool {
	trimmed := strings.TrimSpace(text)
	return len(trimmed) > minlen && len(trimmed) < maxlen+1
}

func IsValidTitle(title string) bool {
	return ValidateLength(0, 70, title)
}

func IsValidDescription(description string) bool {
	return ValidateLength(0, 300, description)
}
