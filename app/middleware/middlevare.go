package middleware

import (
	"errors" // Added import
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth middleware
func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Get token from request
		token := c.GetHeader("Authorization")

		// Check if token empty
		if token == "" {

			// Return 401 error
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No auth token"})
			return
		}

		// Validate token
		err := validateToken(token)

		// Check for errors
		if err != nil {

			// Return 401 error
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// If valid, proceed
		c.Next()
	}
}

// validateToken checks if token is valid
func validateToken(token string) error {

	// Parse and validate token

	if !valid {

		// Return structured error
		return errors.New("Invalid token") // Updated
	}

	return nil

}

// ErrorHandler to catch errors
func ErrorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Execute request
		c.Next()

		// Check for errors
		if err := c.Errors.Last(); err != nil {

			// Return 500 error
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}

}
