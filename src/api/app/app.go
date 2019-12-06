package app

import (
    "github.com/florgm/webplanner_api/src/api/controllers/eventos"
    "github.com/florgm/webplanner_api/src/api/controllers/tareas"
    "github.com/florgm/webplanner_api/src/api/controllers/usuarios"
    sessionService "github.com/florgm/webplanner_api/src/api/services/sessions"
    "github.com/florgm/webplanner_api/src/api/utils/logger"
    "github.com/gin-contrib/cors"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
)

var (
    router = SetupRouter()
)

func Run() {
    if err := router.Run(":8081"); err != nil {
        logger.Error("starting router", err);
    }
}

//SetupRouter esto es una funcion
func SetupRouter() *gin.Engine {
    router := gin.New()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8888"},
        AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    store := cookie.NewStore([]byte("secret"))
    router.Use(sessions.Sessions("webplanner", store))

    router.POST("/login", usuarios.Login)
    router.GET("/logout", usuarios.Logout)

    auth := router.Group("/auth")
    auth.Use(AuthRequired)
    {
        router.GET("/eventos", eventos.GetEventos)
        router.POST("/eventos", eventos.CreateEvento)
        router.DELETE("/eventos", eventos.DeleteEvento)
        router.PUT("/eventos", eventos.UpdateEvento)

        router.GET("/tareas", tareas.GetTareas)
        router.POST("/tareas", tareas.CreateTarea)
        router.DELETE("/tareas", tareas.DeleteTarea)
        router.PUT("/tareas", tareas.CompleteTarea)
    }

    return router
}

//AuthRequired funcion para comprobar que la sesion esta iniciada
func AuthRequired(c *gin.Context) {
    _, apiErr := sessionService.ValidateLoggedUser(c)
    if apiErr != nil {
        c.JSON(apiErr.Status, apiErr.Message)
        return
    }
    c.Next()
}
