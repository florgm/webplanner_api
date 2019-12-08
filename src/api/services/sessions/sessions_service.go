package sessions

import (
	db "github.com/florgm/webplanner_api/src/api/db"
	"fmt"
	"github.com/florgm/webplanner_api/src/api/utils/apierror"
	"github.com/gin-gonic/gin"
    "net/http"
)

//SaveSession funcion
func SaveSession(token string, idUser int64) error {
	stmt, err := db.Init().Prepare("insert into sessions (token, user) values(?,?);")
	if err != nil {
        fmt.Print(err.Error())
	}
	
	_, err = stmt.Exec(token, idUser)
	defer stmt.Close()
	return err
}

//GetSession funcion
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
			Message: "Invalid Token",
		}
	}
	
	defer stmt.Close()
	return idUsuario, nil	
}

//ValidateLoggedUser funcion
func ValidateLoggedUser(c *gin.Context) int64 {
	usuario, exists := c.Get("idUsuario")
	if (exists) {
		return usuario.(int64)
	}
	return 0
}