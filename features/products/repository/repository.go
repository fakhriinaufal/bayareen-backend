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
	repo.Conn.Where("status = ?", true).Find(&records)
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
	err := repo.Conn.Model(record).Updates(map[string]interface{}{"name": record.Name, "price": record.Price, "status": record.Status}).Error
	if err != nil {
		return &products.Core{}, err
	}
	return record.ToCore(), nil
}

func (repo *postgresProductRepository) Delete(id []int) error {
	return repo.Conn.Delete(&Product{}, id).Error
}

func (repo *postgresProductRepository) GetByProviderId(provId int) ([]products.Core, error) {
	var records []Product
	err := repo.Conn.Where("provider_id = ? AND status = ?", provId, true).Find(&records).Error
	if err != nil {
		return nil, err
	}

	return ToCoreSlice(records), nil
}
