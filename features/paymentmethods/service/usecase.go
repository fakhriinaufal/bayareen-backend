package service

import (
	"bayareen-backend/features/paymentmethods"

	"github.com/go-playground/validator/v10"
)

type paymentMethodUsecase struct {
	PaymentMethodData paymentmethods.Data
	Validator         *validator.Validate
}

func NewPaymentMethodUsecase(paymentMethodData paymentmethods.Data) paymentmethods.Business {
	return &paymentMethodUsecase{
		PaymentMethodData: paymentMethodData,
		Validator:         validator.New(),
	}
}

func (pmu *paymentMethodUsecase) Create(data *paymentmethods.Core) (*paymentmethods.Core, error) {
	err := pmu.Validator.Struct(data)
	if err != nil {
		return &paymentmethods.Core{}, err
	}

	resp, err := pmu.PaymentMethodData.Create(data)

	if err != nil {
		return &paymentmethods.Core{}, err
	}

	return resp, nil
}

func (pmu *paymentMethodUsecase) GetAll() []paymentmethods.Core {
	return pmu.PaymentMethodData.GetAll()
}

func (pmu *paymentMethodUsecase) GetById(id int) (*paymentmethods.Core, error) {
	resp, err := pmu.PaymentMethodData.GetById(id)

	if err != nil {
		return &paymentmethods.Core{}, err
	}

	return resp, nil
}

func (pmu *paymentMethodUsecase) Update(data *paymentmethods.Core) (*paymentmethods.Core, error) {
	existedData, err := pmu.PaymentMethodData.GetById(data.Id)

	if err != nil {
		return &paymentmethods.Core{}, err
	}

	data.CreatedAt = existedData.CreatedAt

	resp, err := pmu.PaymentMethodData.Update(data)

	if err != nil {
		return &paymentmethods.Core{}, err
	}

	return resp, nil
}
