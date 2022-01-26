package transaction

import (
	"bayareen-backend/features/products"
	"time"
)

type Core struct {
	Id              int
	UserId          int
	ProductId       int
	Price           int
	InvoiceId       string
	InvoiceUrl      string
	ReferenceId     string
	PaymentMethodId int
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Product         products.Core
}

type XenditCallback struct {
	Id             string
	ReferenceId    string
	Status         string
	PaymentMethod  string
	PaymentChannel string
}

type Data interface {
	Create(*Core) (*Core, error)
	Update(*Core) (*Core, error)
	UpdateByReferenceId(*Core) (*Core, error)
	GetByUserId(userId int) ([]Core, error)
}

type Business interface {
	Create(*Core) (*Core, error)
	UpdatePayment(XenditCallback) error
}
