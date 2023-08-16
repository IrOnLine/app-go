package main

import (
  "fmt"
  "os"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"

  "myapi/config"
  "myapi/middleware"
  "myapi/routes"
  "myapi/utils"
)

func main() {

  // Load environment variables
  if err := godotenv.Load(); err != nil {
    fmt.Println("Error loading .env file")
    os.Exit(1)
  }

  // Initialize logger
  if err := utils.InitLogger(); err != nil {
    fmt.Println("Error initializing logger")
    os.Exit(1)
  }

  // Connect to database
  db, err := config.ConnectDB()
  if err != nil {
    utils.Log.Fatal("Could not connect to database")
    os.Exit(1)
  }

  // Initialize router
  router := gin.Default()

  // Register routes
  routes.Register(router)

  // Register middleware
  router.Use(middleware.ErrorHandler())
  router.Use(middleware.Logger())

  // Start server
  utils.Log.Info("Starting API server on port 8000")
  if err := router.Run(":8000"); err != nil {
    utils.Log.Fatal("Could not start API server")
    os.Exit(1)
  }

}