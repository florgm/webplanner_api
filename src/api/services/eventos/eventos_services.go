package eventos

import (
	eventos "github.com/florgm/webplanner_api/src/api/domain/eventos"
	"github.com/florgm/webplanner_api/src/api/utils/apierror"
	db "github.com/florgm/webplanner_api/src/api/db"
	"encoding/json"
	"io/ioutil"
    "net/http"
    "strconv"
	"time"
)

//GetEventos trae los eventos guardados del usuario y los feriados
func GetEventos(user int64) (*[]eventos.Eventos,*apierror.ApiError) {
	channel := make(chan []eventos.Eventos)
	channelErrors := make(chan apierror.ApiError)
	var errors []apierror.ApiError

	go getFeriados(channel,channelErrors,user,2019)
    go getFeriados(channel,channelErrors,user,2020)
    go getEventosDB(channel,channelErrors,user)

    eventos := <- channel
	eventos = append(eventos, <- channel...)
	eventos = append(eventos, <- channel...)
	
	errors = append(errors,<- channelErrors)
	errors = append(errors,<- channelErrors)
	errors = append(errors,<- channelErrors)
	
	for i := 0; i < len(errors); i++ {
		if(errors[i].Status != http.StatusOK) {
			return &eventos, &errors[i]
		}
	}

    return &eventos, nil
}

func getFeriados(channel chan []eventos.Eventos, errors chan apierror.ApiError, user int64, year int){
    url := "http://nolaborables.com.ar/api/v2/feriados/"
    yearStr := strconv.Itoa(year)
    url = url + yearStr

    var(
		evento eventos.Eventos
		feriados []eventos.EventosFeriados
        eventos []eventos.Eventos
    )

    resp, err := http.Get(url)
    if err != nil {
		apiError := apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while getting the external api data",
		}
		errors <- apiError
		channel <- eventos
		return
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        apiError := apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while reading the external api data",
		}
		errors <- apiError
		channel <- eventos
		return
    }

    err = json.Unmarshal(data, &feriados)
    if err != nil {
        apiError := apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while doing unmarshal of the external api data",
		}
		errors <- apiError
		channel <- eventos
		return
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

	apiError := apierror.ApiError {
		Status: http.StatusOK,
		Message: "nil",
	}

	channel <- eventos
	errors <- apiError
}

func getEventosDB(channel chan []eventos.Eventos, errors chan apierror.ApiError, user int64){
    var (
        evento  eventos.Eventos
        eventos []eventos.Eventos
    )
    stmt, err := db.Init().Prepare("select * from eventos where id_usuario = ?;")
    if err != nil {
        apiError := apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while getting the external api data",
		}
		errors <- apiError
		channel <- eventos
		return
    }

    rows, err := stmt.Query(user)

    if err != nil {
        apiError := apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while getting the external api data",
		}
		errors <- apiError
		channel <- eventos
		return
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

        // if err != nil {
        //     fmt.Print(err.Error())
        // }
    }

	defer rows.Close()
	
	apiError := apierror.ApiError {
		Status: http.StatusOK,
		Message: "nil",
	}

	channel <- eventos
	errors <- apiError
}

//ParseEvento esto es una funcion
func ParseEvento(data []byte) (*eventos.Eventos, error) {
    var evento eventos.Eventos
    if err := json.Unmarshal(data, &evento); err != nil {
        return nil, err
    }

    return &evento, nil
}

//CreateEvento crea un evento en la base de datos
func CreateEvento(evento *eventos.Eventos, user int64) *apierror.ApiError {
    stmt, err := db.Init().Prepare("insert into eventos (id_usuario, title, descripcion, color, textColor, start, end) values(?,?,?,?,?,?,?);")
    if err != nil {
		return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
    }

    _, err = stmt.Exec(user, evento.Title, evento.Descripcion, evento.Color, evento.TextColor, evento.Start, evento.End)
	if err != nil {
        return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while saving the event data",
		}
	}

    defer stmt.Close()
    return nil
}

//DeleteEvento elmina un evento de la base de datos
func DeleteEvento(evento *eventos.Eventos) *apierror.ApiError {
	id := evento.IDEvento
	
	stmt, err := db.Init().Prepare("delete from eventos where id_evento = ?;")
	if err != nil {
		return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
    }
	
	_, err = stmt.Exec(id)
	if err != nil {
        return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while deleting the event data",
		}
	}

    defer stmt.Close()
    return nil
}

//UpdateEvento modifica un evento de la base de datos
func ModifyEvento(evento *eventos.Eventos) *apierror.ApiError {
    stmt, err := db.Init().Prepare("update eventos set title=?, descripcion=?, color=?, textColor=?, start=?, end=? where id_evento=?;")
	if err != nil {
		return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
    }

    _, err = stmt.Exec(evento.Title, evento.Descripcion, evento.Color, evento.TextColor, evento.Start, evento.End, evento.IDEvento)
	if err != nil {
        return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Error while modifying the event data",
		}
	}

    defer stmt.Close()
    return nil
}

