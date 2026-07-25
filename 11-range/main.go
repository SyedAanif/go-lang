package main

import "fmt"

// iterating over data-structures
func main() {
	nums := []int{1, 4, 3, 2, 7}

	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}

	fmt.Println()

	// index, value
	for i, v := range nums {
		fmt.Println(i, v)
	}

	maps := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range maps {
		fmt.Println(k, v)
	}

	// unicode code point
	// here i is actually the num of bytes taken by that character, because some characters can take beyond 1 byte
	str := "golang"
	for i, v := range str {
		fmt.Println(i, v, string(v))
	}
}
