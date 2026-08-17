package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

// expect to return the passed value unless error has occurred
// we cna control panic, error log, fatal, must panic etc
func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// MUST functions in GO either return result or panic/handle on errors.
// this can be used to handle errors efficiently
// NOTE: Use this pattern for initialisation, config load, file load etc
func main() {
	// COMMON PATTERN for handling errors
	// r, err := regexp.Compile("[]")
	// if err != nil {
	// 	panic("")
	// 	log.Fatal("")
	// }
	// fmt.Println(r)

	// use MUST function pattern

	reg := Must(regexp.Compile("123"))
	fmt.Println(reg)

	// FILE HANDLING - copying
	src := "./template.txt"
	dest := "./out/template.txt"

	// r, err := os.Open(src)
	// if err != nil {
	// 	panic(err)
	// }
	// defer r.Close()
	r := Must(os.Open(src))
	defer r.Close()

	// w, err := os.Create(dest)
	// if err != nil {
	// 	panic(err)
	// }
	// defer w.Close()
	w := Must(os.Create(dest))
	defer w.Close()

	// if _, err := io.Copy(w, r); err != nil {
	// 	panic(err)
	// }
	Must(io.Copy(w, r))

	// if err := w.Close(); err != nil {
	// 	panic(err)
	// }
	CheckErr(w.Close())
}
