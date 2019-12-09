package tareas

import (
	"testing"
)

func TestParseTareaFail(t *testing.T) {
	var data []byte
	_, err := ParseTarea(data)
	if err == nil {
		t.Errorf("Se esperaba un error")
	}
}

func TestParseTareaOk(t *testing.T) {
	b := []byte(`{"id_usuario": 1, "id_tarea": 1, "title": "hola"}`)
	_, err := ParseTarea(b)
	if err != nil {
		t.Errorf("No se esperaba error: %v", err)
	}
}
