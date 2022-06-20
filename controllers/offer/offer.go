package offer

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trading/services/auth"
	"trading/services/db"
	"trading/services/offer"
)

type Order struct {
	Quantity float64
}

func GoodOffers(c *gin.Context) {
	id := c.Param("id")
	good := db.GoodByIdPreloaded(id)
	c.JSON(http.StatusOK, good.Offers)
}

func CreateOffer(c *gin.Context) {
	dbOffer := db.Offer{}
	err := c.BindJSON(&dbOffer)
	if err != nil {
		panic(err)
	}
	session := sessions.Default(c)
	email := session.Get(auth.UserKey)
	userId := db.GetUser(email.(string)).ID
	dbOffer.UserID = userId
	db.CreateOffer(dbOffer)
}

func OrderOffer(c *gin.Context) {
	id := c.Param("id")
	order := Order{}
	err := c.BindJSON(&order)
	if err != nil {
		c.Status(400)
		_ = c.Error(err)
		panic(err)
	}
	session := sessions.Default(c)
	email := session.Get(auth.UserKey)
	userId := db.GetUser(email.(string)).ID
	errService := offer.ProcessOrder(id, order.Quantity, strconv.Itoa(int(userId)))
	if errService != nil {
		c.Status(400)
		_ = c.Error(errService)
	}
}
