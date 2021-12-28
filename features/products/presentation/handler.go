package presentation

import (
	"bayareen-backend/features/products"
	_product_request "bayareen-backend/features/products/presentation/request"
	_product_response "bayareen-backend/features/products/presentation/response"
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := ph.ProductBusiness.Create(productRequest.ToCore())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    _product_response.FromCore(resp),
	})
}

func (ph *ProductHandler) GetAll(c echo.Context) error {
	resp := ph.ProductBusiness.GetAll()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _product_response.FromCoreSlice(resp),
	})
}

func (ph *ProductHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := ph.ProductBusiness.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _product_response.FromCore(resp),
	})
}
