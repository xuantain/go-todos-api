package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct{}

func (auth AuthenticationHandler) CheckBasicAuth(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	c.JSON(http.StatusOK, gin.H{"user": user, "message": "Welcome back!"})
}
