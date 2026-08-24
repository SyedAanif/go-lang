package main

import (
	"fmt"
	"io"
	"os"
)

// process the file based on the user-input
// even though the program joins public folder, but we are still able to exploit the path to move out of public folder
// public/safe.txt
// public/../secret.txt
// This can be prevented in GO 1.24 by doing a restricted present-working-directory using os.Root()
func processFile(userInput string) {
	// path-exploited code
	// filePath := path.Join("public", userInput)

	// data, err := os.ReadFile(filePath)
	// if err != nil {
	// 	panic(err)
	// }

	root, err := os.OpenRoot("public") // restrict the root to public folder
	if err != nil {
		panic(err)
	}
	defer root.Close()

	file, err := root.Open(userInput) // open the file with public as the root
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Contents of file are: ", string(data))
}

// generally before GO 1.24, we could do an attack path traversal, by accessing files outside the scope of current directory
func main() {
	// create files with RW permission
	os.WriteFile("public/safe.txt", []byte("This is a publicly safe file"), 0644)

	os.WriteFile("secret.txt", []byte("TOP SECRET"), 0644)

	fmt.Println("Accessing public file")
	processFile("safe.txt")

	fmt.Println("Accessing secret file")
	processFile("../secret.txt")
	// panic: openat ../secret.txt: path escapes from parent -> after os.root enhancement
}
