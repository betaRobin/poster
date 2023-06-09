package controller

import (
	"github.com/betarobin/poster/helper"
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request request.Login
	c.BindJSON(&request)

	status, message := service.Login(request)

	helper.Response(c, status, message)
}

func Register(c *gin.Context) {
	var request request.Register
	c.BindJSON(&request)

	status, message := service.Register(request)

	helper.Response(c, status, message)
}
