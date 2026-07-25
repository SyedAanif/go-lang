package main

import "fmt"

// ... and data type
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {

	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)

	// ... and slice
	nums := []int{1, 2, 3, 4, 5}
	result = sum(nums...) // spread operator
	fmt.Println(result)
}
