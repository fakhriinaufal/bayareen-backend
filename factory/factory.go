package factory

import (
	"bayareen-backend/driver"
	// user domain
	_userHandler "bayareen-backend/features/user/presentation"
	_userData "bayareen-backend/features/user/repository"
	_userUsecase "bayareen-backend/features/user/service"

	// provider domain
	_providerHandler "bayareen-backend/features/providers/presentation"
	_providerData "bayareen-backend/features/providers/repository"
	_providerUsecase "bayareen-backend/features/providers/service"

	// category domain
	_categoryHandler "bayareen-backend/features/categories/presentation"
	_categoryData "bayareen-backend/features/categories/repository"
	_categoryUsecase "bayareen-backend/features/categories/service"

	// payment method domain
	_paymentMethodHandler "bayareen-backend/features/paymentmethods/presentation"
	_paymentMethodData "bayareen-backend/features/paymentmethods/repository"
	_paymentMethodUsecase "bayareen-backend/features/paymentmethods/service"

	// admin domain
	_adminHandler "bayareen-backend/features/admins/presentation"
	_adminData "bayareen-backend/features/admins/repository"
	_adminUsecase "bayareen-backend/features/admins/service"
)

type Presenter struct {
	UserPresenter          *_userHandler.UserHandler
	CategoryPresenter      *_categoryHandler.CategoryHandler
	PaymentMethodPresenter *_paymentMethodHandler.PaymentMethodHandler
	ProviderPresenter      *_providerHandler.ProviderHandler
	AdminPresenter         *_adminHandler.AdminHandler
}

func Init() Presenter {

	userData := _userData.NewMysqlRepository(driver.DB)
	userUsecase := _userUsecase.NewUserUsecase(userData)
	userHandler := _userHandler.NewUserHandler(userUsecase)

	providerData := _providerData.NewPostgresRepository(driver.DB)
	providerUsecase := _providerUsecase.NewProviderUsecase(providerData)
	providerHandler := _providerHandler.NewProviderHandler(providerUsecase)

	categoryData := _categoryData.NewPostgreRepository(driver.DB)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(categoryData)
	categoryHandler := _categoryHandler.NewCategoryHandler(categoryUsecase)

	paymentMethodData := _paymentMethodData.NewPostgresPaymentMethodRepository(driver.DB)
	paymentMethodUsecase := _paymentMethodUsecase.NewPaymentMethodUsecase(paymentMethodData)
	paymentMethodHandler := _paymentMethodHandler.NewPaymentMethodHandler(paymentMethodUsecase)

	adminData := _adminData.NewPostgresUserRepository(driver.DB)
	adminUsecase := _adminUsecase.NewAdminUsecase(adminData)
	adminHandler := _adminHandler.NewAdminHandler(adminUsecase)

	return Presenter{
		UserPresenter:          userHandler,
		CategoryPresenter:      categoryHandler,
		PaymentMethodPresenter: paymentMethodHandler,
		ProviderPresenter:      providerHandler,
		AdminPresenter:         adminHandler,
	}
}
