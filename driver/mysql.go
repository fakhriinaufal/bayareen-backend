package driver

import (
	"bayareen-backend/config"
	_userRepo "bayareen-backend/features/user/repository"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MigrateDB() {
	DB.AutoMigrate(&_userRepo.User{})
}

func InitDB() {
	config, _ := config.LoadConfig(".")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	DB = db.Debug()
	MigrateDB()
}
