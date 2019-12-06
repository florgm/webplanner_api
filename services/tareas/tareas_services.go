package tareas

import (
	"../../db"
	tareasDomain "../../domain/tareas"
	"encoding/json"
	"fmt"
)

func GetTareas() *[]tareasDomain.Tareas {
	var (
		tarea  tareasDomain.Tareas
		tareas []tareasDomain.Tareas
	)

	rows, err := db.Init().Query("select * from tareas;")

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

func ParseTarea(data []byte) (*tareasDomain.Tareas, error) {
	var tarea tareasDomain.Tareas
	if err := json.Unmarshal(data, &tarea); err != nil {
		return nil, err
	}

	return &tarea, nil
}

func CreateTarea(tarea *tareasDomain.Tareas) error {
	stmt, err := db.Init().Prepare("insert into tareas (id_usuario, tarea, completado) values(?,?,?);")

	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(tarea.IDUsuario, tarea.Tarea, tarea.Completado)

	defer stmt.Close()
	return err
}

func DeleteTarea(tarea *tareasDomain.Tareas) error {
	id := tarea.IDTarea
	stmt, err := db.Init().Prepare("delete from tareas where id_tarea = ?;")

	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(id)

	defer stmt.Close()
	return err
}

func CompleteTarea(tarea *tareasDomain.Tareas) error {
	stmt, err := db.Init().Prepare("update tareas set completado=? where id_tarea=?;")

	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(tarea.Completado, tarea.IDTarea)

	defer stmt.Close()
	return err
}
