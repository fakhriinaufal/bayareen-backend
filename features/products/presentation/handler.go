package presentation

import (
	"bayareen-backend/features/products"
	_product_request "bayareen-backend/features/products/presentation/request"
	_product_response "bayareen-backend/features/products/presentation/response"
	"bayareen-backend/helper/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductBusiness products.Business
}

func NewProductHandler(pb products.Business) *ProductHandler {
	return &ProductHandler{
		ProductBusiness: pb,
	}
}

func (ph *ProductHandler) Create(c echo.Context) error {
	productRequest := _product_request.Product{}
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	resp, err := ph.ProductBusiness.Create(productRequest.ToCore())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.BasicResponse{
		Message: "success",
		Data:    _product_response.FromCore(resp),
	})
}

func (ph *ProductHandler) GetAll(c echo.Context) error {
	resp := ph.ProductBusiness.GetAll()
	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    _product_response.FromCoreSlice(resp),
	})
}

func (ph *ProductHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	resp, err := ph.ProductBusiness.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    _product_response.FromCore(resp),
	})
}

func (ph *ProductHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	productRequest := _product_request.Product{}
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	productCore := productRequest.ToCore()
	productCore.Id = id

	_, err = ph.ProductBusiness.Update(productCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, []int{})
}

func (ph *ProductHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	err = ph.ProductBusiness.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusNoContent, []int{})
}
