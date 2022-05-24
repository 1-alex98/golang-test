package test

import (
	"testing"
	"trading/services/db"
)

func TestUserIs(t *testing.T) {
	_, err := db.CreateUser("test@example.com", "12345678")
	if err != nil {
		panic(err)
	}
	db.CheckCredentials("test@example.com", "12345678")
}
