package eventos

import (
    "encoding/json"
    "fmt"
    db "github.com/florgm/webplanner_api/src/api/db"
	"github.com/florgm/webplanner_api/src/api/domain/eventos"
    "io/ioutil"
    eventos2 "src/github.com/florgm/webplanner_api/src/api/domain/eventos"
    "src/github.com/florgm/webplanner_api/src/api/utils/apierror"
    "net/http"
    "strconv"
    "time"
)

//GetEventos funcion que trae todos los eventos guardados en la base de datos
func GetEventos(user int64) *[]eventos.Eventos {

    channel := make(chan []eventos.Eventos)



	go getFeriados(channel,user,2019)
    go getFeriados(channel,user,2020)
    go getEventosDB(channel,user)



    eventos := <- channel
	eventos = append(eventos, <- channel)
    eventos = append(eventos, <- channel)

    return &eventos
}

func getFeriados(channel chan []eventos.Eventos,user int64, year int){

    url := "http://nolaborables.com.ar/api/v2/feriados/"
    yearStr := strconv.Itoa(year)

    url = url + yearStr

    var(
        eventos []Eventos
        evento Eventos
        feriados []EventosFeriados
    )

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }

    err = json.Unmarshal(data, &feriados)
    if err != nil {
        fmt.Println(err)
    }

    for u := 0; u < len(feriados); u++ {
        evento.IDUsuario = user
        evento.IDEvento = 0
        evento.Title = feriados[u].Motivo
        evento.Descripcion = feriados[u].ID
        evento.Color = "#2ECC71"
        evento.TextColor = "#FFFFFF"
        evento.Start = time.Date(year,time.Month(feriados[u].Mes),feriados[u].Dia,0,0,0,0,time.UTC)
        evento.End = evento.Start

        eventos = append(eventos, evento)

    }

    channel <- eventos

}

func getEventosDB (channel chan []eventos.Eventos, user int64){

    var (
        evento  eventos.Eventos
        eventos []eventos.Eventos

    )
    stmt, err := db.Init().Prepare("select * from eventos where id_usuario = ?;")
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

    channel <- eventos
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
func CreateEvento(evento *eventos.Eventos, user int64) error {
    stmt, err := db.Init().Prepare("insert into eventos (id_usuario, title, descripcion, color, textColor, start, end) values(?,?,?,?,?,?,?);")

    if err != nil {
		fmt.Print(err.Error())
		return err
    }

    _, err = stmt.Exec(user, evento.Title, evento.Descripcion, evento.Color, evento.TextColor, evento.Start, evento.End)

    defer stmt.Close()
    return err
}

//DeleteEvento esto es una funcion
func DeleteEvento(evento *eventos.Eventos) error {
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
func ModifyEvento(evento *eventos.Eventos) error {
    stmt, err := db.Init().Prepare("update eventos set title=?, descripcion=?, color=?, textColor=?, start=?, end=? where id_evento=?;")

    if err != nil {
        fmt.Print(err.Error())
    }

    _, err = stmt.Exec(evento.Title, evento.Descripcion, evento.Color, evento.TextColor, evento.Start, evento.End, evento.IDEvento)

    defer stmt.Close()
    return err
}

