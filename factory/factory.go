package factory

import (
	"bayareen-backend/config"
	"bayareen-backend/driver"
	"log"

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
	// _paymentMethodHandler "bayareen-backend/features/paymentmethods/presentation"
	_paymentMethodData "bayareen-backend/features/paymentmethods/repository"
	// _paymentMethodUsecase "bayareen-backend/features/paymentmethods/service"

	// product domain
	_productHandler "bayareen-backend/features/products/presentation"
	_productData "bayareen-backend/features/products/repository"
	_productUsecase "bayareen-backend/features/products/service"

	// admin domain
	_adminHandler "bayareen-backend/features/admins/presentation"
	_adminData "bayareen-backend/features/admins/repository"
	_adminUsecase "bayareen-backend/features/admins/service"

	// xendit
	_paymentGatewayData "bayareen-backend/features/payment_gateway/repository"

	// transaction domain
	_transactionHandler "bayareen-backend/features/transaction/presentation"
	_transactionData "bayareen-backend/features/transaction/repository"
	_transactionUsecase "bayareen-backend/features/transaction/service"
)

type Presenter struct {
	UserPresenter     *_userHandler.UserHandler
	CategoryPresenter *_categoryHandler.CategoryHandler
	// PaymentMethodPresenter *_paymentMethodHandler.PaymentMethodHandler
	ProviderPresenter    *_providerHandler.ProviderHandler
	ProductPresenter     *_productHandler.ProductHandler
	AdminPresenter       *_adminHandler.AdminHandler
	TransactionPresenter *_transactionHandler.TransactionHandler
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
	// paymentMethodUsecase := _paymentMethodUsecase.NewPaymentMethodUsecase(paymentMethodData)
	// paymentMethodHandler := _paymentMethodHandler.NewPaymentMethodHandler(paymentMethodUsecase)

	productData := _productData.NewPostgresProductRepository(driver.DB)
	productUsecase := _productUsecase.NewProductUsecase(productData, categoryData, providerData)
	productHandler := _productHandler.NewProductHandler(productUsecase)

	adminData := _adminData.NewPostgresUserRepository(driver.DB)
	adminUsecase := _adminUsecase.NewAdminUsecase(adminData)
	adminHandler := _adminHandler.NewAdminHandler(adminUsecase)

	xenditKey, err := config.LoadXenditKey(".")
	if err != nil {
		log.Fatal(err)
	}

	paymentGatewayData := _paymentGatewayData.NewPaymentGatewayData(xenditKey.WriteKey, xenditKey.ReadKey)

	transactionData := _transactionData.NewPostgresTransactionRepository(driver.DB)
	transactionUsecase := _transactionUsecase.NewTransactionUsecase(paymentGatewayData, transactionData, productData, userData, paymentMethodData)
	transactionHandler := _transactionHandler.NewTransactionHandler(transactionUsecase)

	return Presenter{
		UserPresenter:     userHandler,
		CategoryPresenter: categoryHandler,
		// PaymentMethodPresenter: paymentMethodHandler,
		ProviderPresenter:    providerHandler,
		ProductPresenter:     productHandler,
		AdminPresenter:       adminHandler,
		TransactionPresenter: transactionHandler,
	}
}
