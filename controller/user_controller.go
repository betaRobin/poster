package controller

import (
	"errors"
	"net/http"

	"github.com/betarobin/poster/enum/errlist"
	"github.com/betarobin/poster/helper"
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if service.IsUserLoggedIn(c) {
		helper.Response(c, http.StatusBadRequest, helper.BaseResponse("User currently logged on"))
		return
	}

	var request request.Login
	c.BindJSON(&request)

	userId, err := service.Login(request)

	if err == nil {
		helper.Response(c, http.StatusOK, gin.H{
			"user-id": userId.String(),
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

	err := service.Register(request)

	if err == nil {
		helper.Response(c, http.StatusOK, helper.BaseResponse("User registration success"))
	} else if errors.Is(err, errlist.ErrInvalidUserName) || errors.Is(err, errlist.ErrUsernameTaken) {
		helper.ErrorResponse(c, http.StatusBadRequest, err)
	} else {
		helper.ErrorResponse(c, http.StatusInternalServerError, errlist.ErrInternalServerError)
	}
}
