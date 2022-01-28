package repository

import (
	"bayareen-backend/features/categories"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (c *Category) ToCore() categories.Core {
	return categories.Core{
		Id:        c.Id,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToCoreSlice(cSlice *[]Category) []categories.Core {
	resp := []categories.Core{}
	for _, val := range *cSlice {
		resp = append(resp, val.ToCore())
	}
	return resp
}

func FromCore(core categories.Core) Category {
	return Category{
		Id:        core.Id,
		Name:      core.Name,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}
