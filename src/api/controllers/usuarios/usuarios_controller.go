package usuarios

import (
	"crypto/rand"
    "fmt"
    "github.com/florgm/webplanner_api/src/api/services/sessions"
    "github.com/florgm/webplanner_api/src/api/services/usuarios"
    "github.com/florgm/webplanner_api/src/api/utils/rest"
    "github.com/gin-gonic/gin"
    "net/http"
)

//Login funcion para logguear al usuario a la pagina web
func Login(c *gin.Context) {
    data, err := rest.GetJSONBody(c.Request)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    usuario, err := usuarios.ParseLoginUsuario(data)
    if err != nil {
        fmt.Println(err)
        c.JSON(
            http.StatusBadRequest,
            err,
        )
        return
    }

    result, err := usuarios.Login(usuario)

    if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
	}
	
	sessionToken := tokenGenerator()
	err = sessions.SaveSession(sessionToken, result.IDUsuario)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

    c.JSON(http.StatusOK, sessionToken)
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

//Logout funcion que cierra la sesion del usuario
func Logout(c *gin.Context) {
    // if err := sessions.Logout(c); err != nil {
    //     c.JSON(http.StatusInternalServerError, err.Error())
    //     return
    // }

    // c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

//CreateUsuario

//UpdateUsuario
