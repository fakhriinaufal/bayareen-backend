package service_test

import (
	"bayareen-backend/features/categories"
	mockCategoryRepo "bayareen-backend/features/categories/mocks"
	categoryService "bayareen-backend/features/categories/service"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var (
	categoryRepo       mockCategoryRepo.Data
	categoryServiceVar categories.Business
	categoryCore       categories.Core
	validate           *validator.Validate
)

func setup() {
	categoryServiceVar = categoryService.NewCategoryUsecase(&categoryRepo)
	categoryCore = categories.Core{
		Id:        1,
		Name:      "Pulsa",
		ImgUrl:    "http://google.com",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	validate = validator.New()
}

func TestCreate(t *testing.T) {
	setup()
	categoryRepo.On("Create", mock.AnythingOfType("categories.Core")).Return(categoryCore, nil).Once()
	categoryRepo.On("Create", mock.AnythingOfType("categories.Core")).Return(categories.Core{}, errors.New("can't create category")).Once()

	t.Run("Test Case 1 | Valid Create", func(t *testing.T) {
		result, err := categoryServiceVar.Create(categoryCore)

		assert.Equal(t, categoryCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Missing field in struct", func(t *testing.T) {
		badInput := categories.Core{Name: "First Category"}
		result, err := categoryServiceVar.Create(badInput)

		badInputErr := validate.Struct(&badInput)

		assert.Equal(t, categories.Core{}, result)
		assert.Equal(t, badInputErr.Error(), err.Error())
	})

	t.Run("Test Case 3 | Error on database", func(t *testing.T) {
		result, err := categoryServiceVar.Create(categoryCore)

		assert.Equal(t, categories.Core{}, result)
		assert.Equal(t, errors.New("can't create category"), err)
	})

}

func TestGetAll(t *testing.T) {
	setup()
	categoryRepo.On("GetAll").Return([]categories.Core{categoryCore})

	t.Run("Test Case 1 | Valid Get All", func(t *testing.T) {
		result := categoryServiceVar.GetAll()

		assert.Equal(t, []categories.Core{categoryCore}, result)
	})
}

func TestGetById(t *testing.T) {
	setup()
	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categoryCore, nil).Once()
	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, errors.New("can't find category")).Once()

	t.Run("Test Case 1 | Valid Get By Id", func(t *testing.T) {
		result, err := categoryServiceVar.GetById(1)

		assert.Equal(t, categoryCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Error Get By Id", func(t *testing.T) {
		result, err := categoryServiceVar.GetById(1)

		assert.Equal(t, categories.Core{}, result)
		assert.Equal(t, errors.New("can't find category"), err)
	})
}

func TestUpdate(t *testing.T) {
	setup()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categoryCore, nil).Once()
	categoryRepo.On("Update", mock.AnythingOfType("categories.Core")).Return(categoryCore, nil).Once()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categories.Core{}, errors.New("category doesn't exist")).Once()

	categoryRepo.On("GetById", mock.AnythingOfType("int")).Return(categoryCore, nil).Once()
	categoryRepo.On("Update", mock.AnythingOfType("categories.Core")).Return(categories.Core{}, errors.New("error updating")).Once()

	t.Run("Test Case 1 | Success update", func(t *testing.T) {
		result, err := categoryServiceVar.Update(categoryCore)

		assert.Equal(t, categoryCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Missing field on struct", func(t *testing.T) {
		badInput := categories.Core{}
		badInputErr := validate.Struct(&badInput)

		result, err := categoryServiceVar.Update(badInput)

		assert.Equal(t, categories.Core{}, result)
		assert.Equal(t, badInputErr.Error(), err.Error())
	})

	t.Run("Test Case 3 | Category doesn't exist", func(t *testing.T) {
		result, err := categoryServiceVar.Update(categoryCore)

		assert.Equal(t, categories.Core{}, result)
		assert.Equal(t, errors.New("category doesn't exist"), err)
	})

	t.Run("Test Case 4 | Update failed", func(t *testing.T) {
		result, err := categoryServiceVar.Update(categoryCore)

		assert.Equal(t, categories.Core{}, result)
		assert.Equal(t, errors.New("error updating"), err)
	})
}

func TestDelete(t *testing.T) {
	setup()
	categoryRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

	t.Run("Test Case 1 | Sending error / nil", func(t *testing.T) {
		err := categoryServiceVar.Delete(1)

		assert.Nil(t, err)
	})
}
