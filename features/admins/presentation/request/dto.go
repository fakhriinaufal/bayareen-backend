package request

import "bayareen-backend/features/admins"

type Admin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Admin) ToCore() *admins.Core {
	return &admins.Core{
		Name:     a.Name,
		Password: a.Password,
	}
}
