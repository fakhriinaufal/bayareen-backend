package presentation

import (
	"bayareen-backend/features/user"
	presentation_request "bayareen-backend/features/user/presentation/request"
	presentation_response "bayareen-backend/features/user/presentation/response"
	"net/http"

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
