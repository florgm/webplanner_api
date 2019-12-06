package tareas

import (
    "encoding/json"
    "fmt"
    db2 "github.com/florgm/webplanner_api/src/api/db"
    "github.com/florgm/webplanner_api/src/api/domain/tareas"
)

//GetTareas esto es una funcion
func GetTareas() *[]tareas.Tareas {
    var (
        tarea  tareas.Tareas
        tareas []tareas.Tareas
    )

    rows, err := db2.Init().Query("select * from tareas;")

    if err != nil {
        fmt.Print(err.Error())
    }

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
func CreateTarea(tarea *tareas.Tareas) error {
    stmt, err := db2.Init().Prepare("insert into tareas (id_usuario, tarea, completado) values(?,?,?);")

    if err != nil {
        fmt.Print(err.Error())
    }

    _, err = stmt.Exec(tarea.IDUsuario, tarea.Tarea, tarea.Completado)

    defer stmt.Close()
    return err
}

//DeleteTarea esto es una funcion
func DeleteTarea(tarea *tareas.Tareas) error {
    id := tarea.IDTarea
    stmt, err := db2.Init().Prepare("delete from tareas where id_tarea = ?;")

    if err != nil {
        fmt.Print(err.Error())
    }

    _, err = stmt.Exec(id)

    defer stmt.Close()
    return err
}

//CompleteTarea esto es una funcion
func CompleteTarea(tarea *tareas.Tareas) error {
    stmt, err := db2.Init().Prepare("update tareas set completado=? where id_tarea=?;")

    if err != nil {
        fmt.Print(err.Error())
    }

    _, err = stmt.Exec(tarea.Completado, tarea.IDTarea)

    defer stmt.Close()
    return err
}
