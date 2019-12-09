package rest

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetJSONBodyFail(t *testing.T) {
	request := httptest.NewRequest("GET", "/", nil)
	request.Body = nil
	_, err := GetJSONBody(request)
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestGetJSONBodyOk(t *testing.T) {
	request := httptest.NewRequest("GET", "/", strings.NewReader("algo"))
	_, err := GetJSONBody(request)
	if err != nil {
		t.Error("Se esperaba un error")
	}
}
