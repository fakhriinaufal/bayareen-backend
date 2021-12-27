package repository

import (
	"bayareen-backend/features/products"

	"gorm.io/gorm"
)

type postgresProductRepository struct {
	Conn *gorm.DB
}

func NewPostgresProductRepository(conn *gorm.DB) products.Data {
	return &postgresProductRepository{
		Conn: conn,
	}
}

func (repo *postgresProductRepository) Create(data *products.Core) (*products.Core, error) {
	record := FromCore(data)
	if err := repo.Conn.Create(record).Error; err != nil {
		return &products.Core{}, err
	}
	return record.ToCore(), nil
}
