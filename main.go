package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "database/sql"
    // _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "robinwantah"
    password = "RobinWantah369"
    dbname   = "postgres"
)

var credentials = map[string]string{
	"UserOne" : "PwOne",
	"UserTwo" : "PwTwo",
}

type credential struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type note struct {
	Title		string	`json:"title"`
	Description	string	`json:"description"`
	Username	string	`json:"username"`
}

var notes = [] note{
	{Title: "Title One", Description: "Title description", Username: "UserOne"},
	{Title: "Title Two", Description: "Title description", Username: "UserOne"},
	{Title: "Title One", Description: "Title description", Username: "UserTwo"},
}

func postLogin(c *gin.Context) {
	var loginCredentials credential

	c.BindJSON(&loginCredentials)

	password, exists := credentials[loginCredentials.Username]

	if (!exists) || (password != loginCredentials.Password){
		fmt.Println("Username or Password incorrect")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"message": "Username or Password incorrect",
		})
	} else {
		fmt.Println("Login success")
		c.JSON(http.StatusOK, gin.H{
			"status": "200",
			"message": "Success",
		})
	}
}

func postRegister(c *gin.Context) {
	var signUpCredentials credential

	c.BindJSON(&signUpCredentials)

	credentials[signUpCredentials.Username] = signUpCredentials.Password

	fmt.Println("Registration success")
	c.JSON(http.StatusOK, gin.H{
		"status": "200",
		"message": "Success",
	})
}

func postNote(c *gin.Context) {
	var newNote note

	c.BindJSON(&newNote)

	notes = append(notes, newNote)
	fmt.Println("Note created")
	c.JSON(http.StatusOK, gin.H{
		"status": "200",
		"message": "Success",
	})
}

type getNotesRequest struct {
	Username	string	`json:"username"`
}

func getNotesByUser(c *gin.Context) {
	var request getNotesRequest

	c.BindJSON(&request)

	var userNotes = []note{}

	for _, note := range notes {
		if note.Username == request.Username {
			userNotes = append(userNotes, note)
		}
	}

	c.JSON(http.StatusOK, userNotes)
}

func main() {
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlconn)
	// defer db.Close()

	router := gin.Default()
	router.POST("/login", postLogin)
	router.POST("/register", postRegister)
	router.POST("/post", postNote)
	router.GET("/all", getNotesByUser)
	router.Run("localhost:8080")
}