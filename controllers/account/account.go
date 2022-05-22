package account

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"trading/services/db"
)

const UserKey = "user"

func GetAccountData(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get(UserKey)
	c.JSON(http.StatusOK, db.GetAccount(email.(string)))
}

func UpdateAccount(c *gin.Context) {
	//c.
	//	c.JSON(http.StatusOK, db.GetAccount(email.(string)))
}

func GetAccountView(c *gin.Context) {
	goods := db.Goods()
	session := sessions.Default(c)
	email := session.Get(UserKey)
	accountData := db.GetAccount(email.(string))
	user := db.GetUser(email.(string))

	var goodsList []struct {
		ID    uint
		Name  string
		Value float64
	}
	for _, good := range goods {
		goodsList = append(goodsList, struct {
			ID    uint
			Name  string
			Value float64
		}{
			ID:    good.ID,
			Name:  good.Name,
			Value: valueOfGood(accountData, good.ID),
		})
	}

	c.HTML(http.StatusOK, "account.html", gin.H{
		"Goods": goodsList,
		"User":  user,
	})
}

func valueOfGood(data []db.AccountEntry, id uint) float64 {
	index := slices.IndexFunc(data, func(ele db.AccountEntry) bool { return ele.GoodID == id })
	if index == -1 {
		return 0
	}
	return data[index].Value
}
