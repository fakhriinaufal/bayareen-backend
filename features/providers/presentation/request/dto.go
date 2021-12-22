package request

import "bayareen-backend/features/providers"

type Provider struct {
	CatId int    `json:"cat_id"`
	Name  string `json:"name"`
}

func (p *Provider) ToCore() *providers.Core {
	return &providers.Core{
		CatId: p.CatId,
		Name:  p.Name,
	}
}
