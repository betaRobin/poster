package post

import (
	"strings"

	"github.com/betarobin/poster/entity"
	"github.com/betarobin/poster/enum/errlist"
	"github.com/betarobin/poster/helper"
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/repository"
	"github.com/google/uuid"
)

func CreatePost(userId string, request request.CreatePostRequest) error {
	if !helper.ValidateTitle(request.Title) {
		return errlist.ErrInvalidTitleLength
	}

	if !helper.ValidateDescription(request.Description) {
		return errlist.ErrInvalidDescriptionLength
	}

	userUUID := uuid.MustParse(userId)
	title := strings.TrimSpace(request.Title)
	description := strings.TrimSpace(request.Description)

	_, result := repository.InsertPost(userUUID, title, description)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func GetPostsByUser(userId string) (*[]entity.Post, error) {
	userUUID := uuid.MustParse(userId)
	posts, result := repository.GetPostsByUserId(userUUID)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return posts, nil
	}
}

func EditPost(title string, description string) error {
	return nil
}
