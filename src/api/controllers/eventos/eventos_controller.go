package eventos

import (
    eventos "github.com/florgm/webplanner_api/src/api/services/eventos"
    "github.com/florgm/webplanner_api/src/api/services/sessions"
    "github.com/florgm/webplanner_api/src/api/utils/rest"
    "github.com/gin-gonic/gin"
    "net/http"
)

//GetEventos trae los eventos guardados del usuario y los feriados
func GetEventos(c *gin.Context) {
	user, apiErr := sessions.ValidateLoggedUser(c);
	if apiErr != nil {
		if apiErr != nil {
			c.JSON(apiErr.Status, apiErr.Message)
			return
		}
	}

	eventos, apiErr := eventos.GetEventos(user)
	if apiErr != nil {
		c.JSON(http.StatusInternalServerError, eventos)	
	}
	c.JSON(http.StatusOK, eventos)
}

//CreateEvento crea un evento en la base de datos
func CreateEvento(c *gin.Context) {
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
		
	evento, err := eventos.ParseEvento(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
		
	if apiErr := eventos.CreateEvento(evento, user); err != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created"})
}

//DeleteEvento elmina un evento de la base de datos
func DeleteEvento(c *gin.Context) {
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

	evento, err := eventos.ParseEvento(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if apiErr := eventos.DeleteEvento(evento); apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}

//UpdateEvento modifica un evento de la base de datos
func UpdateEvento(c *gin.Context) {
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

	evento, err := eventos.ParseEvento(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if apiErr := eventos.ModifyEvento(evento); apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully modified"})
}
