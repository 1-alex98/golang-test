package main

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"trading/controllers/account"
	"trading/controllers/auth"
	"trading/controllers/goods"
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
	router.Run() // listen and serve on 0.0.0.0:8080#
}

func templateFunctions(router *gin.Engine) {
	router.SetFuncMap(template.FuncMap{
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
	})
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
	router.POST("/login", auth.Login)
	router.POST("/logout", auth.Logout)
	router.GET("/api/goods", goods.GetGoods)
	router.GET("/goods", goods.GoodsView)
	router.GET("/goods/:id", goods.GoodView)
	router.GET("/api/goods/:id/course", goods.GoodCourse)
}

func private(router *gin.Engine) {
	private := router.Group("/private")
	private.Use(auth.AuthRequired)
	{
		private.GET("/me", auth.Me)
		private.GET("/status", auth.Status)
		private.GET("/api/account", account.GetAccountData)
		private.PUT("/api/account/:id", account.UpdateAccount)
		private.GET("/account", account.GetAccountView)
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
