package auth

import (
	"log"

	"github.com/betarobin/poster/repository"
	"github.com/gin-gonic/gin"
)

func IsUserLoggedIn(c *gin.Context) bool {
	userId := c.GetHeader("user-id")

	if len(userId) != 0 {
		log.Println("[Login] User already logged on")
		return true
	} else {
		return false
	}
}

func IsValidUser(c *gin.Context) bool {
	userId := c.GetHeader("user-id")

	if len(userId) != 0 {
		_, result := repository.FindUserById(userId)
		return result.Error == nil
	} else {
		return false
	}
}
