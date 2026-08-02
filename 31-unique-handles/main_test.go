package main

import (
	"testing"
	"unique"
)

// UNIQUE module allows normalisation/cannonicalisation of variables
// this helps us in comparison of normalised forms whoch improves the performance
var (
	s1 = "this is just a really long string that is used to test the performance of the unique handles algorithm"
	s2 = "this is just a really long string that is used to test the performance of the unique handles algorithm"
)

func BenchmarkStringComparison(b *testing.B) {
	// for i := 0; i < b.N; i++ { // for N benchmark iterations
	// 	_ = (s1 == s2)
	// }

	// modern-way
	for b.Loop() {
		_ = (s1 == s2)
	}
}

func BenchmarkUniqueHandleComparison(b *testing.B) {
	h1 := unique.Make(s1)
	h2 := unique.Make(s2)

	for b.Loop() {
		_ = (h1 == h2)
	}
}
