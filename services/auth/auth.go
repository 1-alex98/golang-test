package auth

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"trading/services/db"
)

const UserKey = "user"

func Login(c *gin.Context, email string, password string) (err error) {

	success := db.CheckCredentials(email, password)
	if !success {
		return errors.New("db check failed")
	}

	session := sessions.Default(c)

	// Save the email in the session
	session.Options(sessions.Options{
		SameSite: http.SameSiteStrictMode,
	})
	session.Set(UserKey, email) // In real world usage you'd set this to the users ID
	err = session.Save()
	return
}

func Register(c *gin.Context, email string, password string) (err error) {

	_, err = db.CreateUser(email, password)
	if err != nil {
		return errors.New("db check failed")
	}

	session := sessions.Default(c)

	// Save the email in the session
	session.Options(sessions.Options{
		SameSite: http.SameSiteStrictMode,
	})
	session.Set(UserKey, email) // In real world usage you'd set this to the users ID
	err = session.Save()
	return
}
