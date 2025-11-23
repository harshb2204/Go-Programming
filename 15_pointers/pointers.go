package main

import "fmt"

// by value (distinct copy) we are changing the copy here
func changeNum(num int) {
	num = 5
	fmt.Println("Num value in func: ", num)
}

func changeNumByRef(num *int) {
	*num = 5
	fmt.Println("In changeNumByRef: ", *num)
}

func main() {

	num := 1
	fmt.Println("Memory address", &num)
	changeNum(num)

	fmt.Println("After func num value:", num) // still print 1 here
}
