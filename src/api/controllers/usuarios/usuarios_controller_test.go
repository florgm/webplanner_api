package usuarios

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestLoginBadJson(t *testing.T) {
    c, _ := gin.CreateTestContext(httptest.NewRecorder())
    var err error

    body := strings.NewReader(`--usuario, "password": "abc123"}`)
    c.Request, err = http.NewRequest(http.MethodPost, "/login", body)
    if err != nil {
        t.Error(err)
	}
	
    Login(c)
    if c.Writer.Status() != http.StatusBadRequest {
        t.Error("Esperaba un status bad request");
    }
}
