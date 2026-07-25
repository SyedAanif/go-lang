package main

import "fmt"

// for in the only construct for looping in GO
func main() {
	// while loop
	i := 0
	for i < 10 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// infinite loop
	// for {
	// 	fmt.Print(i, " ")
	// }

	// classic loop
	for i := 0; i < 10; i++{
		fmt.Print(i, " ")
	}

	fmt.Println()

	// break and continue
	for i := 0; i < 10; i++ {
		if i == 8 {
			break
		}
		if i % 2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}

	fmt.Println()

	// range, exclude last
	for i := range 10{
		fmt.Print(i, " ")
	}
}
