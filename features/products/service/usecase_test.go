package service_test

import (
	"bayareen-backend/features/categories"
	mockCategoryData "bayareen-backend/features/categories/mocks"
	"bayareen-backend/features/products"
	mockProductData "bayareen-backend/features/products/mocks"
	"bayareen-backend/features/products/service"
	"bayareen-backend/features/providers"
	mockProviderData "bayareen-backend/features/providers/mocks"
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	productRepo    mockProductData.Data
	productService products.Business
	productCore    *products.Core
	categoryRepo   mockCategoryData.Data
	providerRepo   mockProviderData.Data
	validate       *validator.Validate
)

func setup() {
	productService = service.NewProductUsecase(&productRepo, &categoryRepo, &providerRepo)
	validate = validator.New()
	productCore = &products.Core{
		Id:         1,
		ProviderId: 1,
		CatId:      1,
		Name:       "Pulsa",
		Price:      20000,
		Status:     true,
	}
}

func TestCreate(t *testing.T) {
	setup()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, nil).Once()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, nil).Once()
	productRepo.On("Create", mock.AnythingOfType("*products.Core")).Return(productCore, nil).Once()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, errors.New("category doesn't exist")).Once()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, nil).Once()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, errors.New("provider doesn't exist")).Once()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, nil).Once()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, nil).Once()
	productRepo.On("Create", mock.AnythingOfType("*products.Core")).Return(productCore, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Create Product", func(t *testing.T) {
		result, err := productService.Create(productCore)

		assert.Equal(t, productCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Missing field", func(t *testing.T) {
		badInput := &products.Core{}
		badInputErr := validate.Struct(badInput)
		result, err := productService.Create(badInput)

		assert.Equal(t, &products.Core{}, result)
		assert.Equal(t, badInputErr.Error(), err.Error())
	})

	t.Run("Test Case 3 | Category Doesn't exist", func(t *testing.T) {
		result, err := productService.Create(productCore)

		assert.Equal(t, &products.Core{}, result)
		assert.Equal(t, errors.New("category doesn't exist"), err)

	})

	t.Run("Test Case 4 | Provider Doesn't exist", func(t *testing.T) {
		result, err := productService.Create(productCore)

		assert.Equal(t, &products.Core{}, result)
		assert.Equal(t, errors.New("provider doesn't exist"), err)

	})

	t.Run("Test Case 5 | Error from database", func(t *testing.T) {
		result, err := productService.Create(productCore)

		assert.Equal(t, &products.Core{}, result)
		assert.Equal(t, errors.New("error happened"), err)

	})
}

func TestGetAll(t *testing.T) {
	setup()

	productRepo.On("GetAll").Return([]products.Core{*productCore}).Once()

	t.Run("Test Case 1 | Success get all", func(t *testing.T) {
		result := productService.GetAll()

		assert.Equal(t, []products.Core{*productCore}, result)
	})
}

func TestGetById(t *testing.T) {
	setup()

	productRepo.On("GetById", mock.AnythingOfType("int")).Return(productCore, nil).Once()

	t.Run("Test Case 1 | Success get by id", func(t *testing.T) {
		result, err := productService.GetById(1)

		assert.Equal(t, productCore, result)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	setup()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, nil).Once()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, nil).Once()
	productRepo.On("Update", mock.AnythingOfType("*products.Core")).Return(productCore, nil).Once()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, errors.New("category doesn't exist")).Once()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, nil).Once()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, errors.New("provider doesn't exist")).Once()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, nil).Once()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, nil).Once()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, nil).Once()
	providerRepo.On("GetById", mock.AnythingOfType("int")).Return(&providers.Core{}, nil).Once()
	productRepo.On("Update", mock.AnythingOfType("*products.Core")).Return(productCore, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Update Product", func(t *testing.T) {
		result, err := productService.Update(productCore)

		assert.Equal(t, productCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Missing field", func(t *testing.T) {
		badInput := &products.Core{}
		badInputErr := validate.Struct(badInput)
		result, err := productService.Update(badInput)

		assert.Equal(t, &products.Core{}, result)
		assert.Equal(t, badInputErr.Error(), err.Error())
	})

	t.Run("Test Case 3 | Category Doesn't exist", func(t *testing.T) {
		result, err := productService.Update(productCore)

		assert.Equal(t, &products.Core{}, result)
		assert.Equal(t, errors.New("category doesn't exist"), err)

	})

	t.Run("Test Case 4 | Provider Doesn't exist", func(t *testing.T) {
		result, err := productService.Update(productCore)

		assert.Equal(t, &products.Core{}, result)
		assert.Equal(t, errors.New("provider doesn't exist"), err)

	})

	t.Run("Test Case 6 | Error from database", func(t *testing.T) {
		result, err := productService.Update(productCore)

		assert.Equal(t, &products.Core{}, result)
		assert.Equal(t, errors.New("error happened"), err)

	})
}

func TestDelete(t *testing.T) {
	setup()

	productRepo.On("Delete", mock.AnythingOfType("[]int")).Return(nil).Once()

	t.Run("Test Case 1 | Success delete by id", func(t *testing.T) {
		err := productService.Delete([]int{1})

		assert.Nil(t, err)
	})
}

func TestGetByProviderId(t *testing.T) {
	setup()

	productRepo.On("GetByProviderId", mock.AnythingOfType("int")).Return([]products.Core{*productCore}, nil).Once()
	productRepo.On("GetByProviderId", mock.AnythingOfType("int")).Return([]products.Core{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success get by provider id", func(t *testing.T) {
		result, err := productService.GetByProviderId(1)

		assert.Equal(t, []products.Core{*productCore}, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Error from database", func(t *testing.T) {
		result, err := productService.GetByProviderId(1)

		assert.Equal(t, []products.Core(nil), result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}
