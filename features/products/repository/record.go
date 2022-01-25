package repository

import (
	cat_repo "bayareen-backend/features/categories/repository"
	"bayareen-backend/features/products"
	prov_repo "bayareen-backend/features/providers/repository"
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
	Category   cat_repo.Category  `gorm:"foreignKey:CatId"`
	Provider   prov_repo.Provider `gorm:"foreignKey:ProviderId"`
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
