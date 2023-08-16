package utils

import (
  "errors"
  "time"

  "github.com/golang-jwt/jwt"
)

var (
  // ErrInvalidToken is returned when a token is invalid
  ErrInvalidToken = errors.New("invalid token provided")
  
  // ErrCreatingToken is returned when token generation fails
  ErrCreatingToken = errors.New("error creating JWT token")
)

// GenerateToken creates a new JWT token 
func GenerateToken(userId int) (string, error) {

  claims := jwt.StandardClaims{
    Subject:   strconv.Itoa(userId),
    ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  // Sign token
  tokenString, err := token.SignedString([]byte(mySecretKey))
  if err != nil {
    return "", ErrCreatingToken
  }

  return tokenString, nil

}

// ValidateToken checks if a JWT token is valid
func ValidateToken(tokenString string) error {

  // Parse token  
  token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte(mySecretKey), nil
  })

  // Check for parse errors
  if err != nil {
    return ErrInvalidToken
  }

  // Validate claims
  if !token.Valid {
    return ErrInvalidToken
  }

  return nil
}