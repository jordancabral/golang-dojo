package main

import (
	"fmt"

	"github.com/jordancabral/golang-dojo/app/routers"
)

func main() {
	fmt.Println("Starting Mock Server")

	// Init Gin
	r := routers.InitRouter()

	// Run server
	r.Run()

}
