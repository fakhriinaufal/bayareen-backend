package email

import (
	"fmt"
	"time"
)

type invoiceMailData struct {
	Username   string
	Price      string
	Product    string
	InvoiceUrl string
}

type paymentConfirmMailData struct {
	Username  string
	Price     string
	Product   string
	Payment   string
	UpdatedAt string
}

type Request struct {
	From    string
	To      []string
	Subject string
	Body    string
}

type Service interface {
	Send(templateName string, r *Request, items interface{}) error
}

func NewEmailRequest(to []string, subject string) *Request {
	return &Request{
		To:      to,
		Subject: subject,
	}
}

func NewInvoiceMailData(username string, price int, product string) *invoiceMailData {
	return &invoiceMailData{
		Username: username,
		Price:    "RP " + fmt.Sprintf("%v", price),
		Product:  product,
	}
}

func NewPaymentConfirmData(username string, price int, product string, paymentChannel string, paymentMethod string, updatedAt time.Time) *paymentConfirmMailData {
	return &paymentConfirmMailData{
		Username:  username,
		Price:     "RP " + fmt.Sprintf("%v", price),
		Product:   product,
		Payment:   fmt.Sprintf("%v - %v", paymentMethod, paymentChannel),
		UpdatedAt: updatedAt.String(),
	}
}
