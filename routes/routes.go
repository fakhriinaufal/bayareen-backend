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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	user := e.Group("/users")
	user.POST("", presenter.UserPresenter.CreateUser)
	user.GET("", presenter.UserPresenter.GetAllUser)
	user.GET("/:id", presenter.UserPresenter.GetUserById)
	user.PATCH("/:id", presenter.UserPresenter.Update)
	user.DELETE("/:id", presenter.UserPresenter.Delete)

	provider := e.Group("/providers")
	provider.POST("", presenter.ProviderPresenter.Create)
	provider.GET("", presenter.ProviderPresenter.GetAll)
	provider.GET("/:id", presenter.ProviderPresenter.GetById)
	provider.PATCH("/:id", presenter.ProviderPresenter.Update)
	provider.DELETE("/:id", presenter.ProviderPresenter.Delete)

	category := e.Group("/categories")
	category.POST("", presenter.CategoryPresenter.CreateCategory)
	category.GET("", presenter.CategoryPresenter.GetAllCategory)
	category.GET("/:id", presenter.CategoryPresenter.GetCategoryById)
	category.PATCH("/:id", presenter.CategoryPresenter.UpdateCategoryById)
	category.DELETE("/:id", presenter.CategoryPresenter.DeleteCategoryById)

	// paymentMethod := e.Group("/payment-methods")
	// paymentMethod.POST("", presenter.PaymentMethodPresenter.Create)
	// paymentMethod.GET("", presenter.PaymentMethodPresenter.GetAll)
	// paymentMethod.GET("/:id", presenter.PaymentMethodPresenter.GetById)
	// paymentMethod.PATCH("/:id", presenter.PaymentMethodPresenter.Update)
	// paymentMethod.DELETE("/:id", presenter.PaymentMethodPresenter.Delete)

	product := e.Group("/products")
	product.POST("", presenter.ProductPresenter.Create)
	product.GET("", presenter.ProductPresenter.GetAll)
	product.GET("/:id", presenter.ProductPresenter.GetById)
	product.PATCH("/:id", presenter.ProductPresenter.Update)
	product.DELETE("", presenter.ProductPresenter.Delete)

	admin := e.Group("/admins")
	admin.POST("", presenter.AdminPresenter.Create)
	admin.GET("", presenter.AdminPresenter.GetAll)
	admin.GET("/:id", presenter.AdminPresenter.GetById)
	admin.PATCH("/:id", presenter.AdminPresenter.Update)
	admin.DELETE("/:id", presenter.AdminPresenter.Delete)

	transaction := e.Group("/transactions")
	transaction.POST("", presenter.TransactionPresenter.Create)
	transaction.POST("/callbacks", presenter.TransactionPresenter.PaymentCallback)

	return e
}
