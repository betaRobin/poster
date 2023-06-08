package main

import (
	"fmt"
	"net/http"

	controller "github.com/betarobin/poster/controller"
	entity "github.com/betarobin/poster/entity"
	request "github.com/betarobin/poster/model/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var users = []entity.User{
	{Username: "UserOne", Password: "PwOne"},
	{Username: "UserTwo", Password: "PwTwo"},
}

var posts = []entity.Post{
	{Id: uuid.New(), Title: "Title One", Description: "Title description", Username: "UserOne"},
	{Id: uuid.New(), Title: "Title Two", Description: "Title description", Username: "UserOne"},
	{Id: uuid.New(), Title: "Title One", Description: "Title description", Username: "UserTwo"},
}

func register(c *gin.Context) {
	var request request.Register
	c.BindJSON(&request)

	for _, user := range users {
		if user.Username == request.Username {
			fmt.Printf("Registration failed, username \"%s\" already taken\n", user.Username)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "400",
				"message": "Username already taken",
			})
			return
		}
	}

	users = append(users, entity.User{Username: request.Username, Password: request.Password})

	fmt.Println("Registration success")
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
	})
}

func createPost(c *gin.Context) {
	var newPost entity.Post
	c.BindJSON(&newPost)

	newPost.Id = uuid.New()

	posts = append(posts, newPost)
	fmt.Println("Post created")
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
	})
}

func getPostsByUser(c *gin.Context) {
	var request request.GetPostsRequest
	c.BindJSON(&request)

	var userPosts = []entity.Post{}

	for _, post := range posts {
		if post.Username == request.Username {
			userPosts = append(userPosts, post)
		}
	}

	c.JSON(http.StatusOK, userPosts)
}

func editPost(c *gin.Context) {
	var request request.EditPostRequest
	c.BindJSON(&request)

	if request.Title == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Post Title cannot be empty",
		})
		return
	}

	if request.Description == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Post Description cannot be empty",
		})
		return
	}

	for i, post := range posts {
		if post.Id == request.PostId && post.Username == request.Username {
			posts[i].Title = request.Title
			posts[i].Description = request.Description

			c.JSON(http.StatusOK, posts[i])
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"status":  "404",
		"message": "Post not found",
	})
}

func removePost(ps []entity.Post, index int) []entity.Post {
	ret := make([]entity.Post, 0)
	ret = append(ret, ps[:index]...)
	return append(ret, ps[index+1:]...)
}

func deletePost(c *gin.Context) {
	var request request.DeletePostRequest
	c.BindJSON(&request)

	for idx, post := range posts {
		if post.Id == request.PostId {
			if post.Username == request.Username {
				posts = removePost(posts, idx)
				c.JSON(http.StatusOK, gin.H{
					"status":  "200",
					"message": "Success",
				})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"status":  "403",
					"message": "Forbidden",
				})
				return
			}
		}
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"status":  "404",
		"message": "Post not found",
	})
}

func main() {
	router := gin.Default()

	// User
	router.POST("/login", controller.Login)
	router.POST("/register", register)

	// Post
	router.POST("/post", createPost)
	router.GET("/all", getPostsByUser)
	router.PUT("/edit", editPost)
	router.DELETE("/delete", deletePost)

	router.Run("localhost:8080")
}
