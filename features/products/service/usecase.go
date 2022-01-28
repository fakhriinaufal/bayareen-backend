package service

import (
	"bayareen-backend/features/categories"
	"bayareen-backend/features/products"
	"bayareen-backend/features/providers"

	"github.com/go-playground/validator/v10"
)

type productUsecase struct {
	productData  products.Data
	categoryData categories.Data
	providerData providers.Data
	validator    *validator.Validate
}

func NewProductUsecase(prodData products.Data, catData categories.Data, provData providers.Data) products.Business {
	return &productUsecase{
		productData:  prodData,
		categoryData: catData,
		providerData: provData,
		validator:    validator.New(),
	}
}

func (pu *productUsecase) Create(data *products.Core) (*products.Core, error) {
	if err := pu.validator.Struct(data); err != nil {
		return &products.Core{}, err
	}

	if data.CatId != 0 {
		// check is category exist
		_, err := pu.categoryData.GetById(data.CatId)
		if err != nil {
			return &products.Core{}, err
		}
	}

	if data.ProviderId != 0 {
		// check is provider exist
		_, err := pu.providerData.GetById(data.ProviderId)
		if err != nil {
			return &products.Core{}, err
		}
	}

	resp, err := pu.productData.Create(data)
	if err != nil {
		return &products.Core{}, err
	}

	return resp, nil
}

func (pu *productUsecase) GetAll() []products.Core {
	return pu.productData.GetAll()
}

func (pu *productUsecase) GetById(id int) (*products.Core, error) {
	return pu.productData.GetById(id)
}

func (pu *productUsecase) Update(data *products.Core) (*products.Core, error) {
	if err := pu.validator.Struct(data); err != nil {
		return &products.Core{}, err
	}

	if data.CatId != 0 {
		// check is category exist
		_, err := pu.categoryData.GetById(data.CatId)
		if err != nil {
			return &products.Core{}, err
		}
	}

	if data.ProviderId != 0 {
		// check is provider exist
		_, err := pu.providerData.GetById(data.ProviderId)
		if err != nil {
			return &products.Core{}, err
		}
	}

	// existedProduct, err := pu.productData.GetById(data.Id)
	// if err != nil {
	// 	return &products.Core{}, err
	// }

	// data.CreatedAt = existedProduct.CreatedAt

	resp, err := pu.productData.Update(data)
	if err != nil {
		return &products.Core{}, err
	}

	return resp, nil
}

func (pu *productUsecase) Delete(id []int) error {
	return pu.productData.Delete(id)
}

func (pu *productUsecase) GetByProviderId(provId int) ([]products.Core, error) {
	data, err := pu.productData.GetByProviderId(provId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
