package middlewares

import (
	"go-todos-api/pkg/helpers"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func JwtSSAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		authUser := session.Get("user")

		if authUser == nil { // && c.Request.URL.Path != "/page/signin" && !strings.HasPrefix(c.Request.URL.Path, "/page/") {
			c.Redirect(http.StatusTemporaryRedirect, "/page/login")
			c.Abort()
			return
		}

		// slice authUser
		authUserStruct := helpers.AuthUser{
			Id:        authUser.([]interface{})[0].(uint),
			Username:  authUser.([]interface{})[1].(string),
			Email:     authUser.([]interface{})[2].(string),
			Token:     authUser.([]interface{})[3].(string),
			ExpiresAt: authUser.([]interface{})[4].(int),
		}

		c.Set("user", authUserStruct)
		c.Next()
	}
}

func Guest(c *gin.Context) {
	authUser := helpers.SessionGet(c, "user")
	if authUser != nil {
		c.Redirect(http.StatusFound, "/page/todos")
	}
	c.Next()
}
