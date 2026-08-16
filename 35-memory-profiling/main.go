package main

import (
	"fmt"
	"runtime"
)

// get memory statistics before and after memory allocation
func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m) // collects the memory statistics

	fmt.Printf("Alloc = %v MB\n", bToMb(m.Alloc))

	fmt.Printf("Total Alloc = %v MB\n", bToMb(m.TotalAlloc)) // cumulative over time

	fmt.Printf("Sys = %v MB\n", bToMb(m.Sys))

	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1000 / 1000
}

func main() {
	fmt.Println("Mem stats before:")
	printMemStats()

	// simulate memory allocation
	s := make([]int, 10_000_000) // empty slice with 10M records
	for i := range s {
		s[i] = i
	}

	fmt.Println("Mem stats after:")
	printMemStats()

	// force Garbage Collector -> see number of GC and Alloc vs Total Alloc
	runtime.GC()
	fmt.Println("Mem stats  GC:")
	printMemStats()

}
