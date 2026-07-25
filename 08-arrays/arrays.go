package main

import "fmt"

// numbered sequence of specific/fixed length
// memory optimisation
func main() {
	// zero values
	// int 0
	// bool false
	// string ""
	// map, slice nil
	var nums [4]int // 4 element array with zero value of type

	// add elements at indices
	nums[0] = 100
	nums[1] = 76
	nums[2] = 2
	nums[3] = 322

	// whole array
	fmt.Println(nums)

	// access elements
	fmt.Println(nums[1])

	// array length
	fmt.Println(len(nums))

	// initialise and declare
	names := [3]string{"a", "b", "c"} // any further would be out-of bounds
	fmt.Println(names)

	// matrix
	mat := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Print(mat)
}
