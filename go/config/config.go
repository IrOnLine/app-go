package config

import (
  "fmt"
  "os"

  "github.com/joho/godotenv"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

// Load loads environment variables and validates required ones exist
func Load() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	// Validate required environment variables
	requiredVars := []string{"DB_HOST", "DB_USER", "DB_PASSWORD"}
	
	for _, v := range requiredVars {
	  if os.Getenv(v) == "" {
		log.Fatal("Required environment variable not set: " + v) 
	  }
	}
  }
  
  // Database connects to DB and validates connection
  func Database() (*gorm.DB, error) {
  
	// Build DSN
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
	  os.Getenv("DB_USER"), 
	  os.Getenv("DB_PASSWORD"),
	  os.Getenv("DB_HOST"), 
	  os.Getenv("DB_PORT"),
	  os.Getenv("DB_NAME"))
  
	// Open database connection
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
	  return nil, err 
	}
  
	// Validate database connection
	if err = db.Ping(); err != nil {
	  return nil, errors.New("Failed to connect to database")
	}
  
	return db, nil
  }