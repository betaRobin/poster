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
	if !auth.IsValidUser(c) {
		helper.ErrorResponse(c, http.StatusForbidden, errlist.ErrInvalidCredentials)
		return
	}

	var request request.CreatePostRequest
	c.BindJSON(&request)

	userId := c.GetHeader("user-id")
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
	if !auth.IsValidUser(c) {
		helper.ErrorResponse(c, http.StatusForbidden, errlist.ErrInvalidCredentials)
		return
	}

	userId := c.GetHeader("user-id")
	posts, err := post.GetPostsByUser(userId)

	if err == nil {
		c.JSON(http.StatusOK, response.Posts(posts))
	} else {
		helper.ErrorResponse(c, http.StatusInternalServerError, errlist.ErrInternalServerError)
	}
}
