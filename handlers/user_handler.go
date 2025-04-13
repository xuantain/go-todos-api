package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"go-todos-api/models"
	"go-todos-api/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepo repositories.UserRepository
}

func (u *UserHandler) GetAllUsers(c *gin.Context) {

	userList, err := u.UserRepo.List(c.Request.Context(), 0, 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userList})
	c.Abort()
}

func (u *UserHandler) CreateUser(c *gin.Context) {

	var newUser models.User

	// Call BindJSON to bind the received JSON to newUser
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.UserRepo.Create(c.Request.Context(), &newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// newUser = repo.InsertNewUser(newUser)

	c.JSON(http.StatusCreated, newUser)
	c.Abort()
}

func (u *UserHandler) UpdateUser(c *gin.Context) {

	newUser := models.User{}

	// Call BindJSON to bind the received JSON to newUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.UserRepo.Update(c.Request.Context(), &newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"data": newUser})
	c.Abort()
}

func (u *UserHandler) Retrieve(c *gin.Context) {

	if c.Param("id") != "" {

		u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		}
		userId := uint(u64)

		user, err := u.UserRepo.FindByID(c.Request.Context(), userId)
		if err != nil {
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

func (u *UserHandler) DeleteUser(c *gin.Context) {

	userId := c.GetUint("id")
	if userId != 0 {
		c.JSON(http.StatusGone, gin.H{"message": fmt.Sprintf("Deleted user id:%d", userId)})
		return
	}

	if err := u.UserRepo.Delete(c.Request.Context(), userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
