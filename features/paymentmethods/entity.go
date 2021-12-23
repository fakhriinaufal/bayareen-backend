package paymentmethods

import "time"

type Core struct {
	Id        int
	Name      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
}

type Data interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
}
