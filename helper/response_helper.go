package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{
		"status":  strconv.Itoa(httpStatus),
		"message": message,
	})
}
