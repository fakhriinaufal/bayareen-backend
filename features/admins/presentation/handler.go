package presentation

import (
	"bayareen-backend/features/admins"
	_admin_request "bayareen-backend/features/admins/presentation/request"
	_admin_response "bayareen-backend/features/admins/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminBusiness admins.Business
}

func NewAdminHandler(au admins.Business) *AdminHandler {
	return &AdminHandler{
		adminBusiness: au,
	}
}

func (ah *AdminHandler) Create(c echo.Context) error {
	adminRequest := _admin_request.Admin{}

	if err := c.Bind(&adminRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := ah.adminBusiness.Create(adminRequest.ToCore())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    _admin_response.FromCore(resp),
	})
}

func (ah *AdminHandler) GetAll(c echo.Context) error {
	resp := ah.adminBusiness.GetAll()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _admin_response.FromCoreSlice(resp),
	})
}
