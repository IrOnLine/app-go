package models

import (
	"errors"

	"gorm.io/gorm"
)

// User model struct
type User struct {
	gorm.Model
	Name  string
	Email string
}

// CreateUser creates a new user record
func CreateUser(db *gorm.DB, name, email string) (*User, error) {

	// Validate input
	if name == "" {
		return nil, errors.New("Name cannot be blank")
	}

	if email == "" {
		return nil, errors.New("Email cannot be blank")
	}

	// Create user
	user := &User{
		Name:  name,
		Email: email,
	}

	// Save to DB
	result := db.Create(&user)

	// Check for errors
	if result.Error != nil {
		return nil, result.Error
	}

	// Return user
	return user, nil

}

// GetUser returns a user by ID
func GetUser(db *gorm.DB, id uint) (*User, error) {

	// Declare user var
	var user User

	// Query for user
	result := db.First(&user, id)

	// Check for errors
	if result.Error != nil {

		// Map db not found error to custom error
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}

		// Otherwise return error
		return nil, result.Error
	}

	// If found, return user
	return &user, nil

}
