package eventos

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"../../utils"
	"github.com/gin-contrib/sessions"
	eventosService "../../services/eventos"
)

const (
	userkey = 123
)

//GetEventos esto es una función 
func GetEventos(c *gin.Context) {
	session := sessions.Default(c)
	var user int64
	user = session.Get(userkey).(int64)

	eventos := eventosService.GetEventos(user)
	c.JSON(http.StatusOK, eventos)
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
func UpdateEvento (c *gin.Context) {
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