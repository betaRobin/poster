package main

import (
	"github.com/betarobin/poster/controller/post"
	"github.com/betarobin/poster/controller/user"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// User
	router.POST("/login", user.Login)
	router.POST("/register", user.Register)

	// Post
	router.POST("/post", post.CreatePost)
	router.GET("/all", post.GetPostsByUser)
	router.PUT("/edit", post.EditPost)
	router.DELETE("/delete", post.DeletePost)

	router.Run("localhost:8080")
}
