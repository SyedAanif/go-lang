package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
)

type employee struct {
	name string
	age  int
}

// compare (cmp) is used to custom compare values
func main() {
	// returns the first non-false/empty value OR the last value in variadic func
	port := cmp.Or(
		getPortFromEnv(),
		getPortFromFlag(),
		"8080",
	)
	fmt.Println("server started on port:", port)

	// custom sorting algorithm
	employees := []employee{
		{name: "John", age: 25},
		{name: "Jane", age: 30},
		{name: "Doe", age: 35},
		{name: "Alex", age: 25},
	}

	fmt.Println("employees before sorting:", employees)

	sortedEmployees := sortEmployees(employees)

	fmt.Println("employees after sorting:", sortedEmployees)
}

func getPortFromEnv() string {
	return os.Getenv("PORT")
}

func getPortFromFlag() string {
	// return "9090" // is non-empty, hence return 9090
	return "" // is empty, thus port is 8080 from above
}

func sortEmployees(employees []employee) []employee {
	sortedEmployees := employees

	// custom sorting algorithm to sort based on age and if age is equal then name
	slices.SortFunc(sortedEmployees, func(x employee, y employee) int {
		return cmp.Or( // return first non-zero value, if age is equal that means x=y=0, then we sort on name
			cmp.Compare(x.age, y.age),
			cmp.Compare(x.name, y.name),
		)
	})

	return sortedEmployees
}
