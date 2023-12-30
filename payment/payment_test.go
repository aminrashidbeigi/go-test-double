package payment_test

import (
	"errors"
	"testing"

	"github.com/aminrashidbeigi/go-test-double/payment"
)

type PaymentGatewayDummy struct{}

func (d *PaymentGatewayDummy) ProcessPayment(amount float64) error {
	return nil
}

type PaymentGatewaySpy struct {
	called bool
	amount float64
}

func (s *PaymentGatewaySpy) ProcessPayment(amount float64) error {
	s.called = true
	s.amount = amount
	return nil
}

type PaymentGatewayFake struct{}

func (f *PaymentGatewayFake) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	return nil
}

type PaymentGatewayStub struct {
	success bool
}

func (s *PaymentGatewayStub) ProcessPayment(amount float64) error {
	if !s.success {
		return errors.New("payment failed")
	}
	return nil
}

type PaymentGatewayMock struct {
	calls      []string
	callAmount float64
}

func (m *PaymentGatewayMock) ProcessPayment(amount float64) error {
	m.calls = append(m.calls, "ProcessPayment")
	m.callAmount = amount
	return nil
}

func (m *PaymentGatewayMock) VerifyCalls(expectedCalls []string) bool {
	if len(expectedCalls) != len(m.calls) {
		return false
	}
	for i, call := range expectedCalls {
		if call != m.calls[i] {
			return false
		}
	}
	return true
}

func TestPaymentProcessor_WithDummy(t *testing.T) {
	paymentProcessor := payment.NewPaymentProcessor(&PaymentGatewayDummy{})

	err := paymentProcessor.Process(100.0)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestPaymentProcessor_WithSpy(t *testing.T) {
	paymentGatewaySpy := &PaymentGatewaySpy{}
	paymentProcessor := payment.NewPaymentProcessor(paymentGatewaySpy)

	amount := 75.0
	_ = paymentProcessor.Process(amount)

	if !paymentGatewaySpy.called {
		t.Error("Expected PaymentGatewaySpy to be called")
	}
	if paymentGatewaySpy.amount != amount {
		t.Errorf("Expected amount %f, got %f", amount, paymentGatewaySpy.amount)
	}
}

func TestPaymentProcessor_WithFake(t *testing.T) {
	paymentProcessor := payment.NewPaymentProcessor(&PaymentGatewayFake{})

	err := paymentProcessor.Process(0)

	if err == nil {
		t.Error("Expected an error for zero amount, got none")
	}

	errNegative := paymentProcessor.Process(-10)

	if errNegative == nil {
		t.Error("Expected an error for negative amount, got none")
	}
}

func TestPaymentProcessor_WithStub_SuccessfulPayment(t *testing.T) {
	paymentProcessor := payment.NewPaymentProcessor(&PaymentGatewayStub{success: true})

	err := paymentProcessor.Process(100.0)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestPaymentProcessor_WithStub_FailedPayment(t *testing.T) {
	paymentProcessor := payment.NewPaymentProcessor(&PaymentGatewayStub{success: false})

	err := paymentProcessor.Process(50.0)

	if err == nil {
		t.Error("Expected an error, got none")
	}
}

func TestPaymentProcessor_WithMock(t *testing.T) {
	paymentGatewayMock := &PaymentGatewayMock{}
	paymentProcessor := payment.NewPaymentProcessor(paymentGatewayMock)

	amount := 75.0
	_ = paymentProcessor.Process(amount)

	expectedCalls := []string{"ProcessPayment"}
	if !paymentGatewayMock.VerifyCalls(expectedCalls) {
		t.Errorf("Expected calls %v, got %v", expectedCalls, paymentGatewayMock.calls)
	}
	if paymentGatewayMock.callAmount != amount {
		t.Errorf("Expected call amount %f, got %f", amount, paymentGatewayMock.callAmount)
	}
}
