package payment

type PaymentGateway interface {
	ProcessPayment(amount float64) error
}

type PaymentProcessor struct {
	gateway PaymentGateway
}

func NewPaymentProcessor(gateway PaymentGateway) *PaymentProcessor {
	return &PaymentProcessor{
		gateway: gateway,
	}
}

func (p *PaymentProcessor) Process(amount float64) error {
	return p.gateway.ProcessPayment(amount)
}
