package main

import "fmt"

// closure can access values from outside it's scope
// stores/memory the variables per closure
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
}
