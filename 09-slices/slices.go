package main

import (
	"fmt"
	"slices"
)

// slice: dynamic arrays with variable length
func main() {
	// declare: nil
	var nums []int

	fmt.Println(nums)

	fmt.Println(nums == nil)

	// length and capacity
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	// using make
	var nums2 = make([]int, 2)

	fmt.Println(nums2)

	// length and capacity
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))

	nums2 = append(nums2, 1, 2, 3, 4, 5)

	fmt.Println(nums2)

	// length and capacity
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))

	// declare and initialise
	nums3 := []int{1, 2, 3, 4}

	fmt.Println(nums3)

	fmt.Println(len(nums3))
	fmt.Println(cap(nums3))

	// copying a slice
	nums4 := make([]int, len(nums3))
	copy(nums4, nums3)

	fmt.Println(nums4)

	// slice operator
	nums5 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(nums5[:3])

	fmt.Println(nums5[4:6])

	// slice function
	nums6 := []int{1, 2, 3, 4, 5, 6, 7}
	nums7 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(slices.Equal(nums6, nums7))

	// matrix
	mat := [][]int{
		{1, 2, 3},
		{4, 5},
		{1, 2, 3},
	}

	fmt.Println(mat)

}
