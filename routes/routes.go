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
	user.GET("", presenter.UserPresenter.GetAllUser)
	user.GET("/:id", presenter.UserPresenter.GetUserById)
	user.PATCH("/:id", presenter.UserPresenter.Update)
	user.DELETE("/:id", presenter.UserPresenter.Delete)

	category := e.Group("/categories")
	category.POST("", presenter.CategoryPresenter.CreateCategory)
	category.GET("", presenter.CategoryPresenter.GetAllCategory)
	category.GET("/:id", presenter.CategoryPresenter.GetCategoryById)
	category.PATCH("/:id", presenter.CategoryPresenter.UpdateCategoryById)
	category.DELETE("/:id", presenter.CategoryPresenter.DeleteCategoryById)

	paymentMethod := e.Group("/payment-methods")
	paymentMethod.POST("", presenter.PaymentMethodPresenter.Create)

	return e
}
