package main

import (
  "fmt"

  "github.com/gin-gonic/gin"
  eventosControllers "./controllers/eventos"
)

func main() {
  fmt.Println("hello world")

  router := SetupRouter()
  router.Run(":8081")
}


func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/eventos", eventosControllers.GetEventos)
	router.POST("/eventos", eventosControllers.CreateEvento)
	router.DELETE("/eventos", eventosControllers.DeleteEvento)
	router.PUT("/eventos", eventosControllers.ModifyEvento)
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