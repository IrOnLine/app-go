package handlers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
)

// Global database handle 
var DB *gorm.DB

// SetDB sets the database handle
func SetDB(db *gorm.DB) {
  DB = db
}

// GetUsers handles GET /users route
func GetUsers(c *gin.Context) {

  var users []User
  result := DB.Find(&users)
  if result.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
    return
  }

  c.JSON(http.StatusOK, users)
}

// CreateUser handles POST /users route
func CreateUser(c *gin.Context) {

  // Bind input
  var input CreateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
    return
  }

  // Validate input
  if !validateUserInput(input) {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return 
  }

  // Create user in database
  user := User{Name: input.Name, Email: input.Email}
  result := DB.Create(&user)
  
  // Handle errors
  if result.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
    return
  }

  c.JSON(http.StatusCreated, user)
}

// Validate user input
func validateUserInput(input CreateUserInput) bool {
  // Input validation
  
  return true // or false
}