//handler:

package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPProductsHandlerGetByID struct {
	ID int
	Name string
}


//HTTPProductsHandler is a HTTP handler that exposes the products service

type HTTPProductsHandler struct {
	// service is used to access the products service
	service Service
}

// NewHTTPProductsHandler creates a new instance of the HTTPProductsHandler

func NewHTTPProductsHandler(service Service) *HTTPProductsHandler {
	return &HTTPProductsHandler{
		service: service,
	}
}

// GetByID handles a HTTP request to get a product by its ID
//reciVE a product id as a path paramether (:id)
func (h *HTTPProductsHandler) GetByID(ctx *gin.Context) {
	// get the product id from the path parameter
	productID, err := strconv.Atoi(ctx.Param("id"))
	// get the product from the service
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}
	//delegar la llamada al servicio
	product, err := h.service.GetByID(ctx, productID)

	if err != nil {
		switch err {
		case errNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "an unexpected error occurred",
			})
	}
	return
	}
	// return the product as a JSON response
	ctx.JSON(http.StatusOK, HTTPProductsHandlerGetByID{
		ID: product.ID,
		Name: product.Name,
	})
}
