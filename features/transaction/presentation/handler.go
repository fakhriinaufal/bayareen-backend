package presentation

import (
	"bayareen-backend/features/transaction"
	"bayareen-backend/features/transaction/presentation/request"
	"bayareen-backend/helper/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	TransactionBusiness transaction.Business
}

func NewTransactionHandler(transactionBusiness transaction.Business) *TransactionHandler {
	return &TransactionHandler{
		TransactionBusiness: transactionBusiness,
	}
}

func (th *TransactionHandler) Create(c echo.Context) error {
	transactionRequest := &request.TransactionRequest{}
	if err := c.Bind(&transactionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	transaction, inv, err := th.TransactionBusiness.Create(transactionRequest.ToCore())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data: map[string]interface{}{
			"transaction": transaction,
			"invoice":     inv,
		},
	})
}

func (th *TransactionHandler) PaymentCallback(c echo.Context) error {
	data := request.XenditCallback{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err = th.TransactionBusiness.UpdatePayment(*data.ToCore())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
	})
}
