package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//handler is a function that handles an HTTT request
//based on Gin's HandlerFunc, to be use in order to check the health of the application

func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Estatus": "OK",
	})
}
