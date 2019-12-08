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

type EventosFeriados struct {
	Motivo 	 string `json:"motivo"` // String
	Tipo   	 string `json:"tipo"`   // String // inamovible | trasladable | nolaborable | puente
	Dia    	 int  	`json:"dia"`    // Number // Día del mes
	Mes    	 int  	`json:"mes"`    // Number // Número de mes en base 1 (enero = 1)
	ID     	 string `json:"id"`     // String // Identificador único de feriado
	Original string `json:"original"` // String // Fecha original en formato DD-MM
}
