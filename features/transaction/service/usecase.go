package service

import (
	"bayareen-backend/features/email"
	"bayareen-backend/features/payment_gateway"
	"bayareen-backend/features/paymentmethods"
	"bayareen-backend/features/products"
	"bayareen-backend/features/transaction"
	"bayareen-backend/features/user"
	"errors"
	"time"

	"github.com/google/uuid"
)

type transactionUsecase struct {
	PaymentGatewayData payment_gateway.Data
	TransactionData    transaction.Data
	ProductData        products.Data
	UserData           user.Data
	paymentMethodData  paymentmethods.Data
	EmailService       email.Service
}

func NewTransactionUsecase(paymentGatewayData payment_gateway.Data, transactionData transaction.Data, productData products.Data, userData user.Data, paymentMethodData paymentmethods.Data, emailService email.Service) transaction.Business {
	return &transactionUsecase{
		PaymentGatewayData: paymentGatewayData,
		TransactionData:    transactionData,
		ProductData:        productData,
		UserData:           userData,
		paymentMethodData:  paymentMethodData,
		EmailService:       emailService,
	}
}

func (tu *transactionUsecase) Create(data *transaction.Core) (*transaction.Core, error) {
	// get product by id to get price and check existence
	product, err := tu.ProductData.GetById(data.ProductId)
	if err != nil {
		return &transaction.Core{}, err
	}

	// if product is not PDAM/Listrik, use price from product
	if product.Name != "PDAM" && product.Name != "Listrik" {
		data.Price = product.Price
	}

	// check existence of user with specific id
	user, err := tu.UserData.GetById(data.UserId)
	if err != nil {
		return &transaction.Core{}, err
	}
	if user.Id == 0 {
		return &transaction.Core{}, errors.New("user doesn't exist")
	}

	data.Status = "PENDING"
	data.CreatedAt = time.Now()
	referenceId := uuid.NewString()

	// create xendit invoice
	inv, err := tu.PaymentGatewayData.CreateInvoice(payment_gateway.InvoiceObj{
		Id:     referenceId,
		Amount: float64(data.Price),
		Name:   product.Name,
		// Email:       user.Email,
		Description: product.Name,
		Currency:    "IDR",
	})
	if err != nil {
		return &transaction.Core{}, err
	}

	data.ReferenceId = referenceId
	data.InvoiceId = inv.Id
	data.InvoiceUrl = inv.InvoiceUrl
	trans, err := tu.TransactionData.Create(data)
	if err != nil {
		return &transaction.Core{}, err
	}

	// send email to user's email
	mailRequest := email.NewEmailRequest([]string{user.Email}, "Invoice Pembayaran")
	mailData := email.NewInvoiceMailData(user.Name, trans.Price, product.Name, trans.InvoiceUrl)
	// send email with separated goroutine
	go func() {
		err := errors.New("init")
		for err != nil {
			err = tu.EmailService.Send("./features/transaction/service/template/invoice.html", mailRequest, mailData)
		}
	}()

	return trans, nil
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

	trans, err := tu.TransactionData.UpdateByReferenceId(&transaction.Core{
		ReferenceId:     callbackData.ReferenceId,
		Status:          callbackData.Status,
		PaymentMethodId: paymentMethodId,
	})
	if err != nil {
		return err
	}

	user, err := tu.UserData.GetById(trans.UserId)
	if err != nil {
		return err
	}

	mailRequest := email.NewEmailRequest([]string{user.Email}, "Konfirmasi Pembayaran")
	mailData := email.NewPaymentConfirmData(user.Name, trans.Price, trans.Product.Name, callbackData.PaymentChannel, callbackData.PaymentMethod, trans.UpdatedAt)
	go func() {
		err := errors.New("init")
		for err != nil {
			err = tu.EmailService.Send("./features/transaction/service/template/payment.html", mailRequest, mailData)
		}
	}()

	return nil
}

func (tu *transactionUsecase) GetByUserId(userId int) ([]transaction.Core, error) {
	transaction, err := tu.TransactionData.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
