package routes

import (
  "github.com/gin-gonic/gin"
  
  "myapi/handlers"
)

// Register registers routes to router
func Register(router *gin.Engine) {
  api := router.Group("/api/v1")
  {
    api.GET("/users", handlers.GetUsers)
    api.POST("/users", handlers.CreateUser)
    api.GET("/users/:id", handlers.GetUser)
    api.PUT("/users/:id", handlers.UpdateUser)
    api.DELETE("/users/:id", handlers.DeleteUser)
  }
}