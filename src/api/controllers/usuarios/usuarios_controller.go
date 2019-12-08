package usuarios

import (
	"github.com/florgm/webplanner_api/src/api/services/sessions"
    "github.com/florgm/webplanner_api/src/api/services/usuarios"
	"github.com/florgm/webplanner_api/src/api/utils/rest"
	"github.com/gin-gonic/gin"
	"crypto/rand"
	"net/http"
    "fmt"
)

//Login loguea al usuario a la pagina
func Login(c *gin.Context) {
    data, err := rest.GetJSONBody(c.Request)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    usuario, err := usuarios.ParseLoginUsuario(data)
    if err != nil {
        c.JSON(http.StatusBadRequest,gin.H{"message": err.Error()})
        return
    }

    result, apiErr := usuarios.Login(usuario)

    if apiErr != nil {
        c.JSON(apiErr.Status, apiErr.Message)
        return
	}
	
	sessionToken := tokenGenerator()
	
	apiErr = sessions.SaveSession(sessionToken, result.IDUsuario)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

    c.JSON(http.StatusOK, sessionToken)
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

//Logout cierra la sesion del usuario
func Logout(c *gin.Context) {
	user, apiErr := sessions.ValidateLoggedUser(c); 
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	if apiErr := usuarios.Logout(user); apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

//CreateUsuario permite crear un usuario
func CreateUsuario (c *gin.Context) {
	data, err := rest.GetJSONBody(c.Request)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    usuario, err := usuarios.ParseUsuario(data)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
		
	if apiErr := usuarios.CreateUsuario(usuario); apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created"})
}

//UpdateUsuario permite modificar el nombre y la password del usuario
func UpdateUsuario (c *gin.Context) {
	user, apiErr := sessions.ValidateLoggedUser(c);
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	data, err := rest.GetJSONBody(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	usuario, err := usuarios.ParseUsuario(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if apiErr := usuarios.UpdateUsuario(user, usuario); apiErr != nil {
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully modified"})
}
