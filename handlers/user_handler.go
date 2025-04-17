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

// Users			godoc
// @Summary			Get all users
// @Description		Responds with the list of users
// @Tags			users
// @Produce			json
// @Success			200  {array}  []models.User
// @Router			/api/users [get]
func (u *UserHandler) GetAllUsers(c *gin.Context) {

	userList, err := u.UserRepo.List(c.Request.Context(), 0, 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userList)
	c.Abort()
}

// Users			godoc
// @Summary			Create a new user
// @Description		Responds with the new user
// @Tags			users
// @Produce			json
// @Success			201  {object}  models.User
// @Router			/api/users [post]
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

// Users			godoc
// @Summary			Update an existed user
// @Description		Responds with the updated user
// @Tags			users
// @Produce			json
// @Success			202  {object}  models.User
// @Router			/api/users/:id [put]
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

	c.JSON(http.StatusAccepted, newUser)
	c.Abort()
}

// Users			godoc
// @Summary			Retreive an user
// @Description		Responds with the user
// @Tags			users
// @Produce			json
// @Success			302  {object}  models.User
// @Router			/api/users/:id [get]
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

		// User founded!
		c.JSON(http.StatusFound, user)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort() // to stop current handler and ignore calling remaining handlers
}

// Users			godoc
// @Summary			Delete an user
// @Description		Responds with the message
// @Tags			users
// @Produce			json
// @Success			410  {json}  { 'message': 'Deleted user id:%d' }
// @Router			/api/users/:id [delete]
func (u *UserHandler) DeleteUser(c *gin.Context) {

	if c.Param("id") != "" {

		u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		}
		userId := uint(u64)

		if err := u.UserRepo.Delete(c.Request.Context(), userId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(http.StatusGone, gin.H{"message": fmt.Sprintf("Deleted user id:%d", userId)})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
