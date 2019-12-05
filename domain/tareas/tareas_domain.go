package tareas

type Tareas struct {
	IDUsuario   int64  `json:"id_usuario"`
	IDTarea     int64  `jason:"id_tarea"`
	Descripcion string `jason:"tarea"` //Tarea string `jason:tarea`
	Completado  int8   `jason:"completado"`
}

type IDTarea struct {
	IDTarea int64 `json:"id_evento"`
}
