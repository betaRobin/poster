package helper

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatus int, response gin.H) {
	c.JSON(httpStatus, response)
}

func ErrorResponse(c *gin.Context, httpStatus int, err error) {
	c.JSON(httpStatus, gin.H{
		"error": err.Error(),
	})
}

func BaseResponse(message string) gin.H {
	return gin.H{
		"message": message,
	}
}
