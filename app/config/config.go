package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//const (
//	DB_USER     = "irol_user"
//	DB_PASSWORD = "irol_pw"
//	DB_NAME     = "irol"
//	DB_HOST     = "127.0.0.1"
//	DB_PORT     = "3306"
//	APP_PORT    = "500"
//	SALT        = "5CAF60FB9CB35CE169B76E657AB21FC4D1D6B093603"
//	JWT_SECRET  = "35CE169B76E657AB21FC4D1D6B093603"
//)

// Load loads environment variables and validates required ones exist
func Load() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Validate required environment variables
	requiredVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME", "DB_ROOT_PASSWORD", "APP_PORT"}

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatal("Required environment variable not set: " + v)
		}
	}
}

// Database connects to DB and validates connection
//func Database() (*gorm.DB, error) {
// Build DSN
//	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

// Open database connection
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
// Validate database connection
//	if err = db.Ping(); err != nil {
//		return nil, errors.New("Failed to connect to database")
//	}
//	return db, nil
//}

//func connectDB() *gorm.DB {
//	var err error
//	dsn := DB_USER + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
//	fmt.Println("dsn : ", dsn)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Println("Error connecting to database : error=%v", err)
//		return nil
//	}
//	return db
//}

// Find .env file
//	err := godotenv.Load("../.env")
//	if err != nil {
//		log.Fatalf("Error loading .env file: %s", err)
//	}
// Getting and using a value from .env
//	DB_USER := os.Getenv("DB_USER")
