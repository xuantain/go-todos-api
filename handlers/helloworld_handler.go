package handlers

import (
	"fmt"
	"net/http"

	_ "go-todos-api/docs"

	"github.com/gin-gonic/gin"
)

type HelloWorldHandler struct{}

// HelloWorld		godoc
// @Summary			Say "Hello World!"
// @Description		Responds with the greeting word "Hello World!".
// @Tags			hello
// @Produce			json
// @Success			200  {text}  string
// @Router			/hello-world [get]
func (u HelloWorldHandler) SayHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

// HelloWorldBean	godoc
// @Summary			Say "Hello World! to {username}"
// @Description		Responds with the greeting word "Hello World! to {username}".
// @Tags			hello
// @Produce			json
// @Success			200  {text}  string
// @Router			/hello-world-bean/path-variable/:username [get]
func (u HelloWorldHandler) SayHelloWorldTo(c *gin.Context) {
	if c.Param("username") != "" {
		username := c.Param("username")
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello %s!", username)})
		// NOTE: just use IndentedJSON for dev because it would consume more CPU.
		// 	But we already have Pretty-print plugin on Browser, so this seems redundant.
		// c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to", "user": username})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
