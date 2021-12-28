package response

import (
	"bayareen-backend/features/admins"
	"time"
)

type Admin struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password:"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromCore(core *admins.Core) *Admin {
	return &Admin{
		Id:        core.Id,
		Name:      core.Name,
		Email:     core.Email,
		Password:  core.Password,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}

func FromCoreSlice(coreSlice []admins.Core) []Admin {
	resp := []Admin{}
	for _, val := range coreSlice {
		resp = append(resp, *FromCore(&val))
	}
	return resp
}
