package categories

import "time"

type Core struct {
	Id        int
	Name      string `validate:"required"`
	ImgUrl    string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Create(Core) (Core, error)
	GetAll() []Core
	GetById(id int) (Core, error)
	Update(Core) (Core, error)
}

type Data interface {
	Create(Core) (Core, error)
	GetAll() []Core
	GetById(id int) (Core, error)
	Update(Core) (Core, error)
}
