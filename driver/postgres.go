package driver

import (
	"bayareen-backend/config"
	_providerRepo "bayareen-backend/features/providers/repository"
	_userRepo "bayareen-backend/features/user/repository"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MigrateDB() {
	DB.AutoMigrate(&_userRepo.User{})
	DB.AutoMigrate(&_providerRepo.Provider{})
}

func InitDB() {
	config, _ := config.LoadConfig(".")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	config.DBUser,
	// 	config.DBPass,
	// 	config.DBHost,
	// 	config.DBPort,
	// 	config.DBName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost,
		config.DBUser,
		config.DBPass,
		config.DBName,
		config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	DB = db.Debug()
	MigrateDB()
}
