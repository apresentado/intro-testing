package main

import (
	"github.com/gin-gonic/gin"
	"http.go/internal/healthcheck"
	"http.go/internal/products"
)

func main() {
	var err error

	//generate handlers
	productsRepository := products.NewSliceRepository([]products.Products{
		{ID: 1, Name: "Oreos"},
		{ID: 2, Name: "Milk"},
		{ID: 3, Name: "Bread"},
	})
	// genero mi servicio
	productService := products.NewDefaultService(productsRepository)
	productHandler := products.NewHTTPProductsHandler(productService)
	//create a new gin router
	router := gin.Default()

	//add a route to the router
	router.GET("/ping", healthcheck.Handler)
	router.GET("/products/:id", productHandler.GetByID)

	//execute my router
	if err = router.Run(":8080"); err != nil {
		panic(err)
	}
}
