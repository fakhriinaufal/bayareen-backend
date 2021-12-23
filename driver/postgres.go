package driver

import (
	"bayareen-backend/config"
	_categoryRepo "bayareen-backend/features/categories/repository"
	_paymentMethodRepo "bayareen-backend/features/paymentmethods/repository"
	_userRepo "bayareen-backend/features/user/repository"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MigrateDB() {
	DB.AutoMigrate(&_userRepo.User{})
	DB.AutoMigrate(&_categoryRepo.Category{})
	DB.AutoMigrate(&_paymentMethodRepo.PaymentMethod{})
}

func InitDB() {
	config, _ := config.LoadConfig(".")

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
