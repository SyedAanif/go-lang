package main

import (
	"fmt"
	"maps"
)

// maps -> key:value
func main() {
	// declare with nil
	var dict map[string]int

	dict = map[string]int{} // initialise

	// setting an element
	dict["k"] = 1

	fmt.Println(dict["k"])

	// non-existent key returns zeroed value

	// using make
	map1 := make(map[int]string)

	map1[1] = "one"
	map1[2] = "two"

	fmt.Println(map1)

	// delete
	delete(map1, 1)
	fmt.Println(map1)

	// clear the map
	clear(map1)
	fmt.Println(map1)

	// short hand
	map2 := map[string]int{
		"k":  1,
		"k2": 2,
	}
	fmt.Println(map2)

	// check if element exists
	if val, ok := map2["k3"]; ok {
		fmt.Println("Exists", val)
	} else {
		fmt.Println("Doesn't exist")
	}

	// maps functions
	m1 := map[string]int{"age": 20, "height": 180}
	m2 := map[string]int{"age": 20, "height": 180}

	fmt.Println(maps.Equal(m1, m2))
}
