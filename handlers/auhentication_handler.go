package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct{}

func (auth AuthenticationHandler) GetBasicAuth(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"token": "returnedBasicAuth"})
	// return
	// if c.Param("token") != "" {
	// 	token := c.GetString("username")

	// 	if token == "todo" {
	// 		returnedBasicAuth := "sdasdasds"

	// 		c.JSON(http.StatusOK, gin.H{"token": returnedBasicAuth})
	// 		return
	// 	}
	// }

	// c.JSON(http.StatusUnauthorized, "Authenticate Failure!")
	c.Abort()
}

func (auth AuthenticationHandler) Authenticate(c *gin.Context) {
	if c.Param("username") != "" && c.Param("password") != "" {
		username := c.GetString("username")
		password := c.GetString("password")

		if username == "todo" && password == "aaa" {
			returnedBasicAuth := "sdasdasds"

			c.JSON(http.StatusOK, gin.H{"token": returnedBasicAuth})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, "Authenticate Failure!")
	c.Abort()
}
