package service

import (
	"bayareen-backend/config"
	"bayareen-backend/features/admins"
	"bayareen-backend/middleware"
	"errors"

	"github.com/go-playground/validator/v10"
)

type adminUsecase struct {
	AdminData admins.Data
	JWTSecret config.JWTSecret
	validator *validator.Validate
}

func NewAdminUsecase(adminData admins.Data, jwtSecret config.JWTSecret) admins.Business {
	return &adminUsecase{
		AdminData: adminData,
		validator: validator.New(),
		JWTSecret: jwtSecret,
	}
}

func (au *adminUsecase) Login(username, password string) (string, error) {
	admin, err := au.AdminData.Login(username, password)
	if err != nil {
		return "", err
	}
	if admin.Name == "" && admin.Password == "" {
		return "", errors.New("wrong credentials")
	}
	return middleware.CreateToken(admin.Id, true, au.JWTSecret.Secret)
}

func (au *adminUsecase) JWTLogin(id int) error {
	admin, err := au.AdminData.GetById(id)
	if err != nil {
		return err
	}
	if admin.Name == "" && admin.Password == "" {
		return errors.New("wrong credentials")
	}
	return nil
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

func (au *adminUsecase) Delete(id int) error {
	// check is admin exist
	_, err := au.AdminData.GetById(id)
	if err != nil {
		return err
	}

	return au.AdminData.Delete(id)
}
