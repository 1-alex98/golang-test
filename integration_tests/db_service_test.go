package integration_tests

import (
	"gotest.tools/assert"
	"testing"
	"trading/services/db"
)

func TestUserWrite(t *testing.T) {
	_, err := db.CreateUser("integration_tests@example.com", "12345678")
	if err != nil {
		panic(err)
	}
	db.CheckCredentials("integration_tests@example.com", "12345678")
}

func TestGoodWrite(t *testing.T) {
	db.CreateGood(db.Good{Name: "test", Description: "test description", CurrentCourse: 33})
	assert.Equal(t, db.Goods()[0].Name, "test")
	assert.Equal(t, db.Goods()[0].Description, "test description")
	assert.Equal(t, db.Goods()[0].CurrentCourse, 33.0)
}
