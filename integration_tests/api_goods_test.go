package integration_tests

import (
	"testing"
	"trading/services/db"
)

func TestUserIs(t *testing.T) {
	_, err := db.CreateUser("integration_tests@example.com", "12345678")
	if err != nil {
		panic(err)
	}
	db.CheckCredentials("integration_tests@example.com", "12345678")
}
