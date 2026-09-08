package main

import (
	"fmt"
	"strings"
)

// here we will build a simple interface and it's implementation.
type FileProcessor interface {
	ParseLines(content string) ([]string, error)
	// Hello() // uncomment to see error in struct definition
}

// implementation
// PROBLEM: as the code base grows larger, it's dificult to track if the interface implementation is valid or not
// as there is no implicit way of errors: ex: if we change interface definition, then only place error comes is the calling section
// not in the struct definition

// to solve this issue, we use explicit TYPE-CAST in the struct interface implementation file
// it's a compile-time check
// empty pointer of CSVProcessor which should always be of it's interface FileProcessor type
var _ FileProcessor = (*CSVProcessor)(nil)

type CSVProcessor struct{}

func (c *CSVProcessor) ParseLines(content string) ([]string, error) {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	return lines, nil
}

func ProcessFile(p FileProcessor, content string) ([]string, error) {
	return p.ParseLines(content)
}

func main() {
	csv := &CSVProcessor{}

	lines, _ := ProcessFile(csv, "name,age \n Alice,25 \n Bob,35")
	fmt.Println(lines)
}
