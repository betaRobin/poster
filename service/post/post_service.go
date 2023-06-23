package post

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/betarobin/poster/entity"
	"github.com/betarobin/poster/enum/errlist"
	typepost "github.com/betarobin/poster/enum/type_post"
	"github.com/betarobin/poster/helper"
	contenthelper "github.com/betarobin/poster/helper/content"
	"github.com/betarobin/poster/model/content"
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

	if !helper.IsValidPostType(request.Type) {
		return errlist.ErrInvalidPostType
	}

	postType := strings.ToLower(request.Type)
	if !helper.IsValidContent(postType, request.Content) {
		return errlist.ErrInvalidContent
	}

	parsedContent, _ := contenthelper.ParseContent(postType, request.Content)
	contentJsonString := ``
	switch postType {
	case typepost.Text:
		typedContent := parsedContent.(*content.Text)
		contentJson, _ := json.Marshal(typedContent)
		contentJsonString = string(contentJson)
	case typepost.Image:
		typedContent := parsedContent.(*content.Image)
		contentJson, _ := json.Marshal(typedContent)
		contentJsonString = string(contentJson)
	case typepost.Checklist:
		typedContent := parsedContent.(*content.Checklist)
		contentJson, _ := json.Marshal(typedContent)
		contentJsonString = string(contentJson)
	default:
		return errlist.ErrInternalServerError
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return errlist.ErrInternalServerError
	}

	title := strings.TrimSpace(request.Title)

	_, result := repository.InsertPost(userUUID, postType, title, contentJsonString)

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
	// if !auth.IsValidUser(userId) {
	// 	return errlist.ErrInvalidCredentials
	// } else if req.Title == nil &&
	// 	req.Content == nil {
	// 	return errlist.ErrNoFieldToUpdate
	// }

	// postUUID, err := uuid.Parse(req.PostID)

	// if err != nil {
	// 	return errlist.ErrInvalidPostID
	// }

	// selectedPost, result := repository.GetPostById(postUUID)

	// if result.Error != nil {
	// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 		return errlist.ErrPostNotFound
	// 	} else {
	// 		return errlist.ErrInternalServerError
	// 	}
	// } else if selectedPost.UserID.String() != userId {
	// 	return errlist.ErrForbidden
	// }

	// if req.Title != nil {
	// 	if !helper.IsValidTitle(*req.Title) {
	// 		return errlist.ErrInvalidTitleLength
	// 	} else {
	// 		selectedPost.Title = strings.TrimSpace(*req.Title)
	// 	}
	// }

	// if req.Content != nil {
	// 	if !helper.IsValidContent(selectedPost.Type, *req.Content) {
	// 		return errlist.ErrInvalidContent
	// 	} else {
	// 		selectedPost.Content = strings.TrimSpace(*req.Content)
	// 	}
	// }

	// // post.UpdatedAt gets automatically updated by Gorm
	// result = repository.EditPostContent(*selectedPost)

	// return result.Error
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

	return result.Error
}
