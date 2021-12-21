package presentation

import (
	"bayareen-backend/features/categories"
	category_request "bayareen-backend/features/categories/presentation/request"
	category_response "bayareen-backend/features/categories/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryBusiness categories.Business
}

func NewCategoryHandler(cb categories.Business) *CategoryHandler {
	return &CategoryHandler{
		categoryBusiness: cb,
	}
}

func (ch *CategoryHandler) CreateCategory(c echo.Context) error {
	categoryRequest := category_request.Category{}

	err := c.Bind(&categoryRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := ch.categoryBusiness.Create(categoryRequest.ToCore())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    category_response.FromCore(&resp),
	})
}
