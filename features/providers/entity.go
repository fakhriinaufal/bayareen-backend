package providers

import "time"

type Core struct {
	Id        int
	CatId     int    `validate:"required"`
	Name      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	GetByCategoryId(catId int) ([]Core, error)
	Update(data *Core) (*Core, error)
	Delete(id int) error
}

type Data interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	GetByCategoryId(catId int) ([]Core, error)
	Update(data *Core) (*Core, error)
	Delete(id int) error
}
