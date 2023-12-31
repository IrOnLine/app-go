package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
	"os"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dsn := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := os.Getenv("db_user") + ":" + os.Getenv("db_pass") + "@tcp(" + os.Getenv("db_host") + ":" + os.Getenv("db_port") + ")/" + os.Getenv("db_name") + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	conn, err := gorm.Open("mysql", dsn)

	//	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	//	fmt.Println(dbUri)
	//	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

func GetDB() *gorm.DB {
	return db
}
