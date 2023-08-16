package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm" // Added GORM import

	"app/models"
)

// Global DB variable
var DB *gorm.DB

// GetUsers handles GET /users route
func GetUsers(c *gin.Context) {

	// Declare empty user array
	var users []models.User

	// Execute DB query
	result := DB.Find(&users)

	// Check for errors
	if result.Error != nil {

		// Return 500 error if query fails
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	// Return data if successful
	c.JSON(http.StatusOK, users)
}

// CreateUser handles POST /users route
func CreateUser(c *gin.Context) {

	// Declare input struct
	var input models.CreateUserInput

	// Bind JSON data to input
	if err := c.ShouldBindJSON(&input); err != nil {

		// Return 400 error if validation fails
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	if !validateUserInput(input) {

		// Return 400 error if validation fails
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create user object
	user := models.User{Name: input.Name, Email: input.Email}

	// Save user to database
	result := DB.Create(&user)

	// Check for db error
	if result.Error != nil {

		// Return 500 error
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return 201 response if created
	c.JSON(http.StatusCreated, user)

}

// Validate input function
func validateUserInput(input models.CreateUserInput) bool {

	// Input validation logic

	return true
}
