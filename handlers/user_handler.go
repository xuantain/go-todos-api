package handlers

import (
	"fmt"
	"net/http"

	"go-todos-api/models"
	repos "go-todos-api/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

var userRepo = new(repos.UserRepo)

func (u UserHandler) GetAllUsers(c *gin.Context) {
	userList := userRepo.GetAllUsers()

	c.JSON(http.StatusOK, gin.H{"data": userList})
	c.Abort()
}

func (u UserHandler) CreateUser(c *gin.Context) {

	var newUser models.User

	// Call BindJSON to bind the received JSON to newUser
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	newUser = userRepo.InsertNewUser(newUser)

	c.JSON(http.StatusCreated, newUser)
	c.Abort()
}

func (u UserHandler) UpdateUser(c *gin.Context) {

	newUser := models.User{}

	c.JSON(http.StatusAccepted, gin.H{"data": newUser})
	c.Abort()
}

func (u UserHandler) Retrieve(c *gin.Context) {

	if c.Param("id") != "" {

		userId := c.GetUint("id")
		user := userRepo.GetByID(userId)
		fmt.Println(user)

		if user != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": nil})
			c.Abort()
			return
		}

		c.JSON(http.StatusFound, gin.H{"message": "User founded!", "user": user})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort() // to stop current handler and ignore calling remaining handlers
}

func (u UserHandler) DeleteUser(c *gin.Context) {

	if c.Param("id") != "" {
		c.JSON(http.StatusGone, gin.H{"message": "Deleted user id:" + c.Param("id")})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
