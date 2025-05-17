package helpers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var defaultSessionOptions = sessions.Options{
	MaxAge: 60 * 5, // This sets session to expire in 60 * 5 seconds (5 minutes)
	// Add these settings to ensure proper cookie behavior:
	Path:     "/",  // Make cookie available for all paths
	HttpOnly: true, // Protect against XSS
	Secure:   true, // For HTTPS
	SameSite: http.SameSiteStrictMode,
}

func Flash(c *gin.Context, key string, message string) {
	session := sessions.Default(c)
	session.Delete(key)
	session.AddFlash(message, key)
	session.Save()
}

func FlashGet(c *gin.Context, key string) string {
	session := sessions.Default(c)
	message := session.Flashes(key)
	if len(message) == 0 {
		return ""
	}

	session.Save()

	var errorMessage string
	if len(message) > 0 {
		errorMessage = message[0].(string)
	} else {
		panic("Flash message not found")
	}

	return errorMessage
}

func SessionSet(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Delete(key)
	session.Set(key, value)
	session.Options(defaultSessionOptions)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session", "message": err.Error()})
		return
	}
}

func SessionGet(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	value := session.Get(key)
	return value
}

func SessionDelete(c *gin.Context, key string) {
	session := sessions.Default(c)
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	session.Delete(key)
	session.Save()
}
