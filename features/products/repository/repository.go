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

func (repo *postgresProductRepository) GetAll() []products.Core {
	records := []Product{}
	repo.Conn.Find(&records)
	return ToCoreSlice(records)
}

func (repo *postgresProductRepository) GetById(id int) (*products.Core, error) {
	record := Product{Id: id}
	if err := repo.Conn.First(&record).Error; err != nil {
		return &products.Core{}, err
	}
	return record.ToCore(), nil
}

func (repo *postgresProductRepository) Update(data *products.Core) (*products.Core, error) {
	record := FromCore(data)
	if err := repo.Conn.Save(record).Error; err != nil {
		return &products.Core{}, err
	}
	return record.ToCore(), nil
}

func (repo *postgresProductRepository) Delete(id int) error {
	return repo.Conn.Delete(&Product{Id: id}).Error
}
