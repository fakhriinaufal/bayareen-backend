package repository

import (
	"bayareen-backend/features/categories"
	"errors"

	"gorm.io/gorm"
)

type postgreRepository struct {
	Conn *gorm.DB
}

func NewPostgreRepository(conn *gorm.DB) categories.Data {
	return &postgreRepository{
		Conn: conn,
	}
}

func (repo *postgreRepository) Create(core categories.Core) (categories.Core, error) {
	record := FromCore(core)
	if err := repo.Conn.Create(&record).Error; err != nil {
		return categories.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *postgreRepository) GetAll() []categories.Core {
	resp := []Category{}
	repo.Conn.Find(&resp)

	return ToCoreSlice(&resp)
}

func (repo *postgreRepository) GetByName(name string) (categories.Core, error) {
	resp := Category{}
	err := repo.Conn.Where("name = ?", name).First(&resp).Error
	if err != nil {
		return categories.Core{}, err
	}
	if resp.Id == 0 {
		return categories.Core{}, errors.New("category not found")
	}
	return resp.ToCore(), nil
}

func (repo *postgreRepository) GetById(id int) (categories.Core, error) {
	record := Category{
		Id: id,
	}

	err := repo.Conn.First(&record).Error
	if err != nil {
		return categories.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *postgreRepository) Update(core categories.Core) (categories.Core, error) {
	record := FromCore(core)

	if err := repo.Conn.Model(&record).Updates(&record).Error; err != nil {
		return categories.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *postgreRepository) Delete(id int) error {
	return repo.Conn.Delete(&Category{Id: id}).Error
}
