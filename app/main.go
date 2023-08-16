package main

import (
	"gin.com/gin/controllers"
	"gin.com/gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/tasks", controllers.FindTasks)
	r.POST("/api/tasks", controllers.CreateTask)
	r.GET("/api/tasks/one", controllers.FindTask)
	r.PUT("/api/tasks/update", controllers.UpdateTask)
	r.Run()
}
