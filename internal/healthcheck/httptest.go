package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	
	t.Run("it should return an HTTP 200 status code", func(t *testing.T) {
		//arrage  - expectativas
		var (
			expectedStatusCode = http.StatusOK
			expectedHeaders    = http.Header{
				"content-type": []string{"application/json ; charset=utf-8"},
			}
			expectedBody = `{"status":"ok"}`

			request = httptest.NewRequest(http.MethodGet, "/ping", nil)
			//donde almacenar la respuesta
			response = httptest.NewRecorder()
		)

		//generar router -- por lo general, este punto debe ser en una estructura de routes
		engine := gin.New() //es igual que .default pero sin middlewares
		engine.GET("/ping", Handler)

		//act - ejecutar
		engine.ServeHTTP(response, request) //para ejecutar el handler

		//assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())    // los headers que espero son iguales a los de la respuesta?
		assert.JSONEq(t, expectedBody, response.Body.String()) // nos permite comparar los jsons distintos
	})
}
