package email

import (
	"fmt"
	"html/template"
	"time"
)

type invoiceMailData struct {
	Username   string
	Price      string
	Product    string
	InvoiceUrl template.URL
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

func NewInvoiceMailData(username string, price int, product string, invoiceUrl string) *invoiceMailData {
	return &invoiceMailData{
		Username:   username,
		Price:      "RP " + fmt.Sprintf("%v", price),
		Product:    product,
		InvoiceUrl: template.URL(invoiceUrl),
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
