package main

import "fmt"

// for --> only construct in go for looping
func main() {

	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i+1
	// }	

	// infinite loop
	// for {

	// }

	// classic for loop

	// for i:=0; i<3; i++ {

	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// range

	for i:= range 3 { // excludes 3 --> 1 2 
		fmt.Println(i)
	}


}