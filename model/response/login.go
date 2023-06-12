package response

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(userId uuid.UUID) gin.H {
	return gin.H{
		"user_id": userId.String(),
	}
}
