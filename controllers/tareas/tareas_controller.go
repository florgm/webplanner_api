package tareas

import (
	"fmt"
	"/github.com/gin-gonic/gin"
	"net/http"
	"../../utils"
	tareasService "../../services/tareas"
)

func GetTareas(c *gin.Context){
	tareas := tareasService.GetTareas()

	c.JSON(http.StatusOK,tareas)
}

func CreateTarea(c *gin.Context){
	data,err := utils.GetJSONBody(c.Request)
	if err != nil{
		fmt.Print(err)
		c.JSON(
			http.StatusBadRequest,
			err,
		)
		return
	}

	tarea, err := tareasService.ParseTarea(data)
	if err != nil{
		fmt.Print(err)
		c.JSON(
			http.StatusBadRequest,
			err,
		)
		return
	}
	 if err := tareasService.CreateTarea(tarea); err != nil{
	 	fmt.Println(err)
	 	c.JSON(
	 		http.StatusInternalServerError,
	 		err,
	 	)
	 	return
	 }
	 c.JSON(http.StatusOK,gin.H{
	 	"message": fmt.Sprintf("Successfully created"),
	 })
}

func DeleteTarea(c *gin.Context) {
	data, err := utils.GetJSONBody(c.Request)
	if err != nil {
		fmt.Println(err)
		c.JSON(
			http.StatusBadRequest,
			err,
		)
		return
	}

	tarea, err := tareasService.ParseTareaID(data)
	if err != nil {
		fmt.Println(err)
		c.JSON(
			http.StatusBadRequest,
			err,
		)
		return
	}

	if err := tareasService.DeleteTarea(tarea); err != nil {
		fmt.Println(err)
		c.JSON(
			http.StatusInternalServerError,
			err,
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted"),
	})
}

func CompleteTarea (c *gin.Context) {
	data, err := utils.GetJSONBody(c.Request)
	if err != nil {
		fmt.Println(err)
		c.JSON(
			http.StatusBadRequest,
			err,
		)
		return
	}

	tarea, err := tareasService.ParseTarea(data)
	if err != nil {
		fmt.Println(err)
		c.JSON(
			http.StatusBadRequest,
			err,
		)
		return
	}

	if err := tareasService.CompleteTarea(tarea); err != nil {
		fmt.Println(err)
		c.JSON(
			http.StatusInternalServerError,
			err,
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully modified"),
	})
}

