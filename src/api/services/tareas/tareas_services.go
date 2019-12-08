package tareas

import (
    "encoding/json"
    "fmt"
    db "github.com/florgm/webplanner_api/src/api/db"
    "github.com/florgm/webplanner_api/src/api/domain/tareas"
)

//GetTareas esto es una funcion
func GetTareas(user int64) *[]tareas.Tareas {
    var (
        tarea  tareas.Tareas
        tareas []tareas.Tareas
    )

	stmt, err := db.Init().Prepare("select * from tareas where id_usuario = ?;")

    if err != nil {
        fmt.Print(err.Error())
	}
	
	rows, err := stmt.Query(user)

    for rows.Next() {
        err = rows.Scan(&tarea.IDUsuario,
            &tarea.IDTarea,
            &tarea.Tarea,
            &tarea.Completado)
        tareas = append(tareas, tarea)

        if err != nil {
            fmt.Print(err.Error())
        }

    }

    defer rows.Close()

    return &tareas
}

//ParseTarea esto es una funcion
func ParseTarea(data []byte) (*tareas.Tareas, error) {
    var tarea tareas.Tareas
    if err := json.Unmarshal(data, &tarea); err != nil {
        return nil, err
    }

    return &tarea, nil
}

//CreateTarea esto es una funcion
func CreateTarea(tarea *tareas.Tareas, user int64) (idTarea int64, err error) {
    stmt, err := db.Init().Prepare("insert into tareas (id_usuario, tarea, completado) values(?,?,?);")
    if err != nil {
        fmt.Print(err.Error())
    }

    res, err := stmt.Exec(user, tarea.Tarea, tarea.Completado)
	if err != nil {
        fmt.Print(err.Error())
	}
	
	id, err := res.LastInsertId()

    defer stmt.Close()
    return id, err
}

//DeleteTarea esto es una funcion
func DeleteTarea(tarea *tareas.Tareas) error {
    id := tarea.IDTarea
    stmt, err := db.Init().Prepare("delete from tareas where id_tarea = ?;")

    if err != nil {
        fmt.Print(err.Error())
    }

    _, err = stmt.Exec(id)

    defer stmt.Close()
    return err
}

//CompleteTarea esto es una funcion
func CompleteTarea(tarea *tareas.Tareas) error {
    stmt, err := db.Init().Prepare("update tareas set completado=? where id_tarea=?;")
    if err != nil {
        fmt.Print(err.Error())
    }

	_, err = stmt.Exec(tarea.Completado, tarea.IDTarea)
    defer stmt.Close()
    return err
}

//SetTarea funcion
func SetTarea(user int64, idTarea int64, tareastr string, completado int8) *tareas.Tareas {
	var tarea tareas.Tareas
	tarea.IDUsuario = user
	tarea.IDTarea = idTarea
	tarea.Tarea = tareastr
	tarea.Completado = completado

	return &tarea
}
