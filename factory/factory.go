package factory

import (
	"bayareen-backend/driver"
	// user domain
	_userData "bayareen-backend/features/user/data"
	_userHandler "bayareen-backend/features/user/presentation"
	_userUsecase "bayareen-backend/features/user/service"
)

type Presenter struct {
	UserPresenter *_userHandler.UserHandler
}

func Init() Presenter {

	userData := _userData.NewMysqlRepository(driver.DB)
	userUsecase := _userUsecase.NewUserUsecase(userData)
	userHandler := _userHandler.NewUserHandler(userUsecase)

	return Presenter{
		UserPresenter: userHandler,
	}
}
