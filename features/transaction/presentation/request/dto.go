package request

import "bayareen-backend/features/transaction"

type TransactionRequest struct {
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
	Price     int `json:"price"`
}

type XenditCallback struct {
	Id             string `json:"id"`
	ReferenceId    string `json:"external_id"`
	Status         string `json:"status"`
	PaymentMethod  string `json:"payment_method"`
	PaymentChannel string `json:"payment_channel"`
}

func (tr *TransactionRequest) ToCore() *transaction.Core {
	return &transaction.Core{
		UserId:    tr.UserId,
		ProductId: tr.ProductId,
		Price:     tr.Price,
	}
}

func (xc *XenditCallback) ToCore() *transaction.XenditCallback {
	return &transaction.XenditCallback{
		Id:             xc.Id,
		ReferenceId:    xc.ReferenceId,
		Status:         xc.Status,
		PaymentMethod:  xc.PaymentMethod,
		PaymentChannel: xc.PaymentChannel,
	}
}
