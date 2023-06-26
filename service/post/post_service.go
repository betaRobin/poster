package post

import (
	"errors"
	"strings"

	"github.com/betarobin/poster/entity"
	"github.com/betarobin/poster/enum/errlist"
	"github.com/betarobin/poster/helper"
	contenthelper "github.com/betarobin/poster/helper/content"
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/repository"
	auth "github.com/betarobin/poster/service/authentication"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreatePost(userId string, req request.CreatePostRequest) error {
	if !helper.IsValidTitle(req.Title) {
		return errlist.ErrInvalidTitleLength
	}

	if !helper.IsValidPostType(req.Type) {
		return errlist.ErrInvalidPostType
	}

	if !helper.IsValidContent(req.Type, req.Content) {
		return errlist.ErrInvalidContent
	}

	contentJsonString, err := contenthelper.ContentToJsonString(req.Type, req.Content)
	if err != nil {
		return errlist.ErrInvalidContent
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return errlist.ErrInternalServerError
	}

	title := strings.TrimSpace(req.Title)

	_, result := repository.InsertPost(userUUID, req.Type, title, contentJsonString)

	return result.Error
}

func GetPostsByUser(userId string) (*[]entity.Post, error) {
	userUUID, err := uuid.Parse(userId)

	if err != nil {
		return nil, errlist.ErrInternalServerError
	}

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
	} else if !helper.IsValidPostType(req.Type) {
		return errlist.ErrInvalidPostType
	} else if !helper.IsValidTitle(req.Title) {
		return errlist.ErrInvalidTitleLength
	} else if !helper.IsValidContent(req.Type, req.Content) {
		return errlist.ErrInvalidContent
	}

	postUUID, err := uuid.Parse(req.PostID)
	if err != nil {
		return errlist.ErrInvalidPostID
	}

	selectedPost, result := repository.GetPostById(postUUID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errlist.ErrPostNotFound
		} else {
			return errlist.ErrInternalServerError
		}
	} else if selectedPost.UserID.String() != userId {
		return errlist.ErrForbidden
	} else if selectedPost.Type != req.Type { // currently does NOT support post type edits
		return errlist.ErrInvalidPostType
	}

	contentJsonString, err := contenthelper.ContentToJsonString(req.Type, req.Content)
	if err != nil {
		return errlist.ErrInvalidContent
	}

	selectedPost.Content = contentJsonString
	selectedPost.Title = req.Title

	// post.UpdatedAt gets automatically updated by Gorm
	result = repository.EditPostContent(*selectedPost)
	if result.Error != nil {
		return errlist.ErrInternalServerError
	}

	return nil
}

func DeletePost(userId string, req request.DeletePostRequest) error {
	if !auth.IsValidUser(userId) {
		return errlist.ErrInvalidCredentials
	}

	postUUID, err := uuid.Parse(req.PostID)

	if err != nil {
		return errlist.ErrInvalidPostID
	}

	selectedPost, result := repository.GetPostById(postUUID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errlist.ErrPostNotFound
		} else {
			return errlist.ErrInternalServerError
		}
	} else if selectedPost.UserID.String() != userId {
		return errlist.ErrForbidden
	}

	// post.DeletedAt gets automatically updated by Gorm
	result = repository.DeletePostById(postUUID)
	if result.Error != nil {
		return errlist.ErrInternalServerError
	}

	return nil
}
