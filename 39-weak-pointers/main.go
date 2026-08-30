package main

import (
	"fmt"
	"runtime"
	"weak"
)

// WEAK-POINTERS/reference: reference a memory address loosely, so that GC can clean up when no reference point to the address
// example: clean up cache memory when no key points to the local cache memory

type ResourceCache struct {
	// strong reference
	// items map[string]*string // invalidate key-value pair

	// weak reference
	items map[string]weak.Pointer[string]
}

func NewResourceCache() *ResourceCache {
	// Strong reference
	// return &ResourceCache{
	// 	items: make(map[string]*string),
	// }

	// weak reference
	return &ResourceCache{
		items: make(map[string]weak.Pointer[string]),
	}
}

func (rc *ResourceCache) Add(key string, value *string) {
	// Strong reference
	// rc.items[key] = value

	// weak reference
	weakPtr := weak.Make(value)
	rc.items[key] = weakPtr
}

// use pointer for GET for weak case only, because we manipulate map
func (rc *ResourceCache) Get(key string) (*string, bool) {
	// Strong reference
	// value, exists := rc.items[key]
	// return value, exists

	// weak reference
	weakPtr, exists := rc.items[key]
	if !exists {
		return nil, false
	}
	if ptr := weakPtr.Value(); ptr != nil {
		return ptr, true
	}
	delete(rc.items, key)
	return nil, false
}

func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory: %.2f MB\n", float64(m.Alloc)/1024/1024)
}

func main() {
	cache := NewResourceCache()

	// example to determine huge memory leak which is not cleared by even making string-reference to nil
	createBigString := func() *string {
		s := make([]byte, 10<<20) // 10 MB
		str := string(s)
		return &str
	}

	bigData := createBigString()
	cache.Add("big", bigData)
	if _, ok := cache.Get("big"); ok {
		fmt.Println("found value")
	}
	// nullify the big string
	bigData = nil

	fmt.Println("Before GC:")
	printMemoryUsage()

	runtime.GC()

	fmt.Println("After GC:")
	printMemoryUsage()

	if val, ok := cache.Get("big"); ok {
		fmt.Printf("still holding %.2f MB in cache\n", float64(len(*val))/1024/1024)
	}
}
