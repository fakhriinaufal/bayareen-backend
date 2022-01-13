package presentation

import (
	"bayareen-backend/features/providers"
	provider_request "bayareen-backend/features/providers/presentation/request"
	provider_response "bayareen-backend/features/providers/presentation/response"
	"bayareen-backend/helper/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProviderHandler struct {
	ProviderBusiness providers.Business
}

func NewProviderHandler(pb providers.Business) *ProviderHandler {
	return &ProviderHandler{
		ProviderBusiness: pb,
	}
}

func (ph *ProviderHandler) Create(c echo.Context) error {
	providerRequest := provider_request.Provider{}

	if err := c.Bind(&providerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	resp, err := ph.ProviderBusiness.Create(providerRequest.ToCore())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.BasicResponse{
		Message: "success",
		Data:    provider_response.FromCore(resp),
	})
}

func (ph *ProviderHandler) GetAll(c echo.Context) error {
	catIdStr := c.QueryParam("catId")

	var resp []providers.Core
	var err error
	if catIdStr != "" {
		catId, err := strconv.Atoi(catIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Message: err.Error(),
			})
		}
		resp, err = ph.ProviderBusiness.GetByCategoryId(catId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Message: err.Error(),
			})
		}
	} else {
		resp = ph.ProviderBusiness.GetAll()
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BasicResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    provider_response.FromCoreSlice(resp),
	})
}

func (ph *ProviderHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	resp, err := ph.ProviderBusiness.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    provider_response.FromCore(resp),
	})
}

func (ph *ProviderHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	providerRequest := provider_request.Provider{}

	if err := c.Bind(&providerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	core := providerRequest.ToCore()
	core.Id = id

	_, err = ph.ProviderBusiness.Update(core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, []int{})
}

func (ph *ProviderHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	if err = ph.ProviderBusiness.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BasicResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, []int{})
}
