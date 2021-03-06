package presentation

import (
	"bayareen-backend/features/paymentmethods"
)

type PaymentMethodHandler struct {
	PaymentMethodBusiness paymentmethods.Business
}

// func NewPaymentMethodHandler(paymentMethodBusiness paymentmethods.Business) *PaymentMethodHandler {
// 	return &PaymentMethodHandler{
// 		PaymentMethodBusiness: paymentMethodBusiness,
// 	}
// }

// func (pmh *PaymentMethodHandler) Create(c echo.Context) error {
// 	pmRequest := _payment_method_request.PaymentMethod{}

// 	if err := c.Bind(&pmRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
// 			Message: err.Error(),
// 		})
// 	}

// 	resp, err := pmh.PaymentMethodBusiness.Create(pmRequest.ToCore())

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
// 			Message: err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusCreated, response.BasicResponse{
// 		Message: "success",
// 		Data:     _payment_method_response.FromCore(resp),
// 	})

// }

// func (pmh *PaymentMethodHandler) GetAll(c echo.Context) error {
// 	resp := pmh.PaymentMethodBusiness.GetAll()

// 	return c.JSON(http.StatusOK, response.BasicResponse{
// 		Message: "success",
// 		Data:    _payment_method_response.FromCoreSlice(resp),
// 	})
// }

// func (pmh *PaymentMethodHandler) GetById(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
// 			Message: err.Error(),
// 		})
// 	}

// 	resp, err := pmh.PaymentMethodBusiness.GetById(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
// 			Message: err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, response.BasicResponse{
// 		Message: "success",
// 		Data:    _payment_method_response.FromCore(resp),
// 	})
// }

// func (pmh *PaymentMethodHandler) Update(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
// 			Message: err.Error(),
// 		})
// 	}

// 	pmRequest := _payment_method_request.PaymentMethod{}

// 	if err := c.Bind(&pmRequest); err != nil {
// 		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
// 			Message: err.Error(),
// 		})
// 	}

// 	pmCore := pmRequest.ToCore()
// 	pmCore.Id = id

// 	_, err = pmh.PaymentMethodBusiness.Update(pmCore)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
// 			Message: err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusNoContent, []int{})
// }

// func (pmh *PaymentMethodHandler) Delete(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
// 			Message: err.Error(),
// 		})
// 	}

// 	if err := pmh.PaymentMethodBusiness.Delete(id); err != nil {
// 		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
// 			Message: err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusNoContent, []int{})
// }
