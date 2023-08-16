package config

import (
  "fmt"
  "os"

  "github.com/joho/godotenv"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

// Load loads environment variables and validators
func Load() error {

  if err := godotenv.Load(); err != nil {
    return fmt.Errorf("error loading .env file: %v", err)
  }

  requiredVars := []string{"DB_HOST", "DB_USER", "DB_PASS"}
  
  for _, v := range requiredVars {
    if os.Getenv(v) == "" {
      return fmt.Errorf("required env var not defined: %s", v)
    }
  }

  return nil
}

// ConnectDB creates a database connection
func ConnectDB() (*gorm.DB, error) {
  
  dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASS"),
    os.Getenv("DB_HOST"),
    os.Getenv("DB_NAME"),
  ) 
   
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, fmt.Errorf("failed to connect to db: %v", err)
  }

  sqldb, err := db.DB()
  if err != nil {
    return nil, fmt.Errorf("failed to get raw database handle: %v", err) 
  }

  if err := sqldb.Ping(); err != nil {
    return nil, fmt.Errorf("failed to ping database: %v", err)
  }

  return db, nil

}