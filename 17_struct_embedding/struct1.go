package main

import "fmt"

type customer struct {
	name  string
	phone string
}

type order struct {
	id     string
	amount float32
	status string
	customer
}

func main() {

	newcustomer := customer{
		phone: "1234",
		name:  "harsh",
	}

	newOrder := order{
		id:       "1",
		amount:   20,
		status:   "accepted",
		customer: newcustomer,
	}

	fmt.Println(newOrder)

	newOrder.customer.name = "okbacha?"
}
