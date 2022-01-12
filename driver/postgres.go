package driver

import (
	"bayareen-backend/config"
	_adminRepo "bayareen-backend/features/admins/repository"
	_categoryRepo "bayareen-backend/features/categories/repository"
	_paymentMethodRepo "bayareen-backend/features/paymentmethods/repository"
	_productRepo "bayareen-backend/features/products/repository"
	_providerRepo "bayareen-backend/features/providers/repository"
	_userRepo "bayareen-backend/features/user/repository"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MigrateDB() {
	DB.AutoMigrate(&_adminRepo.Admin{})
	DB.AutoMigrate(&_userRepo.User{})
	DB.AutoMigrate(&_providerRepo.Provider{})
	DB.AutoMigrate(&_categoryRepo.Category{})
	DB.AutoMigrate(&_paymentMethodRepo.PaymentMethod{})
	DB.AutoMigrate(&_productRepo.Product{})
}

func InitDB() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("err", err)
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	DB = db.Debug()
	MigrateDB()
}
