package main

import "fmt"

// returns a function which returns an int
func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}

}

// Functione when executed go onto the call stack, when the execution it ends it is removed from it (variables also deleted)
// here when when increment is called second time it should have again returned 1 but it returned 2
// happends due to closures, if there is a variable being used inside a function which is of outer scope then
// it is always avaiable in the function even after execution of it

func main() {

	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
