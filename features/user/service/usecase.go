package service

import (
	"bayareen-backend/features/user"

	"github.com/go-playground/validator/v10"
)

type userUseCase struct {
	userData  user.Data
	validator *validator.Validate
}

func NewUserUsecase(userData user.Data) user.Business {
	return &userUseCase{
		userData:  userData,
		validator: validator.New(),
	}
}

func (uu *userUseCase) Create(data user.UserCore) (resp user.UserCore, err error) {
	if err := uu.validator.Struct(data); err != nil {
		return user.UserCore{}, err
	}

	resp, err = uu.userData.Create(data)

	if err != nil {
		return user.UserCore{}, err
	}

	return resp, nil
}

func (uu *userUseCase) GetAll() []user.UserCore {
	resp := uu.userData.GetAll()
	return resp
}

func (uu *userUseCase) GetById(id int) (user.UserCore, error) {
	resp, err := uu.userData.GetById(id)
	if err != nil {
		return user.UserCore{}, err
	}

	return resp, nil
}
