package response

import "bayareen-backend/features/user"

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Token       string `json:"token"`
}

func FromCore(core *user.UserCore) User {
	return User{
		Id:          core.Id,
		Name:        core.Name,
		PhoneNumber: core.PhoneNumber,
		Email:       core.Email,
		Token:       core.Token,
	}
}

func FromCoreSlice(core []user.UserCore) []User {
	userSlice := []User{}
	for _, val := range core {
		userSlice = append(userSlice, FromCore(&val))
	}
	return userSlice
}
