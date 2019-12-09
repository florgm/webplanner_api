package app

import (
	"fmt"
	"net/http"

	"github.com/florgm/webplanner_api/src/api/controllers/eventos"
	"github.com/florgm/webplanner_api/src/api/controllers/tareas"
	"github.com/florgm/webplanner_api/src/api/controllers/usuarios"
	"github.com/florgm/webplanner_api/src/api/services/sessions"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = SetupRouter()
)

func Run() {
	if err := router.Run(":80"); err != nil {
		fmt.Println(err.Error())
	}
}

func SetupRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/login", usuarios.Login)
	router.POST("/usuarios", usuarios.CreateUsuario)

	auth := router.Group("/auth")
	auth.Use(AuthRequired)
	{
		auth.GET("/eventos", eventos.GetEventos)
		auth.POST("/eventos", eventos.CreateEvento)
		auth.DELETE("/eventos", eventos.DeleteEvento)
		auth.PUT("/eventos", eventos.UpdateEvento)

		auth.GET("/tareas", tareas.GetTareas)
		auth.POST("/tareas", tareas.CreateTarea)
		auth.DELETE("/tareas", tareas.DeleteTarea)
		auth.PUT("/tareas", tareas.CompleteTarea)

		auth.GET("/logout", usuarios.Logout)
		auth.PUT("/usuarios", usuarios.UpdateUsuario)
	}

	return router
}

func AuthRequired(c *gin.Context) {
	if token := c.Request.Header.Get("Authorization"); len(token) > 0 {
		idUsuario, apiErr := sessions.GetSession(token)
		if apiErr != nil {
			c.JSON(apiErr.Status, apiErr.Message)
			return
		}
		c.Set("idUsuario", idUsuario)
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, "Invalid token")
}
