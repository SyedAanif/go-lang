package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // best practice to saty in the struct
}

func (p *post) inc(wg *sync.WaitGroup) {
	// This will lead to race condition, because concurrent resources will try to modify the same resource
	// defer wg.Done()
	// p.views++
	defer func() {
		wg.Done()
		p.mu.Unlock() // release the lock
	}()
	p.mu.Lock() // synchronize at the stage where state is being modified
	p.views++
}

func main() {
	myPost := post{
		views: 0,
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Printf("Post has %d views\n", myPost.views)
}
