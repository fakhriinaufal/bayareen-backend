package service

import (
	mockEmailService "bayareen-backend/features/email/mocks"
	"bayareen-backend/features/payment_gateway"
	mockGatewayData "bayareen-backend/features/payment_gateway/mocks"
	"bayareen-backend/features/paymentmethods"
	mockPaymentMethodData "bayareen-backend/features/paymentmethods/mocks"
	"bayareen-backend/features/products"
	mockProductData "bayareen-backend/features/products/mocks"
	"bayareen-backend/features/transaction"
	mockTransData "bayareen-backend/features/transaction/mocks"
	"bayareen-backend/features/user"
	mockUserData "bayareen-backend/features/user/mocks"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	gatewayRepo       mockGatewayData.Data
	transRepo         mockTransData.Data
	productRepo       mockProductData.Data
	userRepo          mockUserData.Data
	paymentMethodRepo mockPaymentMethodData.Data
	emailService      mockEmailService.Service
	transService      transaction.Business
	transCoreCreate   *transaction.Core
	transCoreResp     *transaction.Core
	productCore       *products.Core
	paymentMethodCore *paymentmethods.Core
	userCore          user.UserCore
	invoiceData       payment_gateway.InvoiceData
	invoiceObj        payment_gateway.InvoiceObj
)

func setup() {
	transService = NewTransactionUsecase(&gatewayRepo, &transRepo, &productRepo, &userRepo, &paymentMethodRepo, &emailService)
	transCoreCreate = &transaction.Core{
		UserId:    1,
		ProductId: 1,
	}

	transCoreResp = &transaction.Core{
		Id:          1,
		UserId:      1,
		ProductId:   1,
		Price:       50000,
		InvoiceId:   "me912eiajaolwiuer029e2q",
		InvoiceUrl:  "https://invoice-test.com",
		ReferenceId: uuid.NewString(),
	}

	productCore = &products.Core{
		Id:         1,
		ProviderId: 1,
		CatId:      1,
		Name:       "test product",
		Price:      50000,
		Status:     true,
	}

	userCore = user.UserCore{
		Id:          1,
		Name:        "test user",
		PhoneNumber: "0918237289",
		Email:       "user@user.com",
	}

	invoiceObj = payment_gateway.InvoiceObj{
		Id:          transCoreResp.InvoiceUrl,
		Amount:      50000,
		Name:        "test product",
		Description: "test product",
		Currency:    "IDR",
	}

	invoiceData = payment_gateway.InvoiceData{
		Id:         "me912eiajaolwiuer029e2q",
		InvoiceUrl: "https://invoice-test.com",
	}

	paymentMethodCore = &paymentmethods.Core{
		Id:             1,
		PaymentMethod:  "E_WALLET",
		PaymentChannel: "OVO",
	}
}

func TestCreate(t *testing.T) {
	setup()

	t.Run("Test Case 1 | Success Create", func(t *testing.T) {
		productRepo.On("GetById", mock.AnythingOfType("int")).Return(productCore, nil).Once()
		userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCore, nil).Once()
		gatewayRepo.On("CreateInvoice", mock.AnythingOfType("payment_gateway.InvoiceObj")).Return(invoiceData, nil).Once()
		transRepo.On("Create", mock.AnythingOfType("*transaction.Core")).Return(transCoreResp, nil).Once()
		emailService.On("Send", mock.AnythingOfType("string"), mock.AnythingOfType("*email.Request"), mock.AnythingOfType("*email.invoiceMailData")).Return(nil).Once()

		resp, err := transService.Create(transCoreCreate)
		assert.Equal(t, transCoreResp.Id, resp.Id)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Error ProductData.GetById", func(t *testing.T) {
		productRepo.On("GetById", mock.AnythingOfType("int")).Return(&products.Core{}, errors.New("error")).Once()

		resp, err := transService.Create(transCoreCreate)
		assert.Equal(t, &transaction.Core{}, resp)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})

	t.Run("Test Case 3 | Error UserData.GetById", func(t *testing.T) {
		productRepo.On("GetById", mock.AnythingOfType("int")).Return(productCore, nil).Once()
		userRepo.On("GetById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("error")).Once()

		resp, err := transService.Create(transCoreCreate)
		assert.Equal(t, &transaction.Core{}, resp)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})

}

func TestUpdatePayment(t *testing.T) {
	setup()

	t.Run("Test Case 1 | Success UpdatePayment", func(t *testing.T) {
		paymentMethodRepo.On("GetByName", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(1, nil).Once()
		transCoreResp.Status = "PAID"
		transRepo.On("UpdateByReferenceId", mock.AnythingOfType("*transaction.Core")).Return(transCoreResp, nil).Once()
		userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCore, nil).Once()
		emailService.On("Send", mock.AnythingOfType("string"), mock.AnythingOfType("*email.Request"), mock.AnythingOfType("*email.paymentConfirmMailData")).Return(nil).Once()

		err := transService.UpdatePayment(transaction.XenditCallback{})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | PaymentMethod.Id == 0", func(t *testing.T) {
		paymentMethodRepo.On("GetByName", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(0, nil).Once()
		paymentMethodRepo.On("Create", mock.AnythingOfType("*paymentmethods.Core")).Return(paymentMethodCore, nil).Once()
		transCoreResp.Status = "PAID"
		transRepo.On("UpdateByReferenceId", mock.AnythingOfType("*transaction.Core")).Return(transCoreResp, nil).Once()
		userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCore, nil).Once()
		emailService.On("Send", mock.AnythingOfType("string"), mock.AnythingOfType("*email.Request"), mock.AnythingOfType("*email.paymentConfirmMailData")).Return(nil).Once()

		err := transService.UpdatePayment(transaction.XenditCallback{})
		assert.Nil(t, err)
	})

	t.Run("Test Case 3 | Error PaymentMethodData.Create", func(t *testing.T) {
		paymentMethodRepo.On("GetByName", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(0, nil).Once()
		paymentMethodRepo.On("Create", mock.AnythingOfType("*paymentmethods.Core")).Return(&paymentmethods.Core{}, errors.New("error")).Once()

		err := transService.UpdatePayment(transaction.XenditCallback{})
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
}
