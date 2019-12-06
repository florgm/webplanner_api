package tareas

type Tareas struct {
	IDUsuario  int64  `json:"id_usuario"`
	IDTarea    int64  `json:"id_tarea"`
	Tarea      string `json:tarea`
	Completado int8   `json:"completado"`
}
