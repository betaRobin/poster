package user

import (
	"errors"
	"net/http"

	"github.com/betarobin/poster/enum/errlist"
	"github.com/betarobin/poster/helper"
	"github.com/betarobin/poster/model/request"
	auth "github.com/betarobin/poster/service/authentication"
	"github.com/betarobin/poster/service/user"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	userId := c.GetHeader("user-id")

	if auth.IsUserLoggedIn(userId) {
		helper.Response(c, http.StatusBadRequest, helper.BaseResponse("user currently logged on"))
		return
	}

	var request request.Login
	c.BindJSON(&request)

	userUUID, err := user.Login(request)

	if err == nil {
		helper.Response(c, http.StatusOK, gin.H{
			"user_id": userUUID.String(),
		})
	} else if errors.Is(err, errlist.ErrInvalidLogin) {
		helper.ErrorResponse(c, http.StatusBadRequest, err)
	} else {
		helper.ErrorResponse(c, http.StatusInternalServerError, errlist.ErrInternalServerError)
	}
}

func Register(c *gin.Context) {
	var request request.Register
	c.BindJSON(&request)

	err := user.Register(request)

	if err == nil {
		helper.Response(c, http.StatusOK, helper.BaseResponse("user registration success"))
	} else {
		switch err {
		case errlist.ErrInvalidUsername:
			fallthrough
		case errlist.ErrInvalidPassword:
			fallthrough
		case errlist.ErrUsernameTaken:
			helper.ErrorResponse(c, http.StatusBadRequest, err)
		default:
			helper.ErrorResponse(c, http.StatusInternalServerError, errlist.ErrInternalServerError)
		}
	}
}
