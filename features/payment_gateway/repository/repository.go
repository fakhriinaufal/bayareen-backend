package repository

import (
	"bayareen-backend/features/payment_gateway"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type PaymentGatewayRepository struct {
	WriteKey string
	ReadKey  string
}

func NewPaymentGatewayData(writeKey string, readKey string) payment_gateway.Data {
	return &PaymentGatewayRepository{
		WriteKey: writeKey,
		ReadKey:  readKey,
	}
}

func (pr *PaymentGatewayRepository) CreateInvoice(inv payment_gateway.InvoiceObj) (payment_gateway.InvoiceData, error) {
	xendit.Opt.SecretKey = pr.WriteKey
	shouldSendEmail := false
	data := invoice.CreateParams{
		ExternalID:         inv.Id,
		Amount:             inv.Amount,
		SuccessRedirectURL: "https://bayareen.my.id",
		FailureRedirectURL: "https://bayareen.my.id",
		Currency:           inv.Currency,
		PayerEmail:         inv.Email,
		ShouldSendEmail:    &shouldSendEmail,
		Description:        inv.Description,

		Customer: xendit.InvoiceCustomer{
			GivenNames: inv.Name,
			Email:      inv.Email,
		},
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return payment_gateway.InvoiceData{}, err
	}

	return payment_gateway.InvoiceData{
		Id:         resp.ID,
		InvoiceUrl: resp.InvoiceURL,
	}, nil
}

func (pr *PaymentGatewayRepository) GetInvoice(id string) (payment_gateway.InvoiceData, error) {
	xendit.Opt.SecretKey = pr.ReadKey

	data := invoice.GetParams{
		ID: id,
	}

	resp, err := invoice.Get(&data)
	if err != nil {
		return payment_gateway.InvoiceData{}, err
	}

	return payment_gateway.InvoiceData{
		Id:         resp.ID,
		InvoiceUrl: resp.InvoiceURL,
	}, nil
}
