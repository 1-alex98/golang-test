package integration_tests

import (
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"trading/services/db"
)

var w *httptest.ResponseRecorder

//Fuzzing is incompatible with test containers :(
func FuzzLoginMain(f *testing.F) {
	db.CreateUser("test@example.de", "banana")
	f.Add("banana")
	w = httptest.NewRecorder()
	f.Fuzz(Login)
}

func Login(t *testing.T, password string) {
	req, _ := http.NewRequest("POST", "/login", nil)
	form := url.Values{}
	form.Add("username", "test@example.de")
	form.Add("password", password)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	var expectedStatus int
	if password == "banana" {
		expectedStatus = 303
	} else {
		expectedStatus = 400
	}
	assert.Equal(t, expectedStatus, w.Code)
}
