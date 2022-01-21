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

	userData, err := uu.userData.GetByEmail(data.Email)
	if err != nil {
		return user.UserCore{}, err
	}
	if userData.Id != 0 {
		return user.UserCore{}, errors.New("email already exist")
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

	userData.Token, err = middleware.CreateToken(userData.Id, false)
	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}

func (uc *userUseCase) UpdatePassword(core user.UserUpdatePasswordCore) (user.UserCore, error) {
	if err := uc.validator.Struct(core); err != nil {
		return user.UserCore{}, err
	}

	existedUser, err := uc.userData.GetById(core.ID)
	if err != nil {
		return user.UserCore{}, err
	}

	if !bcrypt.ValidateHash(core.OldPassword, existedUser.Password) {
		return user.UserCore{}, errors.New("wrong old password")
	}

	existedUser.Password, err = bcrypt.Hash(core.NewPassword)
	if err != nil {
		return user.UserCore{}, err
	}

	updatedUser, err := uc.userData.Update(existedUser)
	if err != nil {
		return user.UserCore{}, err
	}

	return updatedUser, nil
}

func (uc *userUseCase) UpdateProfile(core user.UserCore) (user.UserCore, error) {
	if core.Name == "" {
		return user.UserCore{}, errors.New("name required")
	}

	if core.PhoneNumber == "" {
		return user.UserCore{}, errors.New("phone number required")
	}

	if core.Email == "" {
		return user.UserCore{}, errors.New("email required")
	}

	userData, err := uc.userData.GetByEmail(core.Email)
	if err != nil {
		return user.UserCore{}, err
	}
	if userData.Id != 0 {
		return user.UserCore{}, errors.New("email already exist")
	}

	existedUser, err := uc.userData.GetById(core.Id)
	if err != nil {
		return user.UserCore{}, err
	}

	existedUser.Name = core.Name
	existedUser.PhoneNumber = core.PhoneNumber
	existedUser.Email = core.Email

	updatedUser, err := uc.userData.Update(existedUser)
	if err != nil {
		return user.UserCore{}, err
	}

	return updatedUser, nil
}
