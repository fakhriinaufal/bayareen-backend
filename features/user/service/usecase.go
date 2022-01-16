package service

import (
	"bayareen-backend/features/user"
	"bayareen-backend/helper/bcrypt"
	"bayareen-backend/middleware"
	"errors"

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

	data.Password, err = bcrypt.Hash(data.Password)
	if err != nil {
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

func (uu *userUseCase) Update(data user.UserCore) (user.UserCore, error) {
	if err := uu.validator.Struct(data); err != nil {
		return user.UserCore{}, err
	}

	existedUser, err := uu.userData.GetById(data.Id)
	if err != nil {
		return user.UserCore{}, err
	}

	data.CreatedAt = existedUser.CreatedAt

	resp, err := uu.userData.Update(data)
	if err != nil {
		return user.UserCore{}, err
	}

	return resp, nil
}

func (uu *userUseCase) Delete(id int) error {
	// check user exist
	_, err := uu.userData.GetById(id)
	if err != nil {
		return err
	}

	return uu.userData.Delete(id)
}

func (uc *userUseCase) Login(core user.UserCore) (user.UserCore, error) {
	if core.Email == "" {
		return user.UserCore{}, errors.New("email empty")
	}

	if core.Password == "" {
		return user.UserCore{}, errors.New("password empty")
	}

	userData, err := uc.userData.Login(core)
	if err != nil {
		return user.UserCore{}, err
	}
	temp := bcrypt.ValidateHash(core.Password, userData.Password)

	if !temp {
		return user.UserCore{}, errors.New("password salah")
	}

	userData.Token, err = middleware.CreateToken(core.Id, false)
	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}
