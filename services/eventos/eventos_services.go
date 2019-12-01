package eventos

import (
	"fmt"
	"../../db"
	"encoding/json"
	eventosDomain "../../domain/eventos"
)

//GetEventos funcion que trae todos los eventos guardados en la base de datos
func GetEventos() *[]eventosDomain.Eventos {
	var (
		evento eventosDomain.Eventos
		eventos []eventosDomain.Eventos
	)

	rows, err := db.Init().Query("select * from eventos;")

	if err != nil {
		fmt.Print(err.Error())
	}
	
	for rows.Next() {
		err = rows.Scan(&evento.IDUsuario, 
						&evento.IDEvento, 
						&evento.Title, 
						&evento.Descripcion, 
						&evento.Color, 
						&evento.TextColor,
						&evento.Start,
						&evento.End)

		eventos = append(eventos, evento)
		
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	
	defer rows.Close()
	
	return &eventos
}

//ParseEvento esto es una funcion
func ParseEvento(data []byte) (*eventosDomain.Eventos, error) {
    var evento eventosDomain.Eventos
    if err := json.Unmarshal(data, &evento); err != nil {
        return nil, err
	}
	
    return &evento, nil
}

//ParseEventoID esto es una funcion
func ParseEventoID(data []byte) (*eventosDomain.IDEvento, error) {
    var evento eventosDomain.IDEvento
    if err := json.Unmarshal(data, &evento); err != nil {
        return nil, err
	}
	
    return &evento, nil
}

//CreateEvento funcion para insertar eventos en la base de datos
func CreateEvento(evento *eventosDomain.Eventos) error {
	stmt, err := db.Init().Prepare("insert into eventos (id_usuario, title, descripcion, color, textColor, start, end) values(?,?,?,?,?,?,?);")

	if err != nil {
		fmt.Print(err.Error())
	}
	
	_, err = stmt.Exec(evento.IDUsuario, evento.Title, evento.Descripcion, evento.Color, evento.TextColor, evento.Start, evento.End)
	
	defer stmt.Close()
	return err
}

//DeleteEvento esto es una funcion
func DeleteEvento(evento *eventosDomain.IDEvento) error {
	id := evento.IDEvento
	stmt, err := db.Init().Prepare("delete from eventos where id_evento = ?;")

	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	
	defer stmt.Close()
	return err
}

//ModifyEvento funcion para modificar eventos en la base de datos
func ModifyEvento(evento *eventosDomain.Eventos) error {
	stmt, err := db.Init().Prepare("update eventos set title=?, descripcion=?, color=?, textColor=?, start=?, end=? where id_evento=?;")

	if err != nil {
		fmt.Print(err.Error())
	}
	
	_, err = stmt.Exec(evento.Title, evento.Descripcion, evento.Color, evento.TextColor, evento.Start, evento.End, evento.IDEvento)
	
	defer stmt.Close()
	return err
}