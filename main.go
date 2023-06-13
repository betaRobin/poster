package main

import (
	"github.com/betarobin/poster/controller/post"
	"github.com/betarobin/poster/controller/user"
	entity "github.com/betarobin/poster/entity"
	request "github.com/betarobin/poster/model/request"
	"github.com/gin-gonic/gin"
)

func editPost(c *gin.Context) {
	var request request.EditPostRequest
	c.BindJSON(&request)

	// if request.Title == "" {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"status":  "400",
	// 		"message": "Post Title cannot be empty",
	// 	})
	// 	return
	// }

	// if request.Description == "" {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"status":  "400",
	// 		"message": "Post Description cannot be empty",
	// 	})
	// 	return
	// }

	// for i, post := range posts {
	// 	if post.Id == request.PostId && post.Username == request.Username {
	// 		posts[i].Title = request.Title
	// 		posts[i].Description = request.Description

	// 		c.JSON(http.StatusOK, posts[i])
	// 		return
	// 	}
	// }

	// c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
	// 	"status":  "404",
	// 	"message": "Post not found",
	// })
}

func removePost(ps []entity.Post, index int) []entity.Post {
	ret := make([]entity.Post, 0)
	ret = append(ret, ps[:index]...)
	return append(ret, ps[index+1:]...)
}

func deletePost(c *gin.Context) {
	var request request.DeletePostRequest
	c.BindJSON(&request)

	// for idx, post := range posts {
	// 	if post.Id == request.PostId {
	// 		if post.Username == request.Username {
	// 			posts = removePost(posts, idx)
	// 			c.JSON(http.StatusOK, gin.H{
	// 				"status":  "200",
	// 				"message": "Success",
	// 			})
	// 			return
	// 		} else {
	// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
	// 				"status":  "403",
	// 				"message": "Forbidden",
	// 			})
	// 			return
	// 		}
	// 	}
	// }

	// c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
	// 	"status":  "404",
	// 	"message": "Post not found",
	// })
}

func main() {
	router := gin.Default()

	// User
	router.POST("/login", user.Login)
	router.POST("/register", user.Register)

	// Post
	router.POST("/post", post.CreatePost)
	router.GET("/all", post.GetPostsByUser)
	// router.PUT("/edit", editPost)
	// router.DELETE("/delete", deletePost)

	router.Run("localhost:8080")
}
