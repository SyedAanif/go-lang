package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mark task as done
	fmt.Println("task", id)
}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)       // add tasks to wait-group
		go task(i, &wg) // concurrency

		// go func() {
		// 	fmt.Println("closure task", i)
		// }()
	}
	// time.Sleep(time.Second * 2) // to see the output, but better to use sync.WaitGroup

	wg.Wait() // wait for all tasks to complete
	fmt.Println("done")
}
