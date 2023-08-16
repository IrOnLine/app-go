package handlers

import (
  "net/http"

  "github.com/gin-gonic/gin"

  "myapi/models"
)

/// GetUsers fetches users from database
func GetUsers(c *gin.Context) {

	// Get users from database
	var users []User
	if result := db.Find(&users); result.Error != nil {
	  // Handle error
	  c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
	  log.Error(result.Error)
	  return
	}
  
	// Return response
	c.JSON(http.StatusOK, users)
  }

// CreateUser creates a new user in database
func CreateUser(c *gin.Context) {
  
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  // Handle invalid input error
	  c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	// Create user in database
	user := User{Name: input.Name, Email: input.Email}
	if result := db.Create(&user); result.Error != nil {
	  // Handle db error
	  c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"}) 
	  log.Error(result.Error)
	  return
	}
  
	// Return response
	c.JSON(http.StatusCreated, user)
  }

// GetUser fetches a single user
func GetUser(c *gin.Context) {
  // get user id from url param
  
  // fetch user from database
  
  // return response
}

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
  // get user id from url param
  
  // validate input 
  
  // update user in database
  
  // return response
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
  // get user id from url param
  
  // delete user from database
  
  // return response
}