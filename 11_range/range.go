package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	sum := 0

	for _, num := range nums {

		sum = sum + num
		fmt.Println(num)
	}

	for i, num := range nums {

		fmt.Println(num, i)
	}

	m := map[string]string{"name": "harsh", "surname": "badagandi"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// unicode code point rune
	// starting byte of rune
	for i, c := range "harsh" {
		fmt.Println(i, string(c))
	}

}
