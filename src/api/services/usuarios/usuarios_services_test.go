package usuarios

import (
	"testing"
)

func TestParseUsuarioFail(t *testing.T) {
	var data []byte
	_, err := ParseUsuario(data)
	if err == nil {
		t.Errorf("Se esperaba un error")
	}
}

func TestParseUsuarioOk(t *testing.T) {
	b := []byte(`{"id_usuario": 1, "id_tarea": 1, "title": "hola"}`)
	_, err := ParseUsuario(b)
	if err != nil {
		t.Errorf("No se esperaba error: %v", err)
	}
}

func TestParseLoginUsuarioFail(t *testing.T) {
	var data []byte
	_, err := ParseLoginUsuario(data)
	if err == nil {
		t.Errorf("Se esperaba un error")
	}
}

func TestParseLoginUsuarioOk(t *testing.T) {
	b := []byte(`{"id_usuario": 1, "id_tarea": 1, "title": "hola"}`)
	_, err := ParseLoginUsuario(b)
	if err != nil {
		t.Errorf("No se esperaba error: %v", err)
	}
}
