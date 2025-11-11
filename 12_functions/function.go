package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// func add(a , b int) int {
// 	return a + b
// }

func getLang() (string, string) {
	return "java", "c"
}

func process(fn func(a int) int) {
	fn(1)
}

func main() {

	result := add(3, 4)

	fmt.Println(result)

	fn := func(a int) int {
		return 2
	}

	process(fn)
}
