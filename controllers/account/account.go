package account

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"strconv"
	"trading/services/auth"
	"trading/services/db"
)

func GetAccountData(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get(auth.UserKey)
	c.JSON(http.StatusOK, db.GetAccount(email.(string)))
}

type Update struct {
	Value float64 `form:"Value" json:"Value" binding:"required"`
}

func UpdateAccount(c *gin.Context) {
	var json Update
	err := c.BindJSON(&json)
	if err != nil {
		panic(err)
	}
	goodId := c.Param("id")
	session := sessions.Default(c)
	email := session.Get(auth.UserKey)
	userId := db.GetUser(email.(string)).ID
	goodIdInt, err := strconv.Atoi(goodId)
	if err != nil {
		panic(err)
	}
	db.UpdateAccount(json.Value, uint(goodIdInt), userId)
}

// UpdateCredit godoc
// @Summary      Update your credit
// @Description  api to set a self chosen new account value
// @Accept       json
// @Param        Update  body      Update  true  "Add account value"
// @Success      200
// @Router       /api/credit [post]
func UpdateCredit(c *gin.Context) {
	var json Update
	err := c.BindJSON(&json)
	if err != nil {
		panic(err)
	}
	session := sessions.Default(c)
	email := session.Get(auth.UserKey)
	db.UpdateCredit(email.(string), json.Value)
}

func GetAccountView(c *gin.Context) {
	goods := db.Goods()
	session := sessions.Default(c)
	email := session.Get(auth.UserKey)
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
