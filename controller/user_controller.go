package controller

import (
	"fmt"
	"net/http"

	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request request.Login
	c.BindJSON(&request)

	if service.Login(request) {
		fmt.Println("Login success")
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "Success",
		})
	} else {
		fmt.Println("Login failed")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Invalid username/password",
		})
	}
}

func Register(c *gin.Context) {
	var request request.Register
	c.BindJSON(&request)
}
