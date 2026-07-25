package main

import (
	"fmt"
	"time"
)

func main() {
	age := 11

	if age >= 20 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("person is a minor")
	}

	role := "admin"
	hasPermissions := true

	if role == "admin" || hasPermissions {
		fmt.Println("authorized")
	} else {
		fmt.Println("unauthorized")
	}

	// same block declaration; has precedence on outer scoped variables
	if age := 15; age >= 12 {
		fmt.Println("person is a teenager", age)
	}

	// switch-case; by default break
	// simple switch
	day := "tuesday"

	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
		// fallthrough // moves to next case except default without breaking
	case "saturday":
		fmt.Println("weekend")
	default:
		fmt.Println("learn days of week")
	}

	// multiple condition
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("work day")
	}

	// type switch
	// type assertion
	whoAmI := func(i interface{}) {
		// switch i.(type){}
		switch t := i.(type) {
		case int:
			fmt.Println("integer", t)
		case string:
			fmt.Println("string", t)
		case bool:
			fmt.Println("boolean", t)
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI(10)
	whoAmI("hello")
	whoAmI(true)
	whoAmI(10.10)
}
