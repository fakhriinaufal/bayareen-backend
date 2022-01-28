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

func (pu *providerUsecase) GetByCategoryId(catId int) ([]providers.Core, error) {
	result, err := pu.ProviderData.GetByCategoryId(catId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pu *providerUsecase) GetById(id int) (*providers.Core, error) {
	resp, err := pu.ProviderData.GetById(id)
	if err != nil {
		return &providers.Core{}, err
	}

	return resp, nil
}

func (pu *providerUsecase) Update(data *providers.Core) (*providers.Core, error) {
	if err := pu.Validator.Struct(data); err != nil {
		return &providers.Core{}, err
	}

	// existedUser, err := pu.ProviderData.GetById(data.Id)
	// if err != nil {
	// 	return &providers.Core{}, err
	// }
	// data.CreatedAt = existedUser.CreatedAt

	resp, err := pu.ProviderData.Update(data)
	if err != nil {
		return &providers.Core{}, err
	}

	return resp, nil
}

func (pu *providerUsecase) Delete(id int) error {
	var err error
	_, err = pu.ProviderData.GetById(id)

	if err != nil {
		return err
	}

	return pu.ProviderData.Delete(id)
}
