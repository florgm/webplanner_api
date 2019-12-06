package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  eventosControllers "./controllers/eventos"
  tareasControllers "./controllers/tareas"
)

func main() {
  router := SetupRouter()
  router.Run(":8081")
}

//SetupRouter esto es una funcion
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8888"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/eventos", eventosControllers.GetEventos)
	router.POST("/eventos", eventosControllers.CreateEvento)
	router.DELETE("/eventos", eventosControllers.DeleteEvento)
	router.PUT("/eventos", eventosControllers.UpdateEvento)

	router.GET("/tareas", tareasControllers.GetTareas)
	router.POST("/tareas", tareasControllers.CreateTarea)
	router.DELETE("/tareas", tareasControllers.DeleteTarea)
	router.PUT("/tareas", tareasControllers.CompleteTarea)

//   v1 := router.Group("api/v1")
//   {
//     v1.POST("/flower", controllers.Create)
//     v1.GET("/flower/:id", controllers.GetFlower)
//     v1.GET("/flowers", controllers.GetAllFlower)
//     v1.DELETE("/flower", controllers.DeleteFlower)
//     v1.GET("/check", controllers.HealthCheck)
//   }
  	return router
}