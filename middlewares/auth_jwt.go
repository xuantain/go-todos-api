package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWTSecretKey = []byte("your-secret-key") // Replace with your secret or RSA key
	TokenExpiry  = time.Hour * 24
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Note: Retrieve the token from the Cookie - for the Html return as an option
		// tokenString, err := c.Cookie("token")
		// if err != nil {
		// 	fmt.Println("Token missing in Cookie")
		// 	c.Redirect(http.StatusSeeOther, "/login")
		// 	c.Abort()
		// 	return
		// }

		// Note: Retrieve the token from the Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token missing"})
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JWTSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Add user info to context
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}
