package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
// host     = "localhost"
// port     = 5432
// user     = "robinwantah"
// password = "RobinWantah369"
// dbname   = "postgres"
)

var credentials = map[string]string{
	"UserOne": "PwOne",
	"UserTwo": "PwTwo",
}

type credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type post struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Username    string    `json:"username"`
}

var posts = []post{
	{Id: uuid.New(), Title: "Title One", Description: "Title description", Username: "UserOne"},
	{Id: uuid.New(), Title: "Title Two", Description: "Title description", Username: "UserOne"},
	{Id: uuid.New(), Title: "Title One", Description: "Title description", Username: "UserTwo"},
}

func login(c *gin.Context) {
	var loginCredentials credential
	c.BindJSON(&loginCredentials)

	password, exists := credentials[loginCredentials.Username]

	if (!exists) || (password != loginCredentials.Password) {
		fmt.Println("Username or Password incorrect")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Username or Password incorrect",
		})
	} else {
		fmt.Println("Login success")
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "Success",
		})
	}
}

func register(c *gin.Context) {
	var signUpCredentials credential
	c.BindJSON(&signUpCredentials)

	credentials[signUpCredentials.Username] = signUpCredentials.Password

	fmt.Println("Registration success")
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
	})
}

func createPost(c *gin.Context) {
	var newPost post
	c.BindJSON(&newPost)

	newPost.Id = uuid.New()

	posts = append(posts, newPost)
	fmt.Println("Post created")
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
	})
}

type getPostsRequest struct {
	Username string `json:"username"`
}

func getPostsByUser(c *gin.Context) {
	var request getPostsRequest
	c.BindJSON(&request)

	var userPosts = []post{}

	for _, post := range posts {
		if post.Username == request.Username {
			userPosts = append(userPosts, post)
		}
	}

	c.JSON(http.StatusOK, userPosts)
}

type updatePostRequest struct {
	Username    string    `json:"username"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PostId      uuid.UUID `json:"post_id"`
}

func editPost(c *gin.Context) {
	var request updatePostRequest
	c.BindJSON(&request)

	if request.Title == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Post Title cannot be empty",
		})
	}

	if request.Description == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Post Description cannot be empty",
		})
	}

	for _, post := range posts {
		if post.Id == request.PostId {
			c.JSON(http.StatusOK, post)
		}
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"status":  "404",
		"message": "Post not found",
	})
}

type deletePostRequest struct {
	Username string    `json:"username"`
	PostId   uuid.UUID `json:"post_id"`
}

func removePost(ps []post, index int) []post {
	ret := make([]post, 0)
	ret = append(ret, ps[:index]...)
	return append(ret, ps[index+1:]...)
}

func deletePost(c *gin.Context) {
	var request deletePostRequest
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
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlconn)
	// defer db.Close()

	router := gin.Default()

	// POST
	router.POST("/login", login)
	router.POST("/register", register)
	router.POST("/post", createPost)

	// GET
	router.GET("/all", getPostsByUser)

	// PUT
	router.PUT("/edit", editPost)

	// Delete
	router.DELETE("/delete", deletePost)

	router.Run("localhost:8080")
}
