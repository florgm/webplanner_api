package eventos

import (
	"testing"
)

func TestParseEventoFail(t *testing.T) {
	var data []byte
	_, err := ParseEvento(data)
	if err == nil {
		t.Errorf("Se esperaba un error")
	}
}

func TestParseEventoOk(t *testing.T) {
	b := []byte(`{"id_usuario": 1, "id_evento": 1, "title": "hola"}`)
	_, err := ParseEvento(b)
	if err != nil {
		t.Errorf("No se esperaba error: %v", err)
	}
}
