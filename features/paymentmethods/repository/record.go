package repository

import (
	"bayareen-backend/features/paymentmethods"
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (p *PaymentMethod) ToCore() *paymentmethods.Core {
	return &paymentmethods.Core{
		Id:        p.Id,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToCoreSlice(pmSlice []PaymentMethod) []paymentmethods.Core {
	resp := []paymentmethods.Core{}

	for _, val := range pmSlice {
		resp = append(resp, *val.ToCore())
	}

	return resp
}

func FromCore(core *paymentmethods.Core) *PaymentMethod {
	return &PaymentMethod{
		Id:        core.Id,
		Name:      core.Name,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}
