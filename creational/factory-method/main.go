package main

import "fmt"

func main() {
	var factory IPaymentFactory

	// Use paypal
	factory = &PaypalFactory{}
	paypal := factory.CreatePayment()
	fmt.Println(paypal.Pay(10000))

	// Use Stripe
	factory = &StripeFactory{}
	stripe := factory.CreatePayment()
	fmt.Println(stripe.Pay(20000))

	// Use MoMo
	factory = &MoMoFactory{}
	momo := factory.CreatePayment()
	fmt.Println(momo.Pay(30000))
}
