package handlers

import (
	"go-todos-api/db"
	"go-todos-api/models"
	"go-todos-api/pkg/helpers"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationHandler struct{}

func (auth AuthenticationHandler) CheckBasicAuth(c *gin.Context) {
	username := c.MustGet(gin.AuthUserKey).(string)
	c.JSON(http.StatusOK, gin.H{"user": username, "message": "Welcome back!"})
}

var (
	JWTSecretKey = []byte("your-secret-key") // Replace with your secret or RSA key
	TokenExpiry  = time.Hour * 24
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

func authenticateUser(username, password string) (User, error) {
	// todo: This should be handled by DB layer
	var user models.User
	db := db.Conn()

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return User{}, err
	}

	hash1 := helpers.HashStr(password)

	if strings.Compare(hash1, user.Password) == 0 {
		// return user, nil
		return User{ID: user.ID, Role: "admin"}, nil
	}
	return User{}, gin.Error{}
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

func (auth AuthenticationHandler) Login(c *gin.Context) {
	type LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user (implement your own logic)
	user, err := authenticateUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate token
	token, err := GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Note: Use this for Html return as an option
	// c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"expires_in": TokenExpiry.Seconds(),
	})
}

var tokenBlacklist = make(map[string]bool)

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenBlacklist[tokenString] = true
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func RefreshToken(c *gin.Context) {
	// Extract claims from existing token
	// Verify it's not expired beyond grace period
	// Issue new token with same Token Blacklisting}
}
