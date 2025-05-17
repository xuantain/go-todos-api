package middlewares

import (
	"go-todos-api/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtSSAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authUser := helpers.SessionGet(c, "user")
		if authUser == nil || authUser == "" { // && c.Request.URL.Path != "/page/signin" && !strings.HasPrefix(c.Request.URL.Path, "/page/") {
			// NOTE: c.Redirect with http.StatusTemporaryRedirect would foward the request.method also
			// EX: Request to /update/abc with method POST, if the handler redirect to another url=/delete/abc by http.StatusTemporaryRedirect
			// 	it would delegate the request to url=/delete/abc with method POST
			c.Redirect(http.StatusTemporaryRedirect, "/page/login")
			c.Abort()
			return
		}

		c.Next()
	}
}

func Guest(c *gin.Context) {

	authUser := helpers.Auth(c)
	if authUser.Username != "" {
		c.Redirect(http.StatusFound, "/page/todos")
	}

	c.Next()
}
