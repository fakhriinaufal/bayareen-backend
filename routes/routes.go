package routes

import (
	"bayareen-backend/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	user := e.Group("/users")
	user.POST("", presenter.UserPresenter.CreateUser)

	return e
}
