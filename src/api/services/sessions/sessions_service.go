package sessions

import (
    "github.com/florgm/webplanner_api/src/api/utils/apierror"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "net/http"
)

const (
    userKey = 123
)

func ValidateLoggedUser(c *gin.Context) (int64, *apierror.ApiError) {
    session := sessions.Default(c)
    user := session.Get(userKey)
    if user == nil {
        return 0, &apierror.ApiError{
            Status:  http.StatusUnauthorized,
            Message: "unauthorized",
        }
    }
    return user.(int64), nil
}

func GetLoggedUser(c *gin.Context) int64 {
    return sessions.Default(c).Get(userKey).(int64)
}

func SetLoggedUser(c *gin.Context, userID int64) error {
    session := sessions.Default(c)
    session.Set(userKey, userID)
    return session.Save()
}

func Logout(c *gin.Context) error {
    session := sessions.Default(c)
    session.Delete(userKey)
    return session.Save()
}
