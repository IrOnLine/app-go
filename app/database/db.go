package database

import (
	"app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//const DB_USER = "root"
//const DB_PASSWORD = ""
//const DB_NAME = "my_db"
//const DB_HOST = "127.0.0.1"
//const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}
