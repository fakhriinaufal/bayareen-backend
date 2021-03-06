package products

import "time"

type Core struct {
	Id         int
	ProviderId int
	CatId      int
	Name       string `validate:"required"`
	Price      int    `validate:"required"`
	Status     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Business interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	Update(data *Core) (*Core, error)
	Delete(id []int) error
	GetByProviderId(providerId int) ([]Core, error)
}

type Data interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	Update(data *Core) (*Core, error)
	Delete(id []int) error
	GetByProviderId(providerId int) ([]Core, error)
}
