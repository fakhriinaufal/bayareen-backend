package presentation

import (
	"bayareen-backend/features/providers"
	provider_request "bayareen-backend/features/providers/presentation/request"
	provider_response "bayareen-backend/features/providers/presentation/response"
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := ph.ProviderBusiness.Create(providerRequest.ToCore())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    provider_response.FromCore(resp),
	})
}

func (ph *ProviderHandler) GetAll(c echo.Context) error {
	resp := ph.ProviderBusiness.GetAll()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    provider_response.FromCoreSlice(resp),
	})
}

func (ph *ProviderHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := ph.ProviderBusiness.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    provider_response.FromCore(resp),
	})
}
