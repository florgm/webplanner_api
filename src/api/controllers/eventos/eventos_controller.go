package eventos

import (
    "fmt"
    eventosService "github.com/florgm/webplanner_api/src/api/services/eventos"
    sessions2 "github.com/florgm/webplanner_api/src/api/services/sessions"
    "github.com/florgm/webplanner_api/src/api/utils/rest"
    "github.com/gin-gonic/gin"
    "net/http"
)

//GetEventos esto es una función
func GetEventos(c *gin.Context) {
    eventos := eventosService.GetEventos(
        sessions2.GetLoggedUser(c),
    )
    c.JSON(http.StatusOK, eventos)
}

//CreateEvento esto es una función
func CreateEvento(c *gin.Context) {
    data, err := rest.GetJSONBody(c.Request)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    evento, err := eventosService.ParseEvento(data)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    if err := eventosService.CreateEvento(evento); err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusInternalServerError,
            err,
        )
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Successfully created"),
    })
}

//DeleteEvento esto es una funcion
func DeleteEvento(c *gin.Context) {
    data, err := rest.GetJSONBody(c.Request)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    evento, err := eventosService.ParseEvento(data)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    if err := eventosService.DeleteEvento(evento); err != nil {
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

//UpdateEvento esto es una funcion
func UpdateEvento(c *gin.Context) {
    data, err := rest.GetJSONBody(c.Request)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    evento, err := eventosService.ParseEvento(data)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    if err := eventosService.ModifyEvento(evento); err != nil {
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
