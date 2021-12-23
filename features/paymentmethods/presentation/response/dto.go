package response

import (
	"bayareen-backend/features/paymentmethods"
	"time"
)

type PaymentMethod struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromCore(core *paymentmethods.Core) *PaymentMethod {
	return &PaymentMethod{
		Id:        core.Id,
		Name:      core.Name,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}

func FromCoreSlice(coreSlice []paymentmethods.Core) []PaymentMethod {
	resp := []PaymentMethod{}

	for _, val := range coreSlice {
		resp = append(resp, *FromCore(&val))
	}
	return resp
}
