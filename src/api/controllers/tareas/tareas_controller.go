package tareas

import (
	tareas "github.com/florgm/webplanner_api/src/api/services/tareas"
	"github.com/florgm/webplanner_api/src/api/services/sessions"
	"github.com/florgm/webplanner_api/src/api/utils/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetTareas devuelve todas las tareas del usuario logueado
func GetTareas(c *gin.Context) {
	user, apiErr := sessions.ValidateLoggedUser(c);
	if apiErr != nil {
		if apiErr != nil {
			c.JSON(apiErr.Status, apiErr.Message)
			return
		}
	}

	tareas, apiErr := tareas.GetTareas(user)
	if apiErr != nil {
		if apiErr != nil {
			c.JSON(apiErr.Status, apiErr.Message)
			return
		}
	}

	c.JSON(http.StatusOK, tareas)
}

//CreateTarea crea una nueva tarea para el usuario logueado
func CreateTarea(c *gin.Context) {
	user, apiErr := sessions.ValidateLoggedUser(c);
	if apiErr != nil {
		if apiErr != nil {
			c.JSON(apiErr.Status, apiErr.Message)
			return
		}
	}

	data, err := rest.GetJSONBody(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tr, err := tareas.ParseTarea(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	idTarea, apiErr := tareas.CreateTarea(tr, user); 
	if err != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	tarea := tareas.SetTarea(user, idTarea, tr.Tarea, tr.Completado)
	c.JSON(http.StatusOK, tarea)
}

//DeleteTarea borra una tarea especifica
func DeleteTarea(c *gin.Context) {
	_, apiErr := sessions.ValidateLoggedUser(c);
	if apiErr != nil {
		if apiErr != nil {
			c.JSON(apiErr.Status, apiErr.Message)
			return
		}
	}

	data, err := rest.GetJSONBody(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tarea, err := tareas.ParseTarea(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if apiErr := tareas.DeleteTarea(tarea); apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}

//CompleteTarea modifica una tarea pasando su estado a completado
func CompleteTarea(c *gin.Context) {
	_, apiErr := sessions.ValidateLoggedUser(c);
	if apiErr != nil {
		if apiErr != nil {
			c.JSON(apiErr.Status, apiErr.Message)
			return
		}
	}

	data, err := rest.GetJSONBody(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tarea, err := tareas.ParseTarea(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if apiErr := tareas.CompleteTarea(tarea); apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully modified"})
}
