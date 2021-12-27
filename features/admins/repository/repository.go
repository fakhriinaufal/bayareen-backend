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

func (repo *postgresUserRepository) GetAll() []admins.Core {
	records := []Admin{}
	repo.Conn.Find(&records)
	return ToCoreSlice(records)
}

func (repo *postgresUserRepository) GetById(id int) (*admins.Core, error) {
	record := Admin{
		Id: id,
	}
	if err := repo.Conn.First(&record).Error; err != nil {
		return &admins.Core{}, err
	}
	return record.ToCore(), nil
}

func (repo *postgresUserRepository) Update(data *admins.Core) (*admins.Core, error) {
	record := FromCore(data)
	if err := repo.Conn.Save(&record).Error; err != nil {
		return &admins.Core{}, err
	}
	return record.ToCore(), nil
}

func (repo *postgresUserRepository) Delete(id int) error {
	return repo.Conn.Delete(&Admin{Id: id}).Error
}
