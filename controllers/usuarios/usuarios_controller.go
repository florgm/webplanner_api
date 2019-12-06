package usuarios

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"../../utils"
	usuariosService "../../services/usuarios"
)

// const (
// 	userkey = 123
// )

//Login funcion para logguear al usuario a la pagina web
func Login (c *gin.Context, userkey int64) {
	session := sessions.Default(c)

	data, err := utils.GetJSONBody(c.Request)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
	}

	usuario, err := usuariosService.ParseLoginUsuario(data)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
	}

	result, err := usuariosService.Login(usuario)

	if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

	session.Set(userkey, result.IDUsuario)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

//Logout funcion que cierra la sesion del usuario
func Logout (c *gin.Context, userkey int64) {
	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

//CreateUsuario


//UpdateUsuario