package main

import "fmt"

// pass by value
func wontChangeNum(num int) {
	num = 5
	fmt.Println("In wontChangeNum", num)
}

// pass by reference
func changeNum(num *int) {
	*num = 5 // de-reference, print value of pointer
	fmt.Println("In changeNum", *num)
}

// variable memory location address
func main() {
	num := 1

	wontChangeNum(num)

	fmt.Println("After wontChangeNum in main", num)

	// Pointers
	fmt.Println("Memory Address", &num)

	changeNum(&num)

	fmt.Println("After changeNum in main", num)
}
