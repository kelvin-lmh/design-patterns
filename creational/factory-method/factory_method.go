package main

import (
	"fmt"
)

// -------------------- Payment Gateway Interface --------------------

type IPaymentGateway interface {
	Pay(amount float64) string
}

// -------------------- Concrete Products --------------------

// Paypal
type Paypal struct{}

func (p *Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal", amount)
}

// Stripe
type Stripe struct{}

func (s *Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Stripe", amount)
}

// MoMo
type MoMo struct{}

func (m *MoMo) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using MoMo", amount)
}

// -------------------- Factor Interface--------------------

type IPaymentFactory interface {
	CreatePayment() IPaymentGateway
}

// -------------------- Concrete Factories --------------------

type PaypalFactory struct{}

func (pf *PaypalFactory) CreatePayment() IPaymentGateway {
	return &Paypal{}
}

type StripeFactory struct{}

func (sf *StripeFactory) CreatePayment() IPaymentGateway {
	return &Stripe{}
}

type MoMoFactory struct{}

func (mf *MoMoFactory) CreatePayment() IPaymentGateway {
	return &MoMo{}
}
