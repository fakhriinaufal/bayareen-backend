package repository

import (
	"bayareen-backend/features/paymentmethods"

	"gorm.io/gorm"
)

type posgresPaymentMethodRepository struct {
	Conn *gorm.DB
}

func NewPostgresPaymentMethodRepository(conn *gorm.DB) paymentmethods.Data {
	return &posgresPaymentMethodRepository{
		Conn: conn,
	}
}

func (repo *posgresPaymentMethodRepository) Create(data *paymentmethods.Core) (*paymentmethods.Core, error) {
	record := FromCore(data)
	if err := repo.Conn.Create(record).Error; err != nil {
		return &paymentmethods.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *posgresPaymentMethodRepository) GetAll() []paymentmethods.Core {
	records := []PaymentMethod{}
	repo.Conn.Find(&records)

	return ToCoreSlice(records)
}

func (repo *posgresPaymentMethodRepository) GetById(id int) (*paymentmethods.Core, error) {
	record := PaymentMethod{
		Id: id,
	}
	if err := repo.Conn.First(&record).Error; err != nil {
		return &paymentmethods.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *posgresPaymentMethodRepository) Update(data *paymentmethods.Core) (*paymentmethods.Core, error) {
	record := FromCore(data)
	if err := repo.Conn.Save(&record).Error; err != nil {
		return &paymentmethods.Core{}, err
	}

	return record.ToCore(), nil
}

func (repo *posgresPaymentMethodRepository) Delete(id int) error {
	return repo.Conn.Delete(&PaymentMethod{Id: id}).Error
}

func (repo *posgresPaymentMethodRepository) GetByName(method, channel string) (int, error) {
	var paymentMethod PaymentMethod
	err := repo.Conn.Where("payment_method = ? AND payment_channel = ?", method, channel).First(&paymentMethod).Error
	if err != nil {
		return 0, err
	}
	return paymentMethod.Id, nil
}
