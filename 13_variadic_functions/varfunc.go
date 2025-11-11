package main

import "fmt"

// ...interface{} for getting any type of param
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {
	// result := sum(1,2,3,4,5,6)
	// fmt.Println(result)

	nums := []int{1, 2, 3, 4}

	result := sum(nums...)
	fmt.Println(result)

}
