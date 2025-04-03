package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWorldHandler struct{}

func (u HelloWorldHandler) SayHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func (u HelloWorldHandler) SayHelloWorldTo(c *gin.Context) {
	if c.Param("username") != "" {
		username := c.Param("username")
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to", "user": username})
		// NOTE: just use IndentedJSON for dev because it would consume more CPU.
		// 	But we already have Pretty-print plugin on Browser, so this seems redundant.
		// c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to", "user": username})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
