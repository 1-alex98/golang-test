package offer

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"trading/services/auth"
	"trading/services/db"
)

func GoodOffers(c *gin.Context) {
	id := c.Param("id")
	good := db.GoodByIdPreloaded(id)
	c.JSON(http.StatusOK, good.Offers)
}

func CreateOffer(c *gin.Context) {
	offer := db.Offer{}
	err := c.BindJSON(&offer)
	if err != nil {
		panic(err)
	}
	session := sessions.Default(c)
	email := session.Get(auth.UserKey)
	userId := db.GetUser(email.(string)).ID
	offer.UserID = userId
	db.CreateOffer(offer)
}
