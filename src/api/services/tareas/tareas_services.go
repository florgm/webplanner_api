package tareas

import (
	"encoding/json"
	"net/http"

	db "github.com/florgm/webplanner_api/src/api/db"
	"github.com/florgm/webplanner_api/src/api/domain/tareas"
	"github.com/florgm/webplanner_api/src/api/utils/apierror"
)

//GetTareas devuelve todas las tareas del usuario logueado
func GetTareas(user int64) (*[]tareas.Tareas, *apierror.ApiError) {
	var (
		tarea  tareas.Tareas
		tareas []tareas.Tareas
	)

	stmt, err := db.Init().Prepare("select * from tareas where id_usuario = ?;")
	if err != nil {
		return nil, &apierror.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "Data base error",
		}
	}

	rows, err := stmt.Query(user)
	// if err != nil {
	// 	return nil, &apierror.ApiError{
	// 		Status:  http.StatusNotFound,
	// 		Message: "Incorrect user",
	// 	}
	// }

	for rows.Next() {
		err = rows.Scan(&tarea.IDUsuario,
			&tarea.IDTarea,
			&tarea.Tarea,
			&tarea.Completado)
		tareas = append(tareas, tarea)

		if err != nil {
			return nil, &apierror.ApiError{
				Status:  http.StatusNotFound,
				Message: "The user has no tasks",
			}
		}

	}

	defer rows.Close()
	return &tareas, nil
}

func ParseTarea(data []byte) (*tareas.Tareas, error) {
	var tarea tareas.Tareas
	if err := json.Unmarshal(data, &tarea); err != nil {
		return nil, err
	}

	return &tarea, nil
}

//CreateTarea crea una nueva tarea para el usuario logueado
func CreateTarea(tarea *tareas.Tareas, user int64) (int64, *apierror.ApiError) {
	stmt, err := db.Init().Prepare("insert into tareas (id_usuario, tarea, completado) values(?,?,?);")
	if err != nil {
		return 0, &apierror.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "Data base error",
		}
	}

	res, err := stmt.Exec(user, tarea.Tarea, tarea.Completado)
	if err != nil {
		return 0, &apierror.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "Error while saving the task data",
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, &apierror.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "Error with the last inserted id",
		}
	}

	defer stmt.Close()
	return id, nil
}

//DeleteTarea borra una tarea especifica
func DeleteTarea(tarea *tareas.Tareas) *apierror.ApiError {
	id := tarea.IDTarea

	stmt, err := db.Init().Prepare("delete from tareas where id_tarea = ?;")
	if err != nil {
		return &apierror.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "Data base error",
		}
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return &apierror.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "Error while deleting the task",
		}
	}

	defer stmt.Close()
	return nil
}

//CompleteTarea modifica una tarea pasando su estado a completado
func CompleteTarea(tarea *tareas.Tareas) *apierror.ApiError {
	stmt, err := db.Init().Prepare("update tareas set completado=? where id_tarea=?;")
	if err != nil {
		return &apierror.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "Data base error",
		}
	}

	_, err = stmt.Exec(tarea.Completado, tarea.IDTarea)
	if err != nil {
		return &apierror.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "Error while completing the task",
		}
	}

	defer stmt.Close()
	return nil
}

func SetTarea(user int64, idTarea int64, tareastr string, completado int8) *tareas.Tareas {
	var tarea tareas.Tareas
	tarea.IDUsuario = user
	tarea.IDTarea = idTarea
	tarea.Tarea = tareastr
	tarea.Completado = completado

	return &tarea
}
