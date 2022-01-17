package repository

import (
	"bayareen-backend/features/transaction"

	"gorm.io/gorm"
)

type transactionRepository struct {
	Conn *gorm.DB
}

func NewPostgresTransactionRepository(conn *gorm.DB) transaction.Data {
	return &transactionRepository{
		Conn: conn,
	}
}

func (tr *transactionRepository) Create(data *transaction.Core) (*transaction.Core, error) {
	record := FromCore(data)
	if err := tr.Conn.Create(record).Error; err != nil {
		return &transaction.Core{}, err
	}

	return record.ToCore(), nil
}

func (tr *transactionRepository) Update(data *transaction.Core) (*transaction.Core, error) {
	record := FromCore(data)

	if err := tr.Conn.Updates(record).Error; err != nil {
		return &transaction.Core{}, err
	}

	return record.ToCore(), nil
}

func (tr *transactionRepository) GetByUserId(userId int) ([]transaction.Core, error) {
	var transactions []Transaction
	err := tr.Conn.Debug().Joins("Product").Where(&Transaction{UserId: userId}).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return ToCoreList(transactions), nil
}
