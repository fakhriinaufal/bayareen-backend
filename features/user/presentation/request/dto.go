package request

import "bayareen-backend/features/user"

type User struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (u *User) ToCore() user.UserCore {
	return user.UserCore{
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
		Password:    u.Password,
	}
}

type UserUpdatePasswordPayload struct {
	ID          int    `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (userUpdate *UserUpdatePasswordPayload) ToCore() user.UserUpdatePasswordCore {
	return user.UserUpdatePasswordCore{
		ID:          userUpdate.ID,
		OldPassword: userUpdate.OldPassword,
		NewPassword: userUpdate.NewPassword,
	}
}
