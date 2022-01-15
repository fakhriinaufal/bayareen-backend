package response

import (
	"bayareen-backend/features/payment_gateway"
	"bayareen-backend/features/transaction"
)

type TransactionResponse struct {
	Id              int    `json:"id"`
	UserId          int    `json:"user_id"`
	ProductId       int    `json:"product_id"`
	InvoiceId       string `json:"invoice_id"`
	InvoiceUrl      string `json:"invoice_url"`
	PaymentMethodId int    `json:"payment_method_id"`
	Status          string `json:"status"`
}

func FromCore(trans *transaction.Core, invoice payment_gateway.InvoiceData) TransactionResponse {
	return TransactionResponse{
		Id:              trans.Id,
		UserId:          trans.UserId,
		ProductId:       trans.ProductId,
		InvoiceId:       trans.InvoiceId,
		InvoiceUrl:      invoice.InvoiceUrl,
		PaymentMethodId: trans.PaymentMethodId,
		Status:          trans.Status,
	}
}
