package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// custom errors are made by implementing error interface
type NoTxtError struct { // simulate that no file should be a not txt
	Message string
}

// implement error interface -> Error() string
func (n *NoTxtError) Error() string {
	return n.Message
}

func loadFile(fileName string) (string, error) {
	// return "", fmt.Errorf("go to the unexpexted else block")
	if !strings.HasSuffix(fileName, ".txt") {
		return "", &NoTxtError{
			// Message: "opening a non txt is not allowed",
			"opening a non txt is not allowed", // short hand for position based struct assignment
		}
	}

	f, err := os.Open(fileName)
	if err != nil {
		return "", err // some other error
	}
	defer f.Close()
	return fileName, nil
}

func main() {
	filename, err := loadFile("go.mod") // will return custom error
	// filename, err := loadFile("go.txt") // will return standard error

	// to check the type of error:
	// 1. type assertion and switches
	// 2. errors.Is -> good pre-defined errors but gets ugly for custom struct errors because we need to define the content
	// ex: errors.Is(err, &NoTxtError{Message: "SAME MESSAGE HERE"})
	// 3. thus we use errors.As -> cast an error to a type
	var noTxtErr *NoTxtError
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exist error")
	} else if errors.As(err, &noTxtErr) { // casts error to the specified target(if possible)
		fmt.Println("no txt err:", noTxtErr.Message)
	} else {
		fmt.Println("unexpected error")
	}

	fmt.Println("filename:", filename)
	fmt.Println("error:", err)
}
