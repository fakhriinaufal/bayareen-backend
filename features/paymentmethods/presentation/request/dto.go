package presentation

type PaymentMethod struct {
	Name string `json:"name"`
}

// func (paymentMethod *PaymentMethod) ToCore() *paymentmethods.Core {
// 	return &paymentmethods.Core{
// 		Name: paymentMethod.Name,
// 	}
// }
