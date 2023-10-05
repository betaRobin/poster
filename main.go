package main

import (
	"log"

	"github.com/betarobin/poster/controller/post"
	"github.com/betarobin/poster/controller/user"
	"github.com/betarobin/poster/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// logging setup
	log.SetFlags(0)
	log.SetOutput(new(middleware.LogWriter))

	router := gin.New()

	// middlewares
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())

	// user
	router.POST("/login", user.Login)
	router.POST("/register", user.Register)

	// post
	router.POST("/post", post.CreatePost)
	router.GET("/all", post.GetPostsByUser)
	router.PUT("/edit", post.EditPost)
	router.DELETE("/delete", post.DeletePost)

	router.Run("localhost:8080")
}
