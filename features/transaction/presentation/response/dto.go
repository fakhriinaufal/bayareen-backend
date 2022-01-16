package response

import (
	"bayareen-backend/features/products/presentation/response"
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

type TransactionProductResponse struct {
	Id              int    `json:"id"`
	UserId          int    `json:"user_id"`
	ProductId       int    `json:"product_id"`
	InvoiceId       string `json:"invoice_id"`
	PaymentMethodId int    `json:"payment_method_id"`
	Status          string `json:"status"`
	Product         response.Product
}

func ToTransactionResponse(trans *transaction.Core) TransactionResponse {
	return TransactionResponse{
		Id:              trans.Id,
		UserId:          trans.UserId,
		ProductId:       trans.ProductId,
		InvoiceId:       trans.InvoiceId,
		InvoiceUrl:      trans.InvoiceUrl,
		PaymentMethodId: trans.PaymentMethodId,
		Status:          trans.Status,
	}
}

func ToTransactionProductResponse(trans *transaction.Core) TransactionProductResponse {
	return TransactionProductResponse{
		Id:              trans.Id,
		UserId:          trans.UserId,
		ProductId:       trans.ProductId,
		InvoiceId:       trans.InvoiceId,
		PaymentMethodId: trans.PaymentMethodId,
		Status:          trans.Status,
		Product:         *response.FromCore(&trans.Product),
	}
}

func ToTransactionProductsResponse(transactions []transaction.Core) []TransactionProductResponse {
	conv := []TransactionProductResponse{}
	for _, val := range transactions {
		conv = append(conv, ToTransactionProductResponse(&val))
	}
	return conv
}
