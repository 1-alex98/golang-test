package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"trading/controllers/account"
	"trading/controllers/auth"
	"trading/controllers/goods"
	"trading/controllers/offer"
	"trading/services/course"
	"trading/services/db"
)

var secret = []byte("secret")

func main() {
	db.Init()
	course.Init()
	router := gin.Default()
	router.Use(gin.Logger())
	store := cookie.NewStore(secret)
	router.Use(sessions.Sessions("session", store))
	templateFunctions(router)
	router.LoadHTMLGlob("templates/*.html")
	routes(router)
	err := router.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}

func routes(router *gin.Engine) {
	router.Static("/public", "./public")
	public(router)
	private(router)
}

func public(router *gin.Engine) {
	router.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "error.html", gin.H{"Error": "Not found..."})
	})
	router.GET("/", index)
	router.GET("/login", auth.LoginView)
	router.GET("/register", auth.RegisterView)
	router.POST("/login", auth.Login)
	router.POST("/register", auth.Register)
	router.POST("/logout", auth.Logout)
	router.GET("/api/goods", goods.GetGoods)
	router.GET("/goods", goods.GoodsView)
	router.GET("/goods/:id", goods.GoodView)
	router.GET("/api/goods/:id/course", goods.GoodCourse)
	router.GET("/api/goods/:id/offers", offer.GoodOffers)
}

func private(router *gin.Engine) {
	private := router.Group("/private")
	private.Use(auth.AuthRequired)
	{
		private.GET("/me", auth.Me)
		private.GET("/status", auth.Status)
		private.GET("/api/account", account.GetAccountData)
		private.POST("/api/offer", offer.CreateOffer)
		private.PUT("/api/account/:id", account.UpdateAccount)
		private.PUT("/api/credit", account.UpdateCredit)
		private.GET("/account", account.GetAccountView)
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
