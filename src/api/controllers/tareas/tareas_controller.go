package tareas

import (
    "fmt"
    tareas2 "github.com/florgm/webplanner_api/src/api/services/tareas"
    "github.com/florgm/webplanner_api/src/api/utils/rest"
    "github.com/gin-gonic/gin"
    "net/http"
)

//GetTareas esto es una funcion
func GetTareas(c *gin.Context) {
    tareas := tareas2.GetTareas()

    c.JSON(http.StatusOK, tareas)
}

//CreateTarea esto es una funcion
func CreateTarea(c *gin.Context) {
    data, err := rest.GetJSONBody(c.Request)
    if err != nil {
        fmt.Print(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    tarea, err := tareas2.ParseTarea(data)
    if err != nil {
        fmt.Print(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    if err := tareas2.CreateTarea(tarea); err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusInternalServerError,
            err,
        )
        return
    }

    c.JSON(http.StatusOK, tarea)
}

//DeleteTarea esto es una funcion
func DeleteTarea(c *gin.Context) {
    data, err := rest.GetJSONBody(c.Request)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    tarea, err := tareas2.ParseTarea(data)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    if err := tareas2.DeleteTarea(tarea); err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusInternalServerError,
            err,
        )
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Successfully deleted"),
    })
}

//CompleteTarea esto es una funcion
func CompleteTarea(c *gin.Context) {
    data, err := rest.GetJSONBody(c.Request)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    tarea, err := tareas2.ParseTarea(data)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    if err := tareas2.CompleteTarea(tarea); err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusInternalServerError,
            err,
        )
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Successfully modified"),
    })
}
