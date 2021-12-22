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
}

type Data interface {
	Create(data *Core) (*Core, error)
}
