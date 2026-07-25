package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string, string) {
	return "Go", "Java", "Python", "JavaScript"
}

// first class citizen -> pass function, return func, assign to var
func processIt(fn func(a int) int) int {
	return fn(10)
}

func funcReturn() func(a int) string {
	return func(a int) string {
		return "first class"
	}
}

func main() {
	fmt.Println(add(42, 13))

	fmt.Println(getLanguages())

	fn := func(a int) int {
		return a * 2
	}

	fmt.Println(processIt(fn))

	returnedFunc := funcReturn()
	fmt.Println(returnedFunc(10))
}
