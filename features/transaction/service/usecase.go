package service

import (
	"bayareen-backend/features/payment_gateway"
	"bayareen-backend/features/paymentmethods"
	"bayareen-backend/features/products"
	"bayareen-backend/features/transaction"
	"bayareen-backend/features/user"
	"strconv"
	"time"
)

type transactionUsecase struct {
	PaymentGatewayData payment_gateway.Data
	TransactionData    transaction.Data
	ProductData        products.Data
	UserData           user.Data
	paymentMethodData  paymentmethods.Data
}

func NewTransactionUsecase(paymentGatewayData payment_gateway.Data, transactionData transaction.Data, productData products.Data, userData user.Data, paymentMethodData paymentmethods.Data) transaction.Business {
	return &transactionUsecase{
		PaymentGatewayData: paymentGatewayData,
		TransactionData:    transactionData,
		ProductData:        productData,
		UserData:           userData,
		paymentMethodData:  paymentMethodData,
	}
}

func (tu *transactionUsecase) Create(data *transaction.Core) (*transaction.Core, payment_gateway.InvoiceData, error) {
	data.Status = "Pending"
	data.CreatedAt = time.Now()
	trans, err := tu.TransactionData.Create(data)
	if err != nil {
		return &transaction.Core{}, payment_gateway.InvoiceData{}, err
	}

	product, err := tu.ProductData.GetById(data.ProductId)
	if err != nil {
		return &transaction.Core{}, payment_gateway.InvoiceData{}, err
	}

	user, err := tu.UserData.GetById(data.UserId)
	if err != nil {
		return &transaction.Core{}, payment_gateway.InvoiceData{}, err
	}

	inv, err := tu.PaymentGatewayData.CreateInvoice(payment_gateway.InvoiceObj{
		Id:          strconv.Itoa(trans.Id),
		Amount:      float64(product.Price),
		Name:        product.Name,
		Email:       user.Email,
		Description: product.Name,
		Currency:    "IDR",
	})
	if err != nil {
		return &transaction.Core{}, payment_gateway.InvoiceData{}, err
	}

	trans.InvoiceId = inv.Id
	trans, err = tu.TransactionData.Update(trans)
	if err != nil {
		return &transaction.Core{}, payment_gateway.InvoiceData{}, err
	}

	return trans, inv, nil
}

func (tu *transactionUsecase) UpdatePayment(callbackData transaction.XenditCallback) error {
	paymentMethodId, _ := tu.paymentMethodData.GetByName(callbackData.PaymentMethod, callbackData.PaymentChannel)

	if paymentMethodId == 0 {
		method, err := tu.paymentMethodData.Create(&paymentmethods.Core{
			PaymentMethod:  callbackData.PaymentMethod,
			PaymentChannel: callbackData.PaymentChannel,
		})
		if err != nil {
			return err
		}
		paymentMethodId = method.Id
	}

	id, err := strconv.Atoi(callbackData.TransactionId)
	if err != nil {
		return err
	}

	_, err = tu.TransactionData.Update(&transaction.Core{
		Id:              id,
		Status:          callbackData.Status,
		PaymentMethodId: paymentMethodId,
	})
	if err != nil {
		return err
	}

	return nil
}
