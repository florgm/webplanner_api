package main

import (
	"fmt"
	"net/http"
	// "log"
  	"github.com/gin-gonic/gin"
  	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
  	eventosControllers "./controllers/eventos"
	tareasControllers "./controllers/tareas"
	usuariosControllers "./controllers/usuarios"
)

const (
	userkey = 123
)

func main() {
  router := SetupRouter()
  router.Use(gin.Logger())
  if err:= SetupRouter().Run(":8081"); err != nil {
		fmt.Println("Unable to start: ", err)
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

	router.POST("/login", usuariosControllers.Login(userkey))
	router.GET("/logout", usuariosControllers.Logout(userkey))

	auth := router.Group("/auth")
	auth.Use(AuthRequired) 
	{
		router.GET("/eventos", eventosControllers.GetEventos)
		router.POST("/eventos", eventosControllers.CreateEvento)
		router.DELETE("/eventos", eventosControllers.DeleteEvento)
		router.PUT("/eventos", eventosControllers.UpdateEvento)

		router.GET("/tareas", tareasControllers.GetTareas)
		router.POST("/tareas", tareasControllers.CreateTarea)
		router.DELETE("/tareas", tareasControllers.DeleteTarea)
		router.PUT("/tareas", tareasControllers.CompleteTarea)
	}
	
  	return router
}

//AuthRequired funcion para comprobar que la sesion esta iniciada
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}