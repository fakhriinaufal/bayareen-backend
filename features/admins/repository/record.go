package repository

import (
	"bayareen-backend/features/admins"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	Id        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a *Admin) ToCore() *admins.Core {
	return &admins.Core{
		Id:        a.Id,
		Name:      a.Name,
		Password:  a.Password,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func ToCoreSlice(records []Admin) []admins.Core {
	resp := []admins.Core{}
	for _, val := range records {
		resp = append(resp, *val.ToCore())
	}
	return resp
}

func FromCore(data *admins.Core) *Admin {
	return &Admin{
		Id:        data.Id,
		Name:      data.Name,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
