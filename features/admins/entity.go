package admins

import "time"

type Core struct {
	Id        int
	Name      string `validate:"required"`
	Password  string `validate:"required"`
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	Update(data *Core) (*Core, error)
	Delete(id int) error
	Login(name, password string) (string, error)
	JWTLogin(id int) error
}

type Data interface {
	Create(data *Core) (*Core, error)
	GetAll() []Core
	GetById(id int) (*Core, error)
	Update(data *Core) (*Core, error)
	Delete(id int) error
	Login(username, password string) (*Core, error)
}
