package eventos

import (
    "fmt"
    eventosService "github.com/florgm/webplanner_api/src/api/services/eventos"
    // sessions "github.com/florgm/webplanner_api/src/api/services/sessions"
    "github.com/florgm/webplanner_api/src/api/utils/rest"
    "github.com/gin-gonic/gin"
    "net/http"
)

//GetEventos esto es una función
func GetEventos(c *gin.Context) {
	usuario, exists := c.Get("idUsuario")
	if (exists) {
		fmt.Println("ACA ENTRA")
		eventos := eventosService.GetEventos(usuario)
		fmt.Println(eventos)
		c.JSON(http.StatusOK, eventos)
	}
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
