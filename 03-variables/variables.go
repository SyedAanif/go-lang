package main

import "fmt"

func main() {
	// long  form
	var name string = "hello"
	fmt.Println(name)

	// inferred
	var inferred = "string"
	fmt.Println(inferred)

	// short hand inferred
	age := 18
	fmt.Println(age)

	// declared and later initialised
	var declare string

	// initilaise
	declare = "initialised"
	fmt.Println(declare)
}
