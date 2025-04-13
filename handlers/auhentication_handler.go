package handlers

import (
	"context"
	"go-todos-api/pkg/helpers"
	"go-todos-api/repositories"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationHandler struct {
	userRepo repositories.UserRepository
}

func NewAuthenticationHandler(userRepo repositories.UserRepository) *AuthenticationHandler {
	return &AuthenticationHandler{userRepo: userRepo}
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

func (auth AuthenticationHandler) CheckBasicAuth(c *gin.Context) {
	username := c.MustGet(gin.AuthUserKey).(string)
	c.JSON(http.StatusOK, gin.H{"user": username, "message": "Welcome back!"})
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

	user, err := auth.userRepo.FindByUsername(context.Background(), req.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or credentials"})
		return
	}

	hash1 := helpers.HashStr(req.Password)

	// Todo: Should have a table to store user's token and recheck if it is expired
	// If it is not expired yet => allow access apis
	// If not => require authenticate
	if strings.Compare(hash1, user.Password) != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or credentials"})
		return
	}

	// Todo: Should get user's roles from DB => Create another userRoles & groupRoles & etc
	// Generate token
	token, err := GenerateToken(user.ID, "admin")
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
