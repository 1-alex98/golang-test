package integration_tests

import (
	"github.com/goccy/go-json"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"trading/services/db"
)

func TestGoodsApi(t *testing.T) {
	db.CreateGood(db.Good{Name: "test 2", Description: "test description 2", CurrentCourse: 33})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/goods", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	good := make([]db.Good, 1)
	err := json.Unmarshal([]byte(w.Body.String()), &good)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, good[0].Name, "test")
	assert.Equal(t, good[0].Description, "test description")
	assert.Equal(t, good[0].CurrentCourse, 33.0)
}
