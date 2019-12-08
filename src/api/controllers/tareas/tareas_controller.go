package tareas

import (
	"net/http"
	"github.com/florgm/webplanner_api/src/api/services/sessions"
	tareas "github.com/florgm/webplanner_api/src/api/services/tareas"
	"github.com/florgm/webplanner_api/src/api/utils/rest"
	"github.com/gin-gonic/gin"
)

//GetTareas esto es una funcion
func GetTareas(c *gin.Context) {
	if user := sessions.ValidateLoggedUser(c); user > 0 {
		tareas := tareas.GetTareas(user)
		c.JSON(http.StatusOK, tareas)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error with the get of context"})
}

//CreateTarea esto es una funcion
func CreateTarea(c *gin.Context) {
	if user := sessions.ValidateLoggedUser(c); user > 0 {
		data, err := rest.GetJSONBody(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		tr, err := tareas.ParseTarea(data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		idTarea, err := tareas.CreateTarea(tr, user); 
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		tarea := tareas.SetTarea(user, idTarea, tr.Tarea, tr.Completado)
		c.JSON(http.StatusOK, tarea)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error with the get of context"})
}

//DeleteTarea esto es una funcion
func DeleteTarea(c *gin.Context) {
	if user := sessions.ValidateLoggedUser(c); user > 0 {
		data, err := rest.GetJSONBody(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		tarea, err := tareas.ParseTarea(data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := tareas.DeleteTarea(tarea); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error with the get of context"})
}

//CompleteTarea esto es una funcion
func CompleteTarea(c *gin.Context) {
	if user := sessions.ValidateLoggedUser(c); user > 0 {
		data, err := rest.GetJSONBody(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		tarea, err := tareas.ParseTarea(data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := tareas.CompleteTarea(tarea); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully modified"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error with the get of context"})
}
