package transaction

import (
	"bayareen-backend/features/payment_gateway"
	"time"
)

type Core struct {
	Id              int
	UserId          int
	ProductId       int
	InvoiceId       string
	PaymentMethodId int
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type XenditCallback struct {
	Id             string
	TransactionId  string
	Status         string
	PaymentMethod  string
	PaymentChannel string
}

type Data interface {
	Create(*Core) (*Core, error)
	Update(*Core) (*Core, error)
}

type Business interface {
	Create(*Core) (*Core, payment_gateway.InvoiceData, error)
	UpdatePayment(XenditCallback) error
}
