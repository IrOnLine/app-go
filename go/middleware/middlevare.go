package middleware

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

// Auth checks for valid JWT token
func Auth() gin.HandlerFunc {
  return func(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    
    if token == "" {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No auth token"})
      return
    }

    err := validateToken(token)
    if err != nil {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
      return
    }

    c.Next()
  }
}

// ErrorHandler catches and handles errors
func ErrorHandler() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Next()
    
    err := c.Errors.Last()
    if err != nil {
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
  }
}

// validateToken verifies JWT token
func validateToken(token string) error {
  // JWT parsing and validation
  
  if !valid {
    return errors.New("Invalid token")
  }

  return nil
}