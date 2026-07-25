package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer  customer // struct embedding
}

// constructors
func newOrder(id string, amount float32, status string) *order {
	return &order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
}

// associate behaviour/method to a struct
// pointer receiver
// value receiver
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	myOrder := order{
		id:     "123",
		amount: 12.34,
		status: "pending",
	}
	myOrder.createdAt = time.Now() // can add later also

	fmt.Println(myOrder)

	fmt.Println(myOrder.amount)

	myOrder.changeStatus("paid")

	fmt.Println(myOrder)

	fmt.Println(myOrder.getAmount())

	constrcutedOrder := newOrder("1", 50, "pending")
	fmt.Println(constrcutedOrder)

	anonymousStruct := struct {
		name string
		age  int
	}{
		name: "John",
		age:  30,
	}

	myOrder.customer = customer{
		name:  "John",
		phone: "123456789",
	}

	fmt.Println(myOrder)

	fmt.Println(anonymousStruct)

}
