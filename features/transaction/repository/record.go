package repository

import (
	"time"

	"gorm.io/gorm"

	"bayareen-backend/features/products/repository"
	"bayareen-backend/features/transaction"
)

type Transaction struct {
	Id              int
	UserId          int
	ProductId       int
	PaymentMethodId int
	Price           int
	InvoiceId       string
	InvoiceUrl      string
	ReferenceId     string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Product         repository.Product
}

func (t *Transaction) ToCore() *transaction.Core {
	return &transaction.Core{
		Id:              t.Id,
		UserId:          t.UserId,
		ProductId:       t.ProductId,
		PaymentMethodId: t.PaymentMethodId,
		InvoiceId:       t.InvoiceId,
		InvoiceUrl:      t.InvoiceUrl,
		ReferenceId:     t.ReferenceId,
		Price:           t.Price,
		Status:          t.Status,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
		Product:         *t.Product.ToCore(),
	}
}

func FromCore(data *transaction.Core) *Transaction {
	return &Transaction{
		Id:              data.Id,
		UserId:          data.UserId,
		ProductId:       data.ProductId,
		InvoiceId:       data.InvoiceId,
		InvoiceUrl:      data.InvoiceUrl,
		ReferenceId:     data.ReferenceId,
		PaymentMethodId: data.PaymentMethodId,
		Status:          data.Status,
		Price:           data.Price,
	}
}

func ToCoreList(data []Transaction) []transaction.Core {
	convert := []transaction.Core{}
	for _, tr := range data {
		convert = append(convert, *tr.ToCore())
	}
	return convert
}
