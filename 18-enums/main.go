package main

import "fmt"

// enums using const
type OrderStatus int

// zero indexed grouped values
const (
	Confirmed OrderStatus = iota
	Processing
	Shipped
	Delivered
	Cancelled
	Returned
	Refunded
	Hold
)

func chageOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

// enumerated data types
func main() {
	chageOrderStatus(Returned)
}
