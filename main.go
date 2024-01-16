package main

import (
	"github.com/gin-gonic/gin"
	"http.go/internal/healthcheck"
)

func main() {
	//create a new gin router
	router := gin.Default()

	//add a route to the router
	router.GET("/ping", healthcheck.Handler)

	//execute my router
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
