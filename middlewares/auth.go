package middlewares

import (
	"strings"

	"go-todos-api/config"
	"go-todos-api/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")

		var key string
		var secret string
		if key = config.GetString("http.auth.key"); len(strings.TrimSpace(key)) == 0 {
			c.AbortWithStatus(500)
		}
		if secret = config.GetString("http.auth.secret"); len(strings.TrimSpace(secret)) == 0 {
			c.AbortWithStatus(401)
		}

		isKeysEqual := helpers.SecureCompare(key, reqKey) == 1
		isSecretsEqual := helpers.SecureCompare(secret, reqSecret) == 1
		if !isKeysEqual || !isSecretsEqual {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
