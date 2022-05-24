package main

import (
	"trading/core"
)

func main() {
	router := core.Setup()
	err := router.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
