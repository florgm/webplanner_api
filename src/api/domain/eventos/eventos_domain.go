package eventos

import "time"

//Eventos estructura de la base de datos
type Eventos struct {
	IDUsuario   int64     `json:"id_usuario"`
	IDEvento    int64     `json:"id_evento"`
	Title       string    `json:"title"`
	Descripcion string    `json:"descripcion"`
	Color       string    `json:"color"`
	TextColor   string    `json:"textColor"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}
