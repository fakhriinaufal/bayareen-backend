package response

import (
	"bayareen-backend/features/products"
	"time"
)

type Product struct {
	Id         int       `json:"id"`
	ProviderId int       `json:"provider_id"`
	CatId      int       `json:"cat_id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
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

func FromCoreSlice(coreSlice []products.Core) []Product {
	resp := []Product{}
	for _, val := range coreSlice {
		resp = append(resp, *FromCore(&val))
	}
	return resp
}
