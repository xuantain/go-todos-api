package handlers

import (
	"net/http"

	"go-todos-api/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

var userModel = new(models.User)

func (u UserHandler) Retrieve(c *gin.Context) {

	if c.Param("id") != "" {

		user, err := userModel.GetByID(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}

		c.JSON(http.StatusFound, gin.H{"message": "User founded!", "user": user})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort() // to stop current handler and ignore calling remaining handlers
}

func (u UserHandler) GetAllUsers(c *gin.Context) {
	userList := []models.User{}

	c.JSON(http.StatusOK, gin.H{"data": userList})
	c.Abort()
}

func (u UserHandler) CreateUser(c *gin.Context) {

	newUser := models.User{}

	c.JSON(http.StatusCreated, gin.H{"data": newUser})
	c.Abort()
}

func (u UserHandler) UpdateUser(c *gin.Context) {

	newUser := models.User{}

	c.JSON(http.StatusAccepted, gin.H{"data": newUser})
	c.Abort()
}

func (u UserHandler) DeleteUser(c *gin.Context) {

	if c.Param("id") != "" {
		c.JSON(http.StatusGone, gin.H{"message": "Deleted user id:" + c.Param("id")})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
