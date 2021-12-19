package data

import (
	"bayareen-backend/features/user"

	"gorm.io/gorm"
)

type mySqlRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) user.Data {
	return &mySqlRepository{
		Conn: conn,
	}
}

func (repo *mySqlRepository) Create(data user.UserCore) (user.UserCore, error) {
	record := fromCore(data)

	if err := repo.Conn.Create(&record).Error; err != nil {
		return user.UserCore{}, err
	}

	return record.toCore(), nil
}

func (repo *mySqlRepository) GetAll() []user.UserCore {
	records := []User{}

	repo.Conn.Find(&records)

	return toCoreSlice(records)
}
