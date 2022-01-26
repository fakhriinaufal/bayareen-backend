package service_test

import (
	"bayareen-backend/config"
	"bayareen-backend/features/user"
	mockUserRepo "bayareen-backend/features/user/mocks"
	"bayareen-backend/features/user/service"
	"bayareen-backend/helper/bcrypt"
	"bayareen-backend/middleware"
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepo       mockUserRepo.Data
	userService    user.Business
	userCore       user.UserCore
	userCoreResult user.UserCore
	validate       *validator.Validate
	jwtSecret      config.JWTSecret
)

func setup() {
	jwtSecret = config.JWTSecret{Secret: "lorem123"}
	userService = service.NewUserUsecase(&userRepo, jwtSecret)
	validate = validator.New()
	userCore = user.UserCore{
		Id:          1,
		Name:        "John",
		PhoneNumber: "+62194xxxx",
		Email:       "johnDoe@john.com",
		Password:    "lorem",
	}
	hashedPassword, _ := bcrypt.Hash(userCore.Password)
	userCoreResult = userCore

	userCoreResult.Password = hashedPassword
}

func TestCreate(t *testing.T) {
	setup()

	userRepo.On("GetByEmail", mock.AnythingOfType("string")).Return(user.UserCore{}, nil).Once()
	userRepo.On("Create", mock.AnythingOfType("user.UserCore")).Return(userCoreResult, nil).Once()

	userRepo.On("GetByEmail", mock.AnythingOfType("string")).Return(user.UserCore{}, nil).Once()
	userRepo.On("Create", mock.AnythingOfType("user.UserCore")).Return(user.UserCore{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success create user", func(t *testing.T) {
		result, err := userService.Create(userCore)

		assert.Equal(t, true, bcrypt.ValidateHash(userCore.Password, result.Password))
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Error from database", func(t *testing.T) {
		result, err := userService.Create(userCore)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}

func TestGetAll(t *testing.T) {
	setup()

	userRepo.On("GetAll").Return([]user.UserCore{userCoreResult}).Once()

	t.Run("Test Case 1 | Success get all user", func(t *testing.T) {
		result := userService.GetAll()

		assert.Equal(t, []user.UserCore{userCoreResult}, result)
	})
}

func TestGetById(t *testing.T) {
	setup()

	userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCoreResult, nil).Once()
	userRepo.On("GetById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("user doesn't exist")).Once()

	t.Run("Test Case 1 | Success get by id", func(t *testing.T) {
		result, err := userService.GetById(1)

		assert.Equal(t, userCoreResult, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | User doesn't exist", func(t *testing.T) {
		result, err := userService.GetById(1)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("user doesn't exist"), err)
	})

}

// func TestUpdate(t *testing.T) {
// 	setup()

// 	userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCoreResult, nil).Once()
// 	userRepo.On("Update", mock.AnythingOfType("user.UserCore")).Return(userCoreResult, nil).Once()

// 	userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCoreResult, nil).Once()
// 	userRepo.On("Update", mock.AnythingOfType("user.UserCore")).Return(user.UserCore{}, errors.New("error happened")).Once()

// 	t.Run("Test Case 1 | Success update", func(t *testing.T) {
// 		result, err := userService.Update(userCore)

// 		assert.Equal(t, userCoreResult, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Test Case 2 | Missing field", func(t *testing.T) {
// 		badInput := user.UserCore{}
// 		badInputErr := validate.Struct(&badInput)
// 		result, err := userService.Update(badInput)

// 		assert.Equal(t, user.UserCore{}, result)
// 		assert.Equal(t, badInputErr.Error(), err.Error())
// 	})

// 	t.Run("Test Case 3 | Error from database", func(t *testing.T) {
// 		result, err := userService.Update(userCore)

// 		assert.Equal(t, user.UserCore{}, result)
// 		assert.Equal(t, errors.New("error happened"), err)
// 	})
// }

func TestDelete(t *testing.T) {
	setup()

	userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCoreResult, nil).Once()
	userRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

	userRepo.On("GetById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("user doesn't exist")).Once()

	t.Run("Test Case 1 | Success delete", func(t *testing.T) {
		err := userService.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | User doesn't exist", func(t *testing.T) {
		err := userService.Delete(1)

		assert.Equal(t, errors.New("user doesn't exist"), err)
	})
}

func TestLogin(t *testing.T) {
	setup()

	userLogged := userCoreResult

	userLogged.Token, _ = middleware.CreateToken(1, false, jwtSecret.Secret)

	userRepo.On("Login", mock.AnythingOfType("user.UserCore")).Return(userCoreResult, nil).Once()
	userRepo.On("Login", mock.AnythingOfType("user.UserCore")).Return(userCoreResult, nil).Once()
	userRepo.On("Login", mock.AnythingOfType("user.UserCore")).Return(user.UserCore{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success login", func(t *testing.T) {
		result, err := userService.Login(userCore)

		assert.Equal(t, userLogged, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Email Empty", func(t *testing.T) {
		userInput := user.UserCore{}

		result, err := userService.Login(userInput)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("email empty"), err)
	})

	t.Run("Test Case 3 | Password Empty", func(t *testing.T) {
		userInput := user.UserCore{Email: "john@doe.com"}

		result, err := userService.Login(userInput)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("password empty"), err)
	})

	t.Run("Test Case 4 | Wrong password", func(t *testing.T) {
		userInput := user.UserCore{Email: "john@doe.com", Password: "ipsum"}

		result, err := userService.Login(userInput)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("password salah"), err)
	})

	t.Run("Test Case 5 | Error from database", func(t *testing.T) {
		userInput := user.UserCore{Email: "john@doe.com", Password: "ipsum"}

		result, err := userService.Login(userInput)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}

func TestUpdatePassword(t *testing.T) {
	setup()

	userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCoreResult, nil).Once()
	userRepo.On("Update", mock.AnythingOfType("user.UserCore")).Return(userCoreResult, nil).Once()

	userRepo.On("GetById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("user doesn't exist")).Once()

	userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCoreResult, nil).Once()
	userRepo.On("Update", mock.AnythingOfType("user.UserCore")).Return(user.UserCore{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Update", func(t *testing.T) {
		updateReq := user.UserUpdatePasswordCore{
			ID:          1,
			NewPassword: "New",
			OldPassword: "lorem",
		}

		result, err := userService.UpdatePassword(updateReq)

		assert.Equal(t, userCoreResult, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Missing field", func(t *testing.T) {
		badUpdateReq := user.UserUpdatePasswordCore{
			OldPassword: "lorem",
		}
		badUpdateReqErr := validate.Struct(&badUpdateReq)
		result, err := userService.UpdatePassword(badUpdateReq)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, badUpdateReqErr.Error(), err.Error())
	})

	t.Run("Test Case 3 | User doesn't exist", func(t *testing.T) {
		updateReq := user.UserUpdatePasswordCore{
			ID:          1,
			NewPassword: "New",
			OldPassword: "lorem",
		}

		result, err := userService.UpdatePassword(updateReq)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("user doesn't exist"), err)
	})

	t.Run("Test Case 4 | Error update from database", func(t *testing.T) {
		updateReq := user.UserUpdatePasswordCore{
			ID:          1,
			NewPassword: "New",
			OldPassword: "lorem",
		}

		result, err := userService.UpdatePassword(updateReq)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}

func TestUpdateProfile(t *testing.T) {
	setup()

	userRepo.On("GetByEmail", mock.AnythingOfType("string")).Return(user.UserCore{}, nil).Once()
	userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCoreResult, nil).Once()

	userRepo.On("Update", mock.AnythingOfType("user.UserCore")).Return(userCoreResult, nil).Once()

	userRepo.On("GetByEmail", mock.AnythingOfType("string")).Return(user.UserCore{}, nil).Once()
	userRepo.On("GetById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("user doesn't exist")).Once()

	userRepo.On("GetByEmail", mock.AnythingOfType("string")).Return(user.UserCore{}, nil).Once()
	userRepo.On("GetById", mock.AnythingOfType("int")).Return(userCoreResult, nil).Once()
	userRepo.On("Update", mock.AnythingOfType("user.UserCore")).Return(user.UserCore{}, errors.New("error happened")).Once()

	t.Run("Test Case 1 | Success Update", func(t *testing.T) {
		result, err := userService.UpdateProfile(userCore)

		assert.Equal(t, userCoreResult, result)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Name empty", func(t *testing.T) {
		emptyEmailReq := user.UserCore{}
		result, err := userService.UpdateProfile(emptyEmailReq)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("name required"), err)
	})

	t.Run("Test Case 3 | Phone number empty", func(t *testing.T) {
		emptyPhoneNumber := user.UserCore{
			Email: "email@doe.com",
			Name:  "John",
		}
		result, err := userService.UpdateProfile(emptyPhoneNumber)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("phone number required"), err)
	})

	t.Run("Test Case 4 | Email empty", func(t *testing.T) {
		emptyEmailReq := user.UserCore{
			Name:        "john",
			PhoneNumber: "628343",
		}
		result, err := userService.UpdateProfile(emptyEmailReq)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("email required"), err)
	})

	t.Run("Test Case 5 | User doesn't exist", func(t *testing.T) {
		result, err := userService.UpdateProfile(userCore)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("user doesn't exist"), err)
	})

	t.Run("Test Case 6 | Error update from database", func(t *testing.T) {
		result, err := userService.UpdateProfile(userCore)

		assert.Equal(t, user.UserCore{}, result)
		assert.Equal(t, errors.New("error happened"), err)
	})
}
