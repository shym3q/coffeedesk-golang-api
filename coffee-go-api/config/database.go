package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"fmt"
)

var DB *gorm.DB
var err error

func GetDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(192.168.1.22:3306)/coffee?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DBUSER"), os.Getenv("DBPASS"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
}
