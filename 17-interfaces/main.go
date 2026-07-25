package main

import "fmt"

// logic for payment
type payment struct {
	// gateway stripe
	gateway paymenter
}

// violates Open-Closed principle
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}

	// razorpayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}

	// stripePaymentGw.pay(amount)

	p.gateway.pay(amount)
}

// actual payment gateways
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	// logic to make payment
	fmt.Println("payment using stripe", amount)
}

type mockPayment struct{}

func (m mockPayment) pay(amount float32) {

	fmt.Println("Testing payment using mocking", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("payment using paypal", amount)
}

// solving all by interface
type paymenter interface {
	pay(amount float32)
}

func main() {
	stripeGw := stripe{}
	// razorpayGw := razorpay{}
	// mock := mockPayment{}
	// newPayment := payment{
	// 	// gateway: stripeGw,
	// 	gateway: razorpayGw,
	// }

	newPayment := payment{
		gateway: stripeGw,
		// gateway: razorpayGw,
		// gateway: mock,
	}

	newPayment.makePayment(100.00)
}
