package sessions

import (
	"github.com/florgm/webplanner_api/src/api/utils/apierror"
	db "github.com/florgm/webplanner_api/src/api/db"
	"github.com/gin-gonic/gin"
    "net/http"
)

//SaveSession funcion que guarda la sesion cuando el usuario se loguea
func SaveSession(token string, idUser int64) *apierror.ApiError {
	stmt, err := db.Init().Prepare("insert into sessions (token, user) values(?,?);")
	
	if err != nil {
        return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
	}
	
	_, err = stmt.Exec(token, idUser)

	if err != nil {
		return &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
	}
	
	defer stmt.Close()
	return nil
}

//GetSession funcion que valida que el usuario este logueado en la pagina
func GetSession(token string) (int64, *apierror.ApiError) {
	var idUsuario int64
    stmt, err := db.Init().Prepare("select user from sessions where token = ?;")

    if err != nil {
        return 0, &apierror.ApiError {
			Status: http.StatusInternalServerError,
			Message: "Data base error",
		}
    }

	result := stmt.QueryRow(token)
	err = result.Scan(&idUsuario)
	
	if err != nil {
		return 0, &apierror.ApiError {
			Status: http.StatusUnauthorized,
			Message: "User not logged in",
		}
	}
	
	defer stmt.Close()
	return idUsuario, nil	
}

//ValidateLoggedUser funcion
func ValidateLoggedUser(c *gin.Context) (int64, *apierror.ApiError) {
	usuario, exists := c.Get("idUsuario")
	if (exists) {
		return usuario.(int64), nil
	}

	return 0, &apierror.ApiError {
		Status: http.StatusInternalServerError,
		Message: "Error with the get context",
	}
}