package main

import (
  "fmt"

  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "myapi/config"
  "myapi/routes"
)

// Initialize logger
var log = logrus.New()

func main() {

	// Load config
	config.Load()
	
	// Initialize database
	db, err := config.Database()
	if err != nil {
	  log.Fatal(err)
	}
  
	// Initialize router
	router := gin.Default()
  
	// Register routes
	routes.Register(router)
  
	// Start server
	log.Info("Starting server on port 8000")
	if err := router.Run(":8000"); err != nil {
	  log.Fatal(err)
	}
  }