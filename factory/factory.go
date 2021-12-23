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
)

type Presenter struct {
	UserPresenter     *_userHandler.UserHandler
	ProviderPresenter *_providerHandler.ProviderHandler
	CategoryPresenter *_categoryHandler.CategoryHandler
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

	return Presenter{
		UserPresenter:     userHandler,
		CategoryPresenter: categoryHandler,
    ProviderPresenter: providerHandler,
	}
}
