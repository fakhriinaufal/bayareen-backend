package factory

import (
	"bayareen-backend/driver"
	// user domain
	_userHandler "bayareen-backend/features/user/presentation"
	_userData "bayareen-backend/features/user/repository"
	_userUsecase "bayareen-backend/features/user/service"

	_categoryHandler "bayareen-backend/features/categories/presentation"
	_categoryData "bayareen-backend/features/categories/repository"
	_categoryUsecase "bayareen-backend/features/categories/service"
)

type Presenter struct {
	UserPresenter     *_userHandler.UserHandler
	CategoryPresenter *_categoryHandler.CategoryHandler
}

func Init() Presenter {

	// user domain
	userData := _userData.NewMysqlRepository(driver.DB)
	userUsecase := _userUsecase.NewUserUsecase(userData)
	userHandler := _userHandler.NewUserHandler(userUsecase)

	// category domain
	categoryData := _categoryData.NewPostgreRepository(driver.DB)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(categoryData)
	categoryHandler := _categoryHandler.NewCategoryHandler(categoryUsecase)

	return Presenter{
		UserPresenter:     userHandler,
		CategoryPresenter: categoryHandler,
	}
}
