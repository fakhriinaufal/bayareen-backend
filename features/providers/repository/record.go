package repository

import (
	"bayareen-backend/features/providers"
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	Id        int
	CatId     int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (provider *Provider) ToCore() *providers.Core {
	return &providers.Core{
		Id:        provider.Id,
		CatId:     provider.CatId,
		Name:      provider.Name,
		CreatedAt: provider.CreatedAt,
		UpdatedAt: provider.UpdatedAt,
	}
}

func ToCoreSlice(providerSlice []Provider) []providers.Core {
	resp := []providers.Core{}

	for _, val := range providerSlice {
		resp = append(resp, *val.ToCore())
	}

	return resp
}

func FromCore(core *providers.Core) *Provider {
	return &Provider{
		Id:        core.Id,
		CatId:     core.CatId,
		Name:      core.Name,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}
