package main

import (
	"fmt"
	"time"
)

func main() {

	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Print("other")	 	
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	
	default:
		fmt.Println("workday")
	}


	// type switch

	whoAmi := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other", t)			
		}
	}

	whoAmi("golang")
}