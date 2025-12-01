package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {

		fmt.Println("person is an adult")

	} else if age>=13 {
		fmt.Println("teenager")
	} else {
		fmt.Println("kid")
	}

	// logical operators

	var role = "admin"
	var hasPermissions = true

	if role == "admin" && hasPermissions {
		fmt.Println("Yes")
	}

	// declare a variable inside if construct
	if age:=15; age>=18 {
		fmt.Println("person is an adult", age)
	}

	// go does not have ternary operator use normal if-else


	
}