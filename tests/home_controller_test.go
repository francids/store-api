package tests

import (
	"net/http"
	"net/http/httptest"
	"store-api/controllers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	// Configurar el router de Gin
	r := gin.Default()
	r.GET("/", controllers.Home)

	// Crear una solicitud HTTP
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Enviar la solicitud
	r.ServeHTTP(w, req)

	// Verificar el código de estado y la respuesta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"greeting": "Welcome to store-api"}`, w.Body.String())
}
