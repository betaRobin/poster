package post

import (
	"strings"

	"github.com/betarobin/poster/enum/errlist"
	"github.com/betarobin/poster/helper"
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/repository"
)

func CreatePost(userId string, request request.CreatePostRequest) error {
	if !helper.ValidateTitle(request.Title) {
		return errlist.ErrInvalidTitleLength
	}

	if !helper.ValidateDescription(request.Description) {
		return errlist.ErrInvalidDescriptionLength
	}

	title := strings.TrimSpace(request.Title)
	description := strings.TrimSpace(request.Description)

	_, result := repository.InsertPost(userId, title, description)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
