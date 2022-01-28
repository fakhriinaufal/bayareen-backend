package paymentmethods

import "time"

type Core struct {
	Id             int
	PaymentMethod  string `validate:"required"`
	PaymentChannel string `validate:"required"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Business interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	Update(data *Core) (*Core, error)
	Delete(id int) error
	GetByName(method string, channel string) (int, error)
}

type Data interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	Update(data *Core) (*Core, error)
	Delete(id int) error
	GetByName(method string, channel string) (int, error)
}
