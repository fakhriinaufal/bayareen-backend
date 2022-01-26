package service_test

import (
	"bayareen-backend/features/paymentmethods"
	mockPaymentMethodRepo "bayareen-backend/features/paymentmethods/mocks"
	"bayareen-backend/features/paymentmethods/service"
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	paymentMethodService paymentmethods.Business
	paymentMethodRepo    mockPaymentMethodRepo.Data
	paymentMethodCore    *paymentmethods.Core
	validate             *validator.Validate
)

func setup() {
	paymentMethodService = service.NewPaymentMethodUsecase(&paymentMethodRepo)
	paymentMethodCore = &paymentmethods.Core{
		Id:             1,
		PaymentMethod:  "BRI",
		PaymentChannel: "Anu",
	}
	validate = validator.New()

}

func TestCreate(t *testing.T) {
	setup()

	paymentMethodRepo.On("Create", mock.AnythingOfType("*paymentmethods.Core")).Return(paymentMethodCore, nil).Once()

	paymentMethodRepo.On("Create", mock.AnythingOfType("*paymentmethods.Core")).Return(&paymentmethods.Core{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success create", func(t *testing.T) {
		result, err := paymentMethodService.Create(paymentMethodCore)

		assert.Equal(t, paymentMethodCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Missing field", func(t *testing.T) {
		badInput := &paymentmethods.Core{}
		badInputErr := validate.Struct(badInput)
		result, err := paymentMethodService.Create(badInput)

		assert.Equal(t, &paymentmethods.Core{}, result)
		assert.Equal(t, badInputErr.Error(), err.Error())
	})

	t.Run("Test Case 3 | Error from database", func(t *testing.T) {
		result, err := paymentMethodService.Create(paymentMethodCore)

		assert.Equal(t, &paymentmethods.Core{}, result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}

func TestGetAll(t *testing.T) {
	setup()

	paymentMethodRepo.On("GetAll").Return([]paymentmethods.Core{*paymentMethodCore}).Once()

	t.Run("Test Case 1 | Success get all", func(t *testing.T) {
		result := paymentMethodService.GetAll()

		assert.Equal(t, []paymentmethods.Core{*paymentMethodCore}, result)
	})
}

func TestGetById(t *testing.T) {
	setup()

	paymentMethodRepo.On("GetById", mock.AnythingOfType("int")).Return(paymentMethodCore, nil).Once()
	paymentMethodRepo.On("GetById", mock.AnythingOfType("int")).Return(&paymentmethods.Core{}, errors.New("payment method doesn't exist")).Once()

	t.Run("Test Case 1 | Success get by id", func(t *testing.T) {
		result, err := paymentMethodService.GetById(1)

		assert.Equal(t, paymentMethodCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Payment Method doesn't exist", func(t *testing.T) {
		result, err := paymentMethodService.GetById(1)

		assert.Equal(t, &paymentmethods.Core{}, result)
		assert.Equal(t, errors.New("payment method doesn't exist"), err)
	})
}

func TestUpdate(t *testing.T) {
	setup()

	paymentMethodRepo.On("Update", mock.AnythingOfType("*paymentmethods.Core")).Return(paymentMethodCore, nil).Once()

	paymentMethodRepo.On("Update", mock.AnythingOfType("*paymentmethods.Core")).Return(&paymentmethods.Core{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success create", func(t *testing.T) {
		result, err := paymentMethodService.Update(paymentMethodCore)

		assert.Equal(t, paymentMethodCore, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Error from database", func(t *testing.T) {
		result, err := paymentMethodService.Update(paymentMethodCore)

		assert.Equal(t, &paymentmethods.Core{}, result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}

func TestDelete(t *testing.T) {
	setup()

	paymentMethodRepo.On("GetById", mock.AnythingOfType("int")).Return(paymentMethodCore, nil).Once()
	paymentMethodRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

	paymentMethodRepo.On("GetById", mock.AnythingOfType("int")).Return(&paymentmethods.Core{}, errors.New("payment method doesn't exist")).Once()

	t.Run("Test Case 1 | Success Delete", func(t *testing.T) {
		err := paymentMethodService.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Payment method doesn't exist ", func(t *testing.T) {
		err := paymentMethodService.Delete(1)

		assert.Equal(t, errors.New("payment method doesn't exist"), err)
	})
}

func TestGetByName(t *testing.T) {
	setup()

	paymentMethodRepo.On("GetByName", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(1, nil).Once()
	paymentMethodRepo.On("GetByName", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(0, errors.New("payment method doesn't exist")).Once()

	t.Run("Test Case 1 | Success get by name", func(t *testing.T) {
		result, err := paymentMethodService.GetByName("bri", "0")

		assert.Equal(t, 1, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Payment method doesn't exist", func(t *testing.T) {
		result, err := paymentMethodService.GetByName("bri", "0")

		assert.Equal(t, 0, result)
		assert.Equal(t, errors.New("payment method doesn't exist"), err)
	})
}
