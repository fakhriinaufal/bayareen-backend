package repository

import (
	"bayareen-backend/features/paymentmethods"
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	Id             int
	PaymentMethod  string
	PaymentChannel string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (p *PaymentMethod) ToCore() *paymentmethods.Core {
	return &paymentmethods.Core{
		Id:             p.Id,
		PaymentMethod:  p.PaymentMethod,
		PaymentChannel: p.PaymentChannel,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
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
		Id:             core.Id,
		PaymentMethod:  core.PaymentMethod,
		PaymentChannel: core.PaymentChannel,
		CreatedAt:      core.CreatedAt,
		UpdatedAt:      core.UpdatedAt,
	}
}
