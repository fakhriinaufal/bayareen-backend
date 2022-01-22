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

func (repo *mySqlRepository) GetById(id int) (user.UserCore, error) {
	resp := User{}

	if err := repo.Conn.First(&resp, id).Error; err != nil {
		return user.UserCore{}, err
	}

	return resp.toCore(), nil
}

func (repo *mySqlRepository) Update(data user.UserCore) (user.UserCore, error) {
	record := fromCore(data)
	err := repo.Conn.Save(&record).Error
	if err != nil {
		return user.UserCore{}, err
	}

	return record.toCore(), nil
}

func (repo *mySqlRepository) Delete(id int) error {
	record := User{
		Id: id,
	}
	return repo.Conn.Delete(&record).Error
}

func (repo *mySqlRepository) Login(data user.UserCore) (user.UserCore, error) {
	var userRecord User

	repo.Conn.Raw("SELECT password FROM users WHERE email = ?", data.Email).Scan(&data.Password)

	if err := repo.Conn.First(&userRecord, "email = ?", data.Email).Error; err != nil {
		return user.UserCore{}, err
	}

	return userRecord.toCore(), nil
}

func (repo *mySqlRepository) GetByEmail(email string) (user.UserCore, error) {
	var userRecord User
	err := repo.Conn.Where("email = ?", email).Limit(1).Find(&userRecord).Error
	if err != nil {
		return user.UserCore{}, nil
	}

	return userRecord.toCore(), nil
}
