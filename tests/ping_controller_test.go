package tests

import (
	"net/http"
	"net/http/httptest"
	"store-api/controllers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Configurar el router de Gin
	r := gin.Default()
	r.GET("/ping", controllers.Ping)

	// Crear una solicitud HTTP
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	// Enviar la solicitud
	r.ServeHTTP(w, req)

	// Verificar el código de estado y la respuesta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"ping": "pong"}`, w.Body.String())
}
