package helper

import (
	"regexp"
	"strings"
)

const USERNAME_REGEX = "[a-z0-9]{2,20}"

func ValidateUsername(username string) bool {
	isValid, err := regexp.MatchString(USERNAME_REGEX, username)
	if isValid && err == nil {
		return true
	}
	return false
}

func ValidateLength(minlen int, maxlen int, text string) bool {
	trimmed := strings.TrimSpace(text)
	return len(trimmed) > minlen && len(trimmed) < maxlen+1
}

func ValidateTitle(title string) bool {
	return ValidateLength(0, 70, title)
}

func ValidateDescription(description string) bool {
	return ValidateLength(0, 300, description)
}
