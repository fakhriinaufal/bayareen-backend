package routes

import (
	"bayareen-backend/config"
	"bayareen-backend/factory"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	JWTSecret, err := config.LoadJWTSecret(".")
	if err != nil {
		log.Fatal(err)
	}
	presenter := factory.Init()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Logger())
	user := e.Group("/users")
	user.POST("", presenter.UserPresenter.CreateUser)
	user.GET("", presenter.UserPresenter.GetAllUser, middleware.JWT([]byte(config.JWT_KEY)))
	user.GET("/:id", presenter.UserPresenter.GetUserById, middleware.JWT([]byte(config.JWT_KEY)))
	user.PATCH("/:id", presenter.UserPresenter.Update, middleware.JWT([]byte(config.JWT_KEY)))
	user.PATCH("/:id/profile", presenter.UserPresenter.UpdateProfile, middleware.JWT([]byte(config.JWT_KEY)))
	user.PATCH("/:id/password", presenter.UserPresenter.UpdatePassword, middleware.JWT([]byte(config.JWT_KEY)))
	user.DELETE("/:id", presenter.UserPresenter.Delete, middleware.JWT([]byte(config.JWT_KEY)))
	user.POST("/login", presenter.UserPresenter.Login)
	user.GET("/auth", presenter.UserPresenter.JWTLogin, middleware.JWT([]byte(JWTSecret.Secret)))

	provider := e.Group("/providers")
	provider.POST("", presenter.ProviderPresenter.Create, middleware.JWT([]byte(config.JWT_KEY)))
	provider.GET("", presenter.ProviderPresenter.GetAll, middleware.JWT([]byte(config.JWT_KEY)))
	provider.GET("/:id", presenter.ProviderPresenter.GetById, middleware.JWT([]byte(config.JWT_KEY)))
	provider.PATCH("/:id", presenter.ProviderPresenter.Update, middleware.JWT([]byte(config.JWT_KEY)))
	provider.DELETE("/:id", presenter.ProviderPresenter.Delete, middleware.JWT([]byte(config.JWT_KEY)))

	category := e.Group("/categories")
	category.POST("", presenter.CategoryPresenter.CreateCategory, middleware.JWT([]byte(config.JWT_KEY)))
	category.GET("", presenter.CategoryPresenter.GetAllCategory, middleware.JWT([]byte(config.JWT_KEY)))
	category.GET("/:id", presenter.CategoryPresenter.GetCategoryById, middleware.JWT([]byte(config.JWT_KEY)))
	category.PATCH("/:id", presenter.CategoryPresenter.UpdateCategoryById, middleware.JWT([]byte(config.JWT_KEY)))
	category.DELETE("/:id", presenter.CategoryPresenter.DeleteCategoryById, middleware.JWT([]byte(config.JWT_KEY)))
	category.GET("/name", presenter.CategoryPresenter.GetCategoryByName, middleware.JWT([]byte(config.JWT_KEY)))

	// paymentMethod := e.Group("/payment-methods")
	// paymentMethod.POST("", presenter.PaymentMethodPresenter.Create)
	// paymentMethod.GET("", presenter.PaymentMethodPresenter.GetAll)
	// paymentMethod.GET("/:id", presenter.PaymentMethodPresenter.GetById)
	// paymentMethod.PATCH("/:id", presenter.PaymentMethodPresenter.Update)
	// paymentMethod.DELETE("/:id", presenter.PaymentMethodPresenter.Delete)

	product := e.Group("/products")
	product.POST("", presenter.ProductPresenter.Create, middleware.JWT([]byte(config.JWT_KEY)))
	product.GET("", presenter.ProductPresenter.GetAll, middleware.JWT([]byte(config.JWT_KEY)))
	product.GET("/:id", presenter.ProductPresenter.GetById, middleware.JWT([]byte(config.JWT_KEY)))
	product.PATCH("/:id", presenter.ProductPresenter.Update, middleware.JWT([]byte(config.JWT_KEY)))
	product.DELETE("", presenter.ProductPresenter.Delete, middleware.JWT([]byte(config.JWT_KEY)))
	product.GET("/price", presenter.ProductPresenter.GeneratePrice, middleware.JWT([]byte(config.JWT_KEY)))

	admin := e.Group("/admins")
	admin.POST("", presenter.AdminPresenter.Create)
	admin.GET("", presenter.AdminPresenter.GetAll)
	admin.GET("/:id", presenter.AdminPresenter.GetById)
	admin.PATCH("/:id", presenter.AdminPresenter.Update)
	admin.DELETE("/:id", presenter.AdminPresenter.Delete)
	admin.POST("/login", presenter.AdminPresenter.Login)
	admin.GET("/auth", presenter.AdminPresenter.JWTLogin, middleware.JWT([]byte(JWTSecret.Secret)))

	transaction := e.Group("/transactions")
	// transaction.GET("", presenter.TransactionPresenter.GetByUserId, middleware.JWT([]byte(config.JWT_KEY)))
	transaction.POST("", presenter.TransactionPresenter.Create, middleware.JWT([]byte(config.JWT_KEY)))
	transaction.POST("/callbacks", presenter.TransactionPresenter.PaymentCallback)

	return e
}
