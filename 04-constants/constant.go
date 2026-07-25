package main

import "fmt"

func main() {
	// complete declaration.
	const name string = "golang"

	// name = "hello" // cannot reinitialise

	// grouped constant
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(name, port, host)
}
