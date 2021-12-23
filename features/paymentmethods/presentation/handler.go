package presentation

import (
	"bayareen-backend/features/paymentmethods"
	_payment_method_request "bayareen-backend/features/paymentmethods/presentation/request"
	_payment_method_response "bayareen-backend/features/paymentmethods/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentMethodHandler struct {
	PaymentMethodBusiness paymentmethods.Business
}

func NewPaymentMethodHandler(paymentMethodBusiness paymentmethods.Business) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		PaymentMethodBusiness: paymentMethodBusiness,
	}
}

func (pmh *PaymentMethodHandler) Create(c echo.Context) error {
	pmRequest := _payment_method_request.PaymentMethod{}

	if err := c.Bind(&pmRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := pmh.PaymentMethodBusiness.Create(pmRequest.ToCore())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message:": "success",
		"data":     _payment_method_response.FromCore(resp),
	})

}
