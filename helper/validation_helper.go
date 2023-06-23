package helper

import (
	"regexp"
	"strings"

	typepost "github.com/betarobin/poster/enum/type_post"
	contenthelper "github.com/betarobin/poster/helper/content"
	"github.com/betarobin/poster/model/content"
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

func validateLength(minlen int, maxlen int, text string) bool {
	trimmed := strings.TrimSpace(text)
	return len(trimmed) > minlen && len(trimmed) < maxlen+1
}

func IsValidTitle(title string) bool {
	return validateLength(0, 70, title)
}

func IsValidText(content string) bool {
	return validateLength(0, 300, content)
}

func IsValidPostType(postType string) bool {
	if len(postType) == 0 {
		return false
	}

	return Contains(typepost.GetAllTypes(), strings.ToLower(postType))
}

func IsValidContent(postType string, postContent interface{}) bool {
	parsedContent, err := contenthelper.ParseContent(postType, postContent)

	if err != nil {
		return false
	}

	switch postType {
	case typepost.Text:
		text := parsedContent.(*content.Text)
		return IsValidContentText(text)
	case typepost.Checklist:
		checkbox := parsedContent.(*content.Checklist)
		return IsValidContentChecklist(checkbox)
	case typepost.Image:
		image := parsedContent.(*content.Image)
		return IsValidContentImage(image)
	default:
		return false
	}
}

func IsValidContentText(text *content.Text) bool {
	return IsValidText(text.Text)
}

func IsValidContentImage(image *content.Image) bool {
	return IsValidText(image.Text)
}

func IsValidContentChecklist(checklist *content.Checklist) bool {
	for _, data := range checklist.Checklist {
		if !IsValidText(data.Text) {
			return false
		}
	}
	return true
}
