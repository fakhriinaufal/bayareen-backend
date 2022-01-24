package service

import (
	"bayareen-backend/features/categories"

	"github.com/go-playground/validator/v10"
)

type categoryUsecase struct {
	categoryData categories.Data
	validator    *validator.Validate
}

func NewCategoryUsecase(categoryData categories.Data) categories.Business {
	return &categoryUsecase{
		categoryData: categoryData,
		validator:    validator.New(),
	}
}

func (cu *categoryUsecase) Create(core categories.Core) (resp categories.Core, err error) {
	if err = cu.validator.Struct(core); err != nil {
		return categories.Core{}, err
	}

	resp, err = cu.categoryData.Create(core)

	if err != nil {
		return categories.Core{}, err
	}

	return resp, nil
}

func (cu *categoryUsecase) GetAll() []categories.Core {
	return cu.categoryData.GetAll()
}

func (cu *categoryUsecase) GetByName(name string) (categories.Core, error) {
	var nameQuery string
	if name == "pulsa" {
		nameQuery = "Pulsa"
	}
	if name == "paket" {
		nameQuery = "Paket"
	}
	if name == "pdam" {
		nameQuery = "PDAM"
	}
	if name == "listrik" {
		nameQuery = "Listrik"
	}

	return cu.categoryData.GetByName(nameQuery)
}

func (cu *categoryUsecase) GetById(id int) (categories.Core, error) {
	resp, err := cu.categoryData.GetById(id)

	if err != nil {
		return categories.Core{}, err
	}

	return resp, nil
}

func (cu *categoryUsecase) Update(core categories.Core) (resp categories.Core, err error) {
	if err = cu.validator.Struct(core); err != nil {
		return categories.Core{}, err
	}

	// existedCategory, err := cu.categoryData.GetById(core.Id)
	// if err != nil {
	// 	return categories.Core{}, err
	// }

	// core.CreatedAt = existedCategory.CreatedAt

	resp, err = cu.categoryData.Update(core)
	if err != nil {
		return categories.Core{}, err
	}

	return resp, nil
}

func (cu *categoryUsecase) Delete(id int) error {
	return cu.categoryData.Delete(id)
}
