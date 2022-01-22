package presentation

import (
	"bayareen-backend/features/categories"
	category_request "bayareen-backend/features/categories/presentation/request"
	category_response "bayareen-backend/features/categories/presentation/response"
	"bayareen-backend/helper/response"
	"bayareen-backend/middleware"
	"net/http"
	"strconv"

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
	claims := middleware.ExtractClaim(c)
	isAdmin := claims["is_admin"].(bool)
	if !isAdmin {
		return c.JSON(http.StatusForbidden, response.ErrorResponse{
			Message: "not allowed for related action",
		})
	}

	categoryRequest := category_request.Category{}

	err := c.Bind(&categoryRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	resp, err := ch.categoryBusiness.Create(categoryRequest.ToCore())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.BasicResponse{
		Message: "success",
		Data:    category_response.FromCore(&resp),
	})
}

func (ch *CategoryHandler) GetAllCategory(c echo.Context) error {
	resp := ch.categoryBusiness.GetAll()

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    category_response.FromCoreSlice(resp),
	})
}

func (ch *CategoryHandler) GetCategoryById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	resp, err := ch.categoryBusiness.GetById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BasicResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    category_response.FromCore(&resp),
	})
}

func (ch *CategoryHandler) UpdateCategoryById(c echo.Context) error {
	claims := middleware.ExtractClaim(c)
	isAdmin := claims["is_admin"].(bool)
	if !isAdmin {
		return c.JSON(http.StatusForbidden, response.ErrorResponse{
			Message: "not allowed for related action",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	categoryRequest := category_request.Category{}
	if err := c.Bind(&categoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	categoryCore := categoryRequest.ToCore()
	categoryCore.Id = id
	_, err = ch.categoryBusiness.Update(categoryCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, []int{})
}

func (ch *CategoryHandler) DeleteCategoryById(c echo.Context) error {
	claims := middleware.ExtractClaim(c)
	isAdmin := claims["is_admin"].(bool)
	if !isAdmin {
		return c.JSON(http.StatusForbidden, response.ErrorResponse{
			Message: "not allowed for related action",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	if err = ch.categoryBusiness.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BasicResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, []int{})
}
