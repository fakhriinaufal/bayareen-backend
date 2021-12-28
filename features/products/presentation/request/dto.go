package request

import "bayareen-backend/features/products"

type Product struct {
	ProviderId int    `json:"provider_id"`
	CatId      int    `json:"cat_id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Status     bool   `json:"status"`
}

func (p *Product) ToCore() *products.Core {
	return &products.Core{
		ProviderId: p.ProviderId,
		CatId:      p.CatId,
		Name:       p.Name,
		Price:      p.Price,
		Status:     p.Status,
	}
}
