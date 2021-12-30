package presentation

import (
	"bayareen-backend/features/admins"
	_admin_request "bayareen-backend/features/admins/presentation/request"
	_admin_response "bayareen-backend/features/admins/presentation/response"
	"bayareen-backend/helper/response"
	"net/http"
	"strconv"

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
		return c.JSON(http.StatusBadRequest, response.BasicResponse{
			Message: "failed",
			Data: err.Error(),
		})
	}

	resp, err := ah.adminBusiness.Create(adminRequest.ToCore())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BasicResponse{
			Message: "failed",
			Data: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.BasicResponse{
		Message: "success",
		Data:    _admin_response.FromCore(resp),
	})
}

func (ah *AdminHandler) GetAll(c echo.Context) error {
	resp := ah.adminBusiness.GetAll()
	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    _admin_response.FromCoreSlice(resp),
	})
}

func (ah *AdminHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicResponse{
			Message: "failed",
			Data: err.Error(),
		})
	}

	resp, err := ah.adminBusiness.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BasicResponse{
			Message: "failed",
			Data: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    _admin_response.FromCore(resp),
	})
}

func (ah *AdminHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicResponse{
			Message: "failed",
			Data: err.Error(),
		})
	}

	adminRequest := _admin_request.Admin{}

	if err := c.Bind(&adminRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicResponse{
			Message: "failed",
			Data: err.Error(),
		})
	}

	adminCore := adminRequest.ToCore()
	adminCore.Id = id

	_, err = ah.adminBusiness.Update(adminCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BasicResponse{
			Message: "failed",
			Data: err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, []int{})
}

func (ah *AdminHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicResponse{
			Message: "failed",
			Data: err.Error(),
		})
	}

	err = ah.adminBusiness.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BasicResponse{
			Message: "failed",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, []int{})
}
