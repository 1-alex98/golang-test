package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"trading/services/auth"
)

const UserKey = "user"

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserKey)
	if user == nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": "You must be logged in for this",
		})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

// Login is a handler that parses a form and checks for specific data
func Login(c *gin.Context) {
	email := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	if email == "" || password == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": "Login failed; missing parameters",
		})
		return
	}

	err := auth.Login(c, email, password)
	if err != nil {
		loginViewError(c, "Login failed")
		return
	}

	c.Redirect(303, "/")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserKey)
	if user == nil {
		c.Redirect(303, "/")
		return
	}
	session.Delete(UserKey)
	if err := session.Save(); err != nil {
		c.Redirect(303, "/")
		return
	}
	c.Redirect(303, "/")
}

func Me(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get(UserKey)
	c.JSON(http.StatusOK, gin.H{"email": email})
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

func LoginView(context *gin.Context) {
	loginViewError(context, "")
}
func loginViewError(context *gin.Context, error string) {
	context.HTML(http.StatusOK, "login.html", gin.H{"Error": error})
}
