package main

import "fmt"

type payment struct {
}

func (p payment) makePayement(amount float32) {
	razorpayPayment := razorpay{}
	razorpayPayment.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("Make payment")
}

func main() {

	newPayment := payment{}

	newPayment.makePayement(30)
}
