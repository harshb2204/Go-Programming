package main

import (
	"fmt"
	"time"
)

// like a blueprint
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// as go doesnt have constructors to implement it here we do it in the following way

func newOrder(id string, amount float32, status string) *order {

	myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myorder
}

// receiver type (behaviour for struct) convention is to write first letter of struct
// here also we need to pass a pointer as by default it passes a value
func (o *order) changeStatus(status string) {
	o.status = status // no need for dereferencing here as struct does it for u
}

// no need to pass a pointer here as we are only getting a value no manipulation
func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	// if u dont set any field default value is 0 value
	// int:0 float:0 string:"" bool:false
	myorder := order{
		id:     "1",
		amount: 50,
		status: "accepted",
	}

	myorder.createdAt = time.Now()
	fmt.Println(myorder.status)

	fmt.Println("Orderstruct ", myorder)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
