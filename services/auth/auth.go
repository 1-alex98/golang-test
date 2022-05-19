package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

const UserKey = "user"

func Login(c *gin.Context, username string, password string) (error string) {

	//TODO: check login

	session := sessions.Default(c)

	// Save the username in the session
	session.Options(sessions.Options{
		SameSite: http.SameSiteStrictMode,
	})
	session.Set(UserKey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		return "Failed to save session"
	}
	return
}
