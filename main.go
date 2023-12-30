package main

import (
	"fmt"

	"github.com/aminrashidbeigi/go-test-double/gateway"
	"github.com/aminrashidbeigi/go-test-double/payment"
)

func main() {
	realPaymentGateway := &gateway.RealPaymentGateway{}
	paymentProcessor := payment.NewPaymentProcessor(realPaymentGateway)

	amount := 100.0

	err := paymentProcessor.Process(amount)
	if err != nil {
		fmt.Println("Payment failed:", err)
		return
	}

	fmt.Println("Payment processed successfully!")
}
