package post

import (
	"errors"
	"strings"

	"github.com/betarobin/poster/entity"
	"github.com/betarobin/poster/enum/errlist"
	"github.com/betarobin/poster/helper"
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/repository"
	auth "github.com/betarobin/poster/service/authentication"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreatePost(userId string, request request.CreatePostRequest) error {
	if !helper.IsValidTitle(request.Title) {
		return errlist.ErrInvalidTitleLength
	}

	if !helper.IsValidDescription(request.Description) {
		return errlist.ErrInvalidDescriptionLength
	}

	userUUID, err := uuid.Parse(userId)

	if err != nil {
		return errlist.ErrInternalServerError
	}

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

func EditPost(userId string, req request.EditPostRequest) error {
	if !auth.IsValidUser(userId) {
		return errlist.ErrInvalidCredentials
	} else if req.PostID == "" {
		return errlist.ErrInvalidPostID
	} else if req.Title == nil &&
		req.Description == nil {
		return errlist.ErrNoFieldToUpdate
	}

	postUUID, err := uuid.Parse(req.PostID)

	if err != nil {
		return errlist.ErrInvalidPostID
	}

	selectedPost, result := repository.GetPostById(postUUID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errlist.ErrInvalidPostID
		} else {
			return errlist.ErrInternalServerError
		}
	} else if selectedPost.UserID.String() != userId {
		return errlist.ErrForbidden
	}

	if req.Title != nil {
		if !helper.IsValidTitle(*req.Title) {
			return errlist.ErrInvalidTitleLength
		} else {
			selectedPost.Title = strings.TrimSpace(*req.Title)
		}
	}

	if req.Description != nil {
		if !helper.IsValidDescription(*req.Description) {
			return errlist.ErrInvalidDescriptionLength
		} else {
			selectedPost.Description = strings.TrimSpace(*req.Description)
		}
	}

	repository.EditPostContent(*selectedPost)

	return nil
}
