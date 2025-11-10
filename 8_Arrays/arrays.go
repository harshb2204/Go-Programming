package main

import "fmt"

// numbered sequence of specific length
func main() {

	// default zero values
	var nums [4]int

	nums[0] = 1

	fmt.Println(nums)

	fmt.Println(nums[0])

	fmt.Println(len(nums))

	//declare in single line
	numb := [4]int{1, 2, 3, 4}
	fmt.Println(numb)

	// deafult false
	var vals [4]bool
	fmt.Println(vals)

	// default value " "
	var name [3]string
	fmt.Println(name)

	// 2d arrays
	arr := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(arr)

}
