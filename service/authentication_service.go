package service

import (
	"log"

	"github.com/betarobin/poster/repository"
	"github.com/gin-gonic/gin"
)

func IsUserLoggedIn(c *gin.Context) bool {
	userId := c.GetHeader("poster-id")

	if len(userId) != 0 {
		log.Println("[Login] User already logged on")
		return true
	} else {
		return false
	}
}

func IsValidUser(c *gin.Context) bool {
	userId := c.GetHeader("poster-id")

	if len(userId) != 0 {
		_, err := repository.FindUserByUsername(userId)
		return err == nil
	} else {
		return false
	}
}
