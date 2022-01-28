package service_test

import (
	"bayareen-backend/config"
	"bayareen-backend/features/admins"
	mockAdminRepo "bayareen-backend/features/admins/mocks"
	"bayareen-backend/features/admins/service"
	"bayareen-backend/middleware"
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	adminRepo    mockAdminRepo.Data
	adminService admins.Business
	adminCore    *admins.Core
	validate     *validator.Validate
	jwtSecret    config.JWTSecret
)

func setup() {
	jwtSecret = config.JWTSecret{Secret: "lorem123"}
	adminService = service.NewAdminUsecase(&adminRepo, jwtSecret)
	validate = validator.New()
	adminCore = &admins.Core{
		Id:       1,
		Name:     "Mr Admin",
		Password: "apple",
	}
}

func TestLogin(t *testing.T) {
	setup()

	adminRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(adminCore, nil).Once()
	adminRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&admins.Core{}, nil).Once()

	t.Run("Test Case 1 | Success login", func(t *testing.T) {
		result, err := adminService.Login("lorem", "ipsum")

		tokenResult, _ := middleware.CreateToken(1, true, jwtSecret.Secret)
		assert.Equal(t, tokenResult, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Wrong credentials", func(t *testing.T) {
		result, err := adminService.Login("lorem", "ipsum")

		assert.Equal(t, "", result)
		assert.Equal(t, errors.New("wrong credentials"), err)
	})
}

func TestCreate(t *testing.T) {
	setup()

	adminRepo.On("Create", mock.AnythingOfType("*admins.Core")).Return(adminCore, nil).Once()

	adminRepo.On("Create", mock.AnythingOfType("*admins.Core")).Return(&admins.Core{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Create", func(t *testing.T) {
		result, err := adminService.Create(adminCore)

		assert.Equal(t, adminCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Missing Field", func(t *testing.T) {
		badAdminInput := &admins.Core{}
		badAdminInputErr := validate.Struct(badAdminInput)
		result, err := adminService.Create(badAdminInput)

		assert.Equal(t, &admins.Core{}, result)
		assert.Equal(t, badAdminInputErr.Error(), err.Error())
	})

	t.Run("Test Case 3 | Error from database", func(t *testing.T) {
		result, err := adminService.Create(adminCore)

		assert.Equal(t, &admins.Core{}, result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}

func TestGetAll(t *testing.T) {
	setup()

	adminRepo.On("GetAll").Return([]admins.Core{*adminCore}).Once()

	t.Run("Test Case 1 | Success Get All", func(t *testing.T) {
		result := adminService.GetAll()

		assert.Equal(t, []admins.Core{*adminCore}, result)
	})
}

func TestGetById(t *testing.T) {
	setup()

	adminRepo.On("GetById", mock.AnythingOfType("int")).Return(adminCore, nil).Once()
	adminRepo.On("GetById", mock.AnythingOfType("int")).Return(&admins.Core{}, errors.New("admin doesn't exist")).Once()

	t.Run("Test Case 1 | Success get by id", func(t *testing.T) {
		result, err := adminService.GetById(1)

		assert.Equal(t, adminCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Admin doesn't exist", func(t *testing.T) {
		result, err := adminService.GetById(1)

		assert.Equal(t, &admins.Core{}, result)
		assert.Equal(t, errors.New("admin doesn't exist"), err)
	})
}

func TestUpdate(t *testing.T) {
	setup()
	adminRepo.On("GetById", mock.AnythingOfType("int")).Return(adminCore, nil).Once()
	adminRepo.On("Update", mock.AnythingOfType("*admins.Core")).Return(adminCore, nil).Once()

	adminRepo.On("GetById", mock.AnythingOfType("int")).Return(&admins.Core{}, errors.New("admin doesn't exist")).Once()

	adminRepo.On("GetById", mock.AnythingOfType("int")).Return(adminCore, nil).Once()

	adminRepo.On("GetById", mock.AnythingOfType("int")).Return(adminCore, nil).Once()
	adminRepo.On("Update", mock.AnythingOfType("*admins.Core")).Return(&admins.Core{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Update", func(t *testing.T) {
		result, err := adminService.Update(adminCore)

		assert.Equal(t, adminCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Admin doesn't exist", func(t *testing.T) {
		result, err := adminService.Update(adminCore)

		assert.Equal(t, &admins.Core{}, result)
		assert.Equal(t, errors.New("admin doesn't exist"), err)
	})

	t.Run("Test Case 3 | Missing field", func(t *testing.T) {
		badAdminInput := &admins.Core{}
		badAdminInputErr := validate.Struct(badAdminInput)
		result, err := adminService.Update(badAdminInput)

		assert.Equal(t, &admins.Core{}, result)
		assert.Equal(t, badAdminInputErr.Error(), err.Error())
	})

	t.Run("Test Case 4 | Error from database", func(t *testing.T) {
		result, err := adminService.Update(adminCore)

		assert.Equal(t, &admins.Core{}, result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}

func TestDelete(t *testing.T) {
	setup()

	adminRepo.On("GetById", mock.AnythingOfType("int")).Return(adminCore, nil).Once()
	adminRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

	adminRepo.On("GetById", mock.AnythingOfType("int")).Return(&admins.Core{}, errors.New("admin doesn't exist")).Once()

	adminRepo.On("GetById", mock.AnythingOfType("int")).Return(adminCore, nil).Once()
	adminRepo.On("Delete", mock.AnythingOfType("int")).Return(errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success delete", func(t *testing.T) {
		err := adminService.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Admin doesn't exist", func(t *testing.T) {
		err := adminService.Delete(1)

		assert.Equal(t, errors.New("admin doesn't exist"), err)
	})

	t.Run("Test Case 3 | Error from database", func(t *testing.T) {
		err := adminService.Delete(1)

		assert.Equal(t, errors.New("error happened"), err)
	})
}
