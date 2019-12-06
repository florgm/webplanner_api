package eventos

import (
    "encoding/json"
    "fmt"
    db2 "github.com/florgm/webplanner_api/src/api/db"
    "github.com/florgm/webplanner_api/src/api/domain/eventos"
)

//GetEventos funcion que trae todos los eventos guardados en la base de datos
func GetEventos(user int64) *[]eventos.Eventos {
    var (
        evento  eventos.Eventos
        eventos []eventos.Eventos
    )
    stmt, err := db2.Init().Prepare("select * from eventos where id_usuario = ?;")
    if err != nil {
        fmt.Print(err.Error())
    }

    rows, err := stmt.Query(user)

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
func ParseEvento(data []byte) (*eventos.Eventos, error) {
    var evento eventos.Eventos
    if err := json.Unmarshal(data, &evento); err != nil {
        return nil, err
    }

    return &evento, nil
}

//CreateEvento funcion para insertar eventos en la base de datos
func CreateEvento(evento *eventos.Eventos) error {
    stmt, err := db2.Init().Prepare("insert into eventos (id_usuario, title, descripcion, color, textColor, start, end) values(?,?,?,?,?,?,?);")

    if err != nil {
        fmt.Print(err.Error())
    }

    _, err = stmt.Exec(evento.IDUsuario, evento.Title, evento.Descripcion, evento.Color, evento.TextColor, evento.Start, evento.End)

    defer stmt.Close()
    return err
}

//DeleteEvento esto es una funcion
func DeleteEvento(evento *eventos.Eventos) error {
    id := evento.IDEvento
    stmt, err := db2.Init().Prepare("delete from eventos where id_evento = ?;")

    if err != nil {
        fmt.Print(err.Error())
    }
    _, err = stmt.Exec(id)

    defer stmt.Close()
    return err
}

//ModifyEvento funcion para modificar eventos en la base de datos
func ModifyEvento(evento *eventos.Eventos) error {
    stmt, err := db2.Init().Prepare("update eventos set title=?, descripcion=?, color=?, textColor=?, start=?, end=? where id_evento=?;")

    if err != nil {
        fmt.Print(err.Error())
    }

    _, err = stmt.Exec(evento.Title, evento.Descripcion, evento.Color, evento.TextColor, evento.Start, evento.End, evento.IDEvento)

    defer stmt.Close()
    return err
}
