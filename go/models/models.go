package models

import (
  "errors"

  "gorm.io/gorm"
)

// User model
type User struct {
  gorm.Model
  Name string
  Email string
}

// CreateUser creates a new user record
func CreateUser(db *gorm.DB, name, email string) (*User, error) {

  // Basic validation
  if name == "" {
    return nil, errors.New("name cannot be blank")
  }

  if email == "" {
    return nil, errors.New("email cannot be blank") 
  }

  user := &User{
    Name: name,
    Email: email,
  }

  result := db.Create(user)

  if result.Error != nil {
    return nil, result.Error
  }

  return user, nil
}

// GetUser fetches a single user
func GetUser(db *gorm.DB, id uint) (*User, error) {

  var user User
  result := db.First(&user, id)

  if result.Error != nil {
    // If user not found, return nicer error
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return nil, errors.New("user not found")
    }
    return nil, result.Error
  }

  return &user, nil  
}