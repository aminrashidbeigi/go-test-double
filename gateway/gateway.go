package gateway

import "errors"

type PaymentGateway interface {
	ProcessPayment(amount float64) error
}
type RealPaymentGateway struct{}

func (r *RealPaymentGateway) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid payment amount")
	}
	return nil
}
