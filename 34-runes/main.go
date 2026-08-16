package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type textAnalyzer struct {
	wordCount  int
	lineCount  int
	charCount  int
	upperCount int
	lowerCount int
}

func NewTextAnalyzer() *textAnalyzer {
	return &textAnalyzer{}
}

// value receiver, as we are just retrieving data -> pass by value
func (ta textAnalyzer) PrintStats() {
	fmt.Println("Total lines: ", ta.lineCount)
	fmt.Println("Total characters: ", ta.charCount)
	fmt.Println("Total words: ", ta.wordCount)
	fmt.Println("Total uppercase: ", ta.upperCount)
	fmt.Println("Total lowercase: ", ta.lowerCount)
}

// pointer receiver, as we are manipulating struct fields -> pass by reference
func (ta *textAnalyzer) Analyze(text string) {
	// // example for len returns the bytes, which can be misleading for multi-byte characters
	// runes := []rune{'A', 'ü', '🚀'}
	// for _, r := range runes {
	// 	str := string(r)             // cast to string
	// 	byteCount := utf8.RuneLen(r) // number of bytes
	// 	fmt.Printf("Rune: %c, UTF-8: %q, Bytes: %d\n", r, str, byteCount)
	// }

	runeCount := utf8.RuneCountInString(text)
	ta.lineCount = runeCount - utf8.RuneCountInString(strings.ReplaceAll(text, "\n", "")) // number of runes - ignore all new line
	ta.charCount = runeCount
	ta.wordCount = len(strings.Fields(text)) // split by whitespace around bytes
	for _, r := range text {
		if unicode.IsUpper(r) {
			ta.upperCount++
		} else if unicode.IsLower(r) {
			ta.lowerCount++
		}
	}

}

// alias for int32, handling unicode points, especially for special non-ASCII text
func main() {
	// var r rune = 'a'
	// fmt.Println(r)

	s := "Hello, üä 🚀" // non-ASCII
	fmt.Println(s)

	r := []rune(s) // convert to slice which contains all unicode points

	fmt.Println(r) // [72 101 108 108 111 44 32 252 228 32 128640]

	for i, r := range s { // converts string to runes
		fmt.Printf("Rune at position: %d, unicode: %v, string: %c\n", i, r, r)
	}

	// used for internationalization, text processing and analysing, as it describes the code point of unicode characters

	text := `Hello, 世界! This is a test.
	Here's some text with uppercase and lowercase letters, and even some emoji: 🚀
	And here's a line with no words to test line
	counting.`

	a := NewTextAnalyzer()

	a.Analyze(text)
	a.PrintStats()
}
