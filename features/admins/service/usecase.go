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

func (au *adminUsecase) GetAll() []admins.Core {
	return au.AdminData.GetAll()
}

func (au *adminUsecase) GetById(id int) (*admins.Core, error) {
	return au.AdminData.GetById(id)
}

func (au *adminUsecase) Update(data *admins.Core) (*admins.Core, error) {
	existedAdmin, err := au.AdminData.GetById(data.Id)
	if err != nil {
		return &admins.Core{}, err
	}

	if err := au.validator.Struct(data); err != nil {
		return &admins.Core{}, err
	}

	data.CreatedAt = existedAdmin.CreatedAt

	resp, err := au.AdminData.Update(data)
	if err != nil {
		return &admins.Core{}, err
	}
	return resp, nil
}
