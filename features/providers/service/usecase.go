package service

import (
	"bayareen-backend/features/providers"

	"github.com/go-playground/validator/v10"
)

type providerUsecase struct {
	ProviderData providers.Data
	Validator    *validator.Validate
}

func NewProviderUsecase(providerData providers.Data) providers.Business {
	return &providerUsecase{
		ProviderData: providerData,
		Validator:    validator.New(),
	}
}

func (pu *providerUsecase) Create(data *providers.Core) (*providers.Core, error) {
	if err := pu.Validator.Struct(data); err != nil {
		return &providers.Core{}, err
	}

	resp, err := pu.ProviderData.Create(data)

	if err != nil {
		return &providers.Core{}, err
	}

	return resp, nil
}

func (pu *providerUsecase) GetAll() []providers.Core {
	return pu.ProviderData.GetAll()
}

func (pu *providerUsecase) GetById(id int) (*providers.Core, error) {
	resp, err := pu.ProviderData.GetById(id)
	if err != nil {
		return &providers.Core{}, err
	}

	return resp, nil
}
