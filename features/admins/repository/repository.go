package repository

import (
	"bayareen-backend/features/admins"

	"gorm.io/gorm"
)

type postgresUserRepository struct {
	Conn *gorm.DB
}

func NewPostgresUserRepository(conn *gorm.DB) admins.Data {
	return &postgresUserRepository{
		Conn: conn,
	}
}

func (repo *postgresUserRepository) Create(data *admins.Core) (*admins.Core, error) {
	record := FromCore(data)
	if err := repo.Conn.Create(&record).Error; err != nil {
		return &admins.Core{}, err
	}
	return record.ToCore(), nil
}
