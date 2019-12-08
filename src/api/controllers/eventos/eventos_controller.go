package eventos

import (
    eventos "github.com/florgm/webplanner_api/src/api/services/eventos"
    "github.com/florgm/webplanner_api/src/api/services/sessions"
    "github.com/florgm/webplanner_api/src/api/utils/rest"
    "github.com/gin-gonic/gin"
    "net/http"
)

//GetEventos esto es una función
func GetEventos(c *gin.Context) {
	if user := sessions.ValidateLoggedUser(c); user > 0 {
		eventos := eventos.GetEventos(user)
		c.JSON(http.StatusOK, eventos)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error with the get of context"})
}

//CreateEvento esto es una función
func CreateEvento(c *gin.Context) {
	if user := sessions.ValidateLoggedUser(c); user > 0 {
		data, err := rest.GetJSONBody(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest,err)
			return
		}
		
		evento, err := eventos.ParseEvento(data)
		if err != nil {
			c.JSON(http.StatusBadRequest,err)
			return
		}
		
		if err := eventos.CreateEvento(evento, user); err != nil {
			c.JSON(http.StatusInternalServerError,err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully created"})
		return
	}
    c.JSON(http.StatusInternalServerError, gin.H{"message": "Error with the get of context"})
}

//DeleteEvento esto es una funcion
func DeleteEvento(c *gin.Context) {
	if user := sessions.ValidateLoggedUser(c); user > 0 { 
		data, err := rest.GetJSONBody(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest,err)
			return
		}

		evento, err := eventos.ParseEvento(data)
		if err != nil {
			c.JSON(http.StatusBadRequest,err)
			return
		}

		if err := eventos.DeleteEvento(evento); err != nil {
			c.JSON(http.StatusInternalServerError,err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error with the get of context"})
}

//UpdateEvento esto es una funcion
func UpdateEvento(c *gin.Context) {
	if user := sessions.ValidateLoggedUser(c); user > 0 { 
		data, err := rest.GetJSONBody(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest,err)
			return
		}

		evento, err := eventos.ParseEvento(data)
		if err != nil {
			c.JSON(http.StatusBadRequest,err)
			return
		}

		if err := eventos.ModifyEvento(evento); err != nil {
			c.JSON(http.StatusInternalServerError,err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully modified"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error with the get of context"})
}
