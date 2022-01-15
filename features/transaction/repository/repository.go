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
