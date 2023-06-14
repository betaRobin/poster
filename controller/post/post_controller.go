package post

import (
	"errors"
	"net/http"

	"github.com/betarobin/poster/enum/errlist"
	"github.com/betarobin/poster/helper"
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/model/response"
	auth "github.com/betarobin/poster/service/authentication"
	"github.com/betarobin/poster/service/post"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	userId := c.GetHeader("user-id")

	if !auth.IsValidUser(userId) {
		helper.ErrorResponse(c, http.StatusForbidden, errlist.ErrInvalidCredentials)
		return
	}

	var request request.CreatePostRequest
	c.BindJSON(&request)

	err := post.CreatePost(userId, request)

	if err == nil {
		helper.Response(c, http.StatusOK, helper.BaseResponse("post creation success"))
	} else if errors.Is(err, errlist.ErrInvalidTitleLength) || errors.Is(err, errlist.ErrInvalidDescriptionLength) {
		helper.ErrorResponse(c, http.StatusBadRequest, err)
	} else {
		helper.ErrorResponse(c, http.StatusInternalServerError, errlist.ErrInternalServerError)
	}
}

func GetPostsByUser(c *gin.Context) {
	userId := c.GetHeader("user-id")

	if !auth.IsValidUser(userId) {
		helper.ErrorResponse(c, http.StatusForbidden, errlist.ErrInvalidCredentials)
		return
	}

	posts, err := post.GetPostsByUser(userId)

	if err == nil {
		c.JSON(http.StatusOK, response.Posts(posts))
	} else {
		helper.ErrorResponse(c, http.StatusInternalServerError, errlist.ErrInternalServerError)
	}
}

func EditPost(c *gin.Context) {
	userId := c.GetHeader("user-id")

	if !auth.IsValidUser(userId) {
		helper.ErrorResponse(c, http.StatusForbidden, errlist.ErrInvalidCredentials)
		return
	}

	var request request.EditPostRequest
	c.BindJSON(&request)

	err := post.EditPost(userId, request)

	if err == nil {
		c.JSON(http.StatusOK, helper.BaseResponse("edit post success"))
	} else {
		switch err {
		case errlist.ErrInvalidPostID:
			fallthrough
		case errlist.ErrInvalidTitleLength:
			fallthrough
		case errlist.ErrInvalidDescriptionLength:
			fallthrough
		case errlist.ErrNoFieldToUpdate:
			helper.ErrorResponse(c, http.StatusBadRequest, err)
		case errlist.ErrForbidden:
			helper.ErrorResponse(c, http.StatusForbidden, err)
		default:
			helper.ErrorResponse(c, http.StatusInternalServerError, errlist.ErrInternalServerError)
		}
	}
}
