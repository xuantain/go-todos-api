package controllers

import (
	"go-todos-api/pkg/helpers"
	"go-todos-api/repositories"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationController struct {
	UserRepo repositories.UserRepository
}

var (
	JWTSecretKey = []byte("your-secret-key") // Replace with your secret or RSA key
	TokenExpiry  = time.Hour * 2
)

// Declare the required infos / claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type User struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

func GenerateToken(userID uint, role string) (string, error) {
	// Create a new JWT token with claims
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-todos-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// To ensure the integrity and authenticity of the token, we need to sign it using our secret key.
	return token.SignedString(JWTSecretKey)
}

func (ctl AuthenticationController) Login(c *gin.Context) {

	type LoginRequest struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Flash(c, "error", "StatusBadRequest")
		c.Redirect(http.StatusFound, "/page/login")
		// c.HTML(http.StatusBadRequest, "Index", gin.H{
		// 	"title":   "Login",
		// 	"message": err.Error(),
		// })
		return
	}

	user, err := ctl.UserRepo.FindByUsername(c, req.Username)

	if err != nil {
		helpers.Flash(c, "error", "Invalid username or credentials")
		c.Redirect(http.StatusFound, "/page/login")
		// c.HTML(http.StatusUnauthorized, "Index", gin.H{
		// 	"title":   "Login",
		// 	"message": "Invalid username or credentials",
		// })
		return
	}

	hash1 := helpers.HashStr(req.Password)

	// Todo: Should have a table to store user's token and recheck if it is expired
	// If it is not expired yet => allow access apis
	// If not => require authenticate
	if strings.Compare(hash1, user.Password) != 0 {
		helpers.Flash(c, "error", "Invalid username or credentials")
		c.Redirect(http.StatusFound, "/page/login")
		// c.HTML(http.StatusUnauthorized, "Index", gin.H{
		// 	"title":   "Login",
		// 	"message": "Invalid username or credentials",
		// })
		return
	}

	// Todo: Should get user's roles from DB => Create another userRoles & groupRoles & etc
	// Generate token
	token, err := GenerateToken(user.ID, "admin")
	if err != nil {
		helpers.Flash(c, "error", "Login Failed")
		c.Redirect(http.StatusFound, "/page/login")
		// c.HTML(http.StatusInternalServerError, "Index", gin.H{
		// 	"title":   "Login",
		// 	"message": "Login Failed",
		// })
		return
	}

	// Note: Use this for Html return as an option
	// userId := fmt.Sprintf("%d", user.ID)
	// c.SetCookie("token", token, int(TokenExpiry.Seconds()), "/", c.Request.Host, false, true)
	// c.SetCookie("userId", userId, int(TokenExpiry.Seconds()), "/", c.Request.Host, false, true)
	// c.SetCookie("username", user.Username, int(TokenExpiry.Seconds()), "/", c.Request.Host, false, true)

	// todo: create a model for this
	userData := []interface{}{
		user.ID,
		user.Username,
		user.Email,
		token,
		int(TokenExpiry.Seconds()),
	}

	helpers.SessionSet(c, "user", userData)
	helpers.Flash(c, "success", "Login Success")

	// IMPORTANT: Need to set http.StatusFound for the c.redirect here
	c.Redirect(http.StatusFound, "/page/welcome")
}

// todo: Do we need to store this to DB and refresh all the time?
var tokenBlacklist = make(map[string]bool)

func (ctl AuthenticationController) Logout(c *gin.Context) {

	// authHeader := c.GetHeader("Authorization")
	// tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	user := helpers.Auth(c)
	tokenString := user.Token
	tokenBlacklist[tokenString] = true

	helpers.SessionDelete(c, "user")

	c.Redirect(http.StatusTemporaryRedirect, "/page/login")
}

// func (ctl AuthenticationController) Logout(c *gin.Context) {
// 	// Clear the cookie
// 	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)
// 	c.SetCookie("userId", "", -1, "/", c.Request.Host, false, true)
// 	c.SetCookie("username", "", -1, "/", c.Request.Host, false, true)
// 	c.Redirect(http.StatusOK, "/page/")
// }

func RefreshToken(c *gin.Context) {
	// Extract claims from existing token
	// Verify it's not expired beyond grace period
	// Issue new token with same Token Blacklisting}
}
