package main

import "fmt"

func printIntSlice(items []int) {
	for _, v := range items {
		fmt.Print(v, " ")
	}
}

func printStringSlice(items []string) {
	for _, v := range items {
		fmt.Print(v, " ")
	}
}

// thus DRY
// func printSlice[T any or interface{}](items []T) {
// func printSlice[T comparable](items []T) {
func printSlice[T int | string](items []T) { // scoped to int and string
	for _, v := range items {
		fmt.Print(v, " ")
	}
}

type stack[T int | string] struct {
	elements []T
}

// generics introduced in go 1.18
func main() {
	nums := []int{1, 2, 3, 4}
	printIntSlice(nums)

	fmt.Println()

	names := []string{"a", "b", "c", "d"}
	printStringSlice(names)

	fmt.Println()

	printSlice(nums)
	fmt.Println()
	printSlice(names)

	// printSlice([]bool{}) // error

	myStack := stack[int]{
		elements: []int{1, 2, 3, 4},
	}

	fmt.Println(myStack)

}
