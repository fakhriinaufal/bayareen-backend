package response

import (
	"bayareen-backend/features/providers"
	"time"
)

type Provider struct {
	Id        int       `json:"id"`
	CatId     int       `json:"cat_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromCore(core *providers.Core) *Provider {
	return &Provider{
		Id:        core.CatId,
		CatId:     core.CatId,
		Name:      core.Name,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}

func FromCoreSlice(coreSlice []providers.Core) []Provider {
	resp := []Provider{}

	for _, val := range coreSlice {
		resp = append(resp, *FromCore(&val))
	}

	return resp
}
