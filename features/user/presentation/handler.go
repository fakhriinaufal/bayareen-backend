package presentation

import (
	"bayareen-backend/features/user"
	presentation_request "bayareen-backend/features/user/presentation/request"
	presentation_response "bayareen-backend/features/user/presentation/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBussiness user.Business
}

func NewUserHandler(ub user.Business) *UserHandler {
	return &UserHandler{
		userBussiness: ub,
	}
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	var userRequest presentation_request.User
	c.Bind(&userRequest)

	resp, err := uh.userBussiness.Create(userRequest.ToCore())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCore(&resp),
	})
}

func (uh *UserHandler) GetAllUser(c echo.Context) error {
	resp := uh.userBussiness.GetAll()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCoreSlice(resp),
	})
}

func (uh *UserHandler) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	resp, err := uh.userBussiness.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCore(&resp),
	})
}

func (uh *UserHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userRequest := presentation_request.User{}

	c.Bind(&userRequest)

	userCore := userRequest.ToCore()
	userCore.Id = id

	resp, err := uh.userBussiness.Update(userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCore(&resp),
	})
}

func (uh *UserHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Print("id", id)

	if err := uh.userBussiness.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    []int{},
	})
}
