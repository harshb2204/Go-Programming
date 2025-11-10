package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]string)

	m["name"] = "harsh"
	m["surname"] = "badagandi"

	fmt.Println(m["name"], m["surname"])

	// if key does not exist the it returns zero value

	fmt.Println(len(m))

	delete(m, "surname")

	clear(m)

	mp := map[string]int{"price": 10, "tv": 1}

	fmt.Println(mp)

	v, ok := mp["price"]

	if ok {
		fmt.Println("all good")
	} else {
		fmt.Println("NOT GOOD")
	}

	fmt.Println(v) // value is also returned in the ok

	mp1 := map[string]int{"price": 10, "tv": 1}

	fmt.Println(maps.Equal(mp, mp1))

}
