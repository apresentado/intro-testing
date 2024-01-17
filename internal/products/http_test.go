//handler - test

// la cantidad minima de casos de test es la cantidad de retorno

package products

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPProductsHandler(t *testing.T) {
	t.Run("Given an nvalid product identifier, debe retornar un error 400", func(t *testing.T) {
		//arrange
		var (
			expectedStatusCode = http.StatusBadRequest
			expectedHeaders    = http.Header{
				"content-type": []string{"application/json ; charset=utf-8"},
			}
			expectedBody = `{"error": "invalid product id"}`
			repository   = NewSliceRepository([]Products{
				{ID: 1, Name: "oreos"},
				{ID: 2, Name: "chocolate"},
				{ID: 3, Name: "milk"},
			})

			service = NewDefaultService(repository)
			handler = NewHTTPProductsHandler(service)
			// cuando se envia un id no numerico, se debe retornar un error 400
			request  = httptest.NewRequest(http.MethodGet, "/products/invalid", nil)
			response = httptest.NewRecorder()
		)
		engine := gin.New()
		engine.GET("/products/:id", handler.GetByID)

		//act - genero mi engine de gin
		engine.ServeHTTP(response, request)
		//assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedBody, response.Body.String())

})
}
