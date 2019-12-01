package eventos

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"../../utils"
	eventosService "../../services/eventos"
)

//GetEventos esto es una función 
func GetEventos(c *gin.Context) {
    
    eventos := eventosService.GetEventos()

	c.JSON(http.StatusOK, gin.H{
		"result": eventos,
	})
}

//CreateEvento esto es una función
func CreateEvento(c *gin.Context) {
	data, err := utils.GetJSONBody(c.Request)
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
	data, err := utils.GetJSONBody(c.Request)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
	}

    evento, err := eventosService.ParseEventoID(data)
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

//ModifyEvento esto es una funcion
func ModifyEvento (c *gin.Context) {
	data, err := utils.GetJSONBody(c.Request)
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