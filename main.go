package main

import (
	"trading/core"
)

// @title           Trading Golang API
// @version         1.0
// @description     This is a golang playground application

// @contact.email  alexander.von.trostorff@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	router := core.Setup()
	err := router.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
