package service

import (
	"bayareen-backend/features/admins"

	"github.com/go-playground/validator/v10"
)

type adminUsecase struct {
	AdminData admins.Data
	validator *validator.Validate
}

func NewAdminUsecase(adminData admins.Data) admins.Business {
	return &adminUsecase{
		AdminData: adminData,
		validator: validator.New(),
	}
}

func (au *adminUsecase) Create(data *admins.Core) (*admins.Core, error) {
	if err := au.validator.Struct(data); err != nil {
		return &admins.Core{}, err
	}

	resp, err := au.AdminData.Create(data)
	if err != nil {
		return &admins.Core{}, err
	}
	return resp, nil
}
