package categories

import "time"

type Core struct {
	Id        int
	Name      string
	ImgUrl    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Create(Core) (Core, error)
	GetAll() []Core
}

type Data interface {
	Create(Core) (Core, error)
	GetAll() []Core
}
