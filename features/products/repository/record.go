package repository

import (
	"bayareen-backend/features/products"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id         int
	ProviderId int
	CatId      int
	Name       string
	Price      int
	Status     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (p *Product) ToCore() *products.Core {
	return &products.Core{
		Id:         p.Id,
		ProviderId: p.ProviderId,
		CatId:      p.CatId,
		Name:       p.Name,
		Price:      p.Price,
		Status:     p.Status,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func ToCoreSlice(productSlice []Product) []products.Core {
	coreSlice := []products.Core{}
	for _, val := range productSlice {
		coreSlice = append(coreSlice, *val.ToCore())
	}
	return coreSlice
}

func FromCore(data *products.Core) *Product {
	return &Product{
		Id:         data.Id,
		ProviderId: data.ProviderId,
		CatId:      data.CatId,
		Name:       data.Name,
		Price:      data.Price,
		Status:     data.Status,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}
