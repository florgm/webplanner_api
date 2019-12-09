package sessions

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestValidateLoggedUsserOk(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	var idUsuario int64
	idUsuario = 1234
	c.Set("idUsuario", idUsuario)
	_, err := ValidateLoggedUser(c)
	if err != nil {
		t.Errorf("El usuario tendria que existir, Error: %v", err)
	}
}

func TestValidateLoggedUsserFail(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	_, err := ValidateLoggedUser(c)
	if err == nil {
		t.Errorf("El usuario no tendria que existir")
	}
}
