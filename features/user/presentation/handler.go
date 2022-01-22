package presentation

import (
	"bayareen-backend/features/user"
	presentation_request "bayareen-backend/features/user/presentation/request"
	presentation_response "bayareen-backend/features/user/presentation/response"
	"bayareen-backend/helper/response"
	"bayareen-backend/middleware"
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.BasicResponse{
		Message: "success",
		Data:    presentation_response.FromCore(&resp),
	})
}

func (uh *UserHandler) GetAllUser(c echo.Context) error {
	resp := uh.userBussiness.GetAll()

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    presentation_response.FromCoreSlice(resp),
	})
}

func (uh *UserHandler) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	resp, err := uh.userBussiness.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    presentation_response.FromCore(&resp),
	})
}

func (uh *UserHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userRequest := presentation_request.User{}

	err := c.Bind(&userRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	userCore := userRequest.ToCore()
	userCore.Id = id

	_, err = uh.userBussiness.Update(userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, []int{})
}

func (uh *UserHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}
	if err := uh.userBussiness.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, []int{})
}

func (uh *UserHandler) Login(c echo.Context) error {
	var user presentation_request.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	userData, err := uh.userBussiness.Login(user.ToCore())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    presentation_response.FromCore(&userData),
	})
}

func (uh *UserHandler) UpdatePassword(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	var userUpdateRequest presentation_request.UserUpdatePasswordPayload

	if err := c.Bind(&userUpdateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	updateRequestData := userUpdateRequest.ToCore()
	updateRequestData.ID = id
	_, err = uh.userBussiness.UpdatePassword(updateRequestData)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (uh *UserHandler) UpdateProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	var updateRequest presentation_request.User
	err = c.Bind(&updateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	userData := updateRequest.ToCore()
	userData.Id = id
	updatedUser, err := uh.userBussiness.UpdateProfile(userData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "OK",
		Data:    presentation_response.FromCore(&updatedUser),
	})
}

func (uh *UserHandler) JWTLogin(c echo.Context) error {
	claims := middleware.ExtractClaim(c)

	id := int(claims["id"].(float64))
	user, err := uh.userBussiness.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    presentation_response.FromCore(&user),
	})
}
