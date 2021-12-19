package data

import (
	"bayareen-backend/features/user"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int
	Name        string
	PhoneNumber string
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (u *User) toCore() user.UserCore {
	return user.UserCore{
		Id:          u.ID,
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
		Password:    u.Password,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func toCoreSlice(users []User) []user.UserCore {
	resp := []user.UserCore{}
	for _, val := range users {
		resp = append(resp, val.toCore())
	}
	return resp
}

func fromCore(core user.UserCore) User {
	return User{
		ID:          core.Id,
		Name:        core.Name,
		PhoneNumber: core.PhoneNumber,
		Email:       core.Email,
		Password:    core.Password,
		CreatedAt:   core.CreatedAt,
		UpdatedAt:   core.UpdatedAt,
	}
}
