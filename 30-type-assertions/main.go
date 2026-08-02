package main

import (
	"fmt"
	"io"
	"os"
)

// TYPE ASSERTION can be used to assert:
// 1. concrete type of an interface{}/any
// 2. if a method exists for an interface and then execute
// 3. combine with TYPE SWITCH to determine the first type of the interface/generics

// METHOD CHECKING
type Fooer interface {
	Foo()
}

type MyStruct struct {
}

func (m *MyStruct) Foo() {}

func (m *MyStruct) Bar() {}

// TYPE SWITCH
func process[T any](value T) { // GENERIC function that expects an interface
	switch v := any(value).(type) { // type cast to any and the get the DYNAMIC TYPE. CAST is compile time, TYPE is runtime
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("unknown type")
	}
}

// real-world example
type ReadError struct {
	FileName  string
	Operation string
}

func (e *ReadError) Error() string {
	return fmt.Sprintf("error reading file %s: %s\n", e.FileName, e.Operation)
}

func processFile(fileName string) error { // returns PathError or ReadError
	file, err := os.Open(fileName)
	if err != nil {
		return &os.PathError{ // return different error, so that we can do type assertion and switches
			Path: fileName,
			Op:   "open",
			Err:  err,
		}
	}
	defer file.Close()

	buffer := make([]byte, 1024) // parse into the buffer
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return &ReadError{
			FileName:  fileName,
			Operation: "read",
		}
	}
	return nil
}

func main() {

	// TYPE CHECKING
	// var a interface{} = "hello"
	var a any = "hello" // I can be of ANY type

	// b := int(a) // this is CASTING. assign one type to another comparable type

	if b, ok := a.(string); ok { // TYPE ASSERTING to a string. if we change above to int, the line will !ok
		fmt.Println(b)
	}

	// METHOD CHECKING

	var i Fooer = &MyStruct{} // check if Bar() is part of interface before invoking

	if v, ok := i.(interface{ Bar() }); ok { // if the interface holds the method Bar(), then only execute
		v.Bar()
	}

	// TYPE SWITCHES
	process(3)
	process("hello")
	process(true)

	// real-world example of type assertion and type switches; reading, opening files
	err := processFile("text.txt")
	if err != nil {
		switch e := err.(type) {
		case *os.PathError:
			fmt.Println("path error:", e)
		case *ReadError:
			fmt.Println("read error:", e)
		default:
			fmt.Println("unknown error:", e)
		}

		// TYPE ASSERTION using if-else
		// if _, ok := err.(*os.PathError); ok {
		// 	fmt.Println("path error:", err)
		// } else if _, ok := err.(*ReadError); ok {
		// 	fmt.Println("read", err)
		// } else {
		// 	fmt.Println("unknown error:", err)
		// }
	}
}
