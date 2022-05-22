package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trading/services/db"
)

func GetGoods(c *gin.Context) {
	goods := db.Goods()
	c.JSON(http.StatusOK, goods)
}

func GoodsView(c *gin.Context) {
	c.HTML(http.StatusOK, "goods.html", gin.H{
		"Goods": db.Goods(),
	})
}

func GoodView(c *gin.Context) {
	c.HTML(http.StatusOK, "good.html", gin.H{
		"Good": db.GoodById(c.Param("id")),
	})
}

func GoodCourse(c *gin.Context) {
	id := c.Param("id")
	good := db.GoodByIdPreloaded(id)
	c.JSON(http.StatusOK, good.DataPoints)
}
