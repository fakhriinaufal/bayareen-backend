package service_test

import (
	"bayareen-backend/features/providers"
	mockProviderRepo "bayareen-backend/features/providers/mocks"
	"bayareen-backend/features/providers/service"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	providerRepo    mockProviderRepo.Data
	providerService providers.Business
	providerCore    *providers.Core
	validate        *validator.Validate
)

func setup() {
	providerService = service.NewProviderUsecase(&providerRepo)
	providerCore = &providers.Core{
		Id:    1,
		CatId: 1,
		Name:  "Telkomsel",
	}
	validate = validator.New()
}

func TestCreate(t *testing.T) {
	setup()
	providerRepo.On("Create", mock.AnythingOfType("*providers.Core")).Return(providerCore, nil).Once()
	providerRepo.On("Create", mock.AnythingOfType("*providers.Core")).Return(&providers.Core{}, errors.New("can't create provider")).Once()

	t.Run("Test Case 1 | Success Create Provider", func(t *testing.T) {
		result, err := providerService.Create(providerCore)

		assert.Equal(t, providerCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Missing field on struct", func(t *testing.T) {
		badInput := providers.Core{}
		badInputErr := validate.Struct(&badInput)
		result, err := providerService.Create(&badInput)

		assert.Equal(t, &providers.Core{}, result)
		assert.Equal(t, badInputErr.Error(), err.Error())
	})

	t.Run("Test Case 3 | Error from db", func(t *testing.T) {
		result, err := providerService.Create(providerCore)

		assert.Equal(t, &providers.Core{}, result)
		assert.Equal(t, errors.New("can't create provider"), err)
	})
}

func TestGetAll(t *testing.T) {
	setup()
	providerRepo.On("GetAll").Return([]providers.Core{*providerCore}).Once()

	t.Run("Test Case 1 | Success Get All", func(t *testing.T) {
		result := providerService.GetAll()

		assert.Equal(t, []providers.Core{*providerCore}, result)
	})
}

func TestGetByCategoryId(t *testing.T) {
	setup()
	providerRepo.On("GetByCategoryId", mock.AnythingOfType("int")).Return([]providers.Core{*providerCore}, nil).Once()
	providerRepo.On("GetByCategoryId", mock.AnythingOfType("int")).Return([]providers.Core{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Get By Category Id", func(t *testing.T) {
		result, err := providerService.GetByCategoryId(1)

		assert.Equal(t, []providers.Core{*providerCore}, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Error from database", func(t *testing.T) {
		result, err := providerService.GetByCategoryId(1)

		assert.Equal(t, []providers.Core(nil), result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}

func TestGetById(t *testing.T) {
	setup()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(providerCore, nil).Once()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Get By Id", func(t *testing.T) {
		result, err := providerService.GetById(1)

		assert.Equal(t, providerCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Error from database", func(t *testing.T) {
		result, err := providerService.GetById(1)

		assert.Equal(t, &providers.Core{}, result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}

func TestUpdate(t *testing.T) {
	setup()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(providerCore, nil).Once()
	providerRepo.On("Update", mock.AnythingOfType("*providers.Core")).Return(providerCore, nil).Once()

	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, errors.New("provider doesn't exist")).Once()

	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(providerCore, nil).Once()

	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(providerCore, nil).Once()
	providerRepo.On("Update", mock.AnythingOfType("*providers.Core")).Return(&providers.Core{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Update", func(t *testing.T) {
		result, err := providerService.Update(providerCore)

		assert.Equal(t, providerCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Provider doesn't exist", func(t *testing.T) {
		result, err := providerService.Update(providerCore)

		assert.Equal(t, &providers.Core{}, result)
		assert.Equal(t, errors.New("provider doesn't exist"), err)
	})

	t.Run("Test Case 3 | Missing field", func(t *testing.T) {
		badInput := providers.Core{}
		badInputErr := validate.Struct(&badInput)
		result, err := providerService.Update(&badInput)

		assert.Equal(t, &providers.Core{}, result)
		assert.Equal(t, badInputErr.Error(), err.Error())
	})

	t.Run("Test Case 4 | Error from database", func(t *testing.T) {
		result, err := providerService.Update(providerCore)

		assert.Equal(t, &providers.Core{}, result)
		assert.Equal(t, errors.New("error happened"), err)

	})
}

func TestDelete(t *testing.T) {
	setup()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(providerCore, nil).Once()
	providerRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, errors.New("provider doesn't exist")).Once()

	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(providerCore, nil).Once()
	providerRepo.On("Delete", mock.AnythingOfType("int")).Return(errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Delete", func(t *testing.T) {
		err := providerService.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Provider doesn't exist", func(t *testing.T) {
		err := providerService.Delete(1)

		assert.Equal(t, errors.New("provider doesn't exist"), err)
	})

	t.Run("Test Case 2 | Error from database", func(t *testing.T) {
		err := providerService.Delete(1)

		assert.Equal(t, errors.New("error happened"), err)
	})
}
