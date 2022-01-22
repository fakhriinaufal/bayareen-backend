package repository

import (
	"bayareen-backend/features/providers"

	"gorm.io/gorm"
)

type posgresRepository struct {
	Conn *gorm.DB
}

func NewPostgresRepository(conn *gorm.DB) providers.Data {
	return &posgresRepository{
		Conn: conn,
	}
}

func (repo *posgresRepository) Create(data *providers.Core) (*providers.Core, error) {
	record := FromCore(data)

	if err := repo.Conn.Create(record).Error; err != nil {
		return &providers.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *posgresRepository) GetByCategoryId(catId int) ([]providers.Core, error) {
	records := []Provider{}
	if err := repo.Conn.Where("cat_id = ?", catId).Find(&records).Error; err != nil {
		return nil, err
	}

	return ToCoreSlice(records), nil
}

func (repo *posgresRepository) GetAll() []providers.Core {
	records := []Provider{}
	repo.Conn.Find(&records)

	return ToCoreSlice(records)
}

func (repo *posgresRepository) GetById(id int) (*providers.Core, error) {
	record := Provider{Id: id}
	if err := repo.Conn.First(&record).Error; err != nil {
		return &providers.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *posgresRepository) Update(data *providers.Core) (*providers.Core, error) {
	record := FromCore(data)
	if err := repo.Conn.Model(record).Updates(record).Error; err != nil {
		return &providers.Core{}, err
	}
	return record.ToCore(), nil
}

func (repo *posgresRepository) Delete(id int) error {
	return repo.Conn.Delete(&Provider{Id: id}).Error
}
