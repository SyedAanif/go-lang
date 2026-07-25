package main

import (
	"fmt"
	"math/rand"
	"time"
)

// send data
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number:", num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, a int, b int) {
	result <- a + b
}

func task(done chan bool) {
	defer func() { done <- true }() // defer ensures that the done channel is sent to even if the above code fails. it must be a function call in defer
	fmt.Println("Processing...")
	// done <- true // marks the task as done, but the above can fail making the code dead. so we use defer
}

// communication in concurrency
func main() {
	// messageChan := make(chan string) // create channel

	// // send data to channel
	// messageChan <- "ping" // blocks there until someone receives the data. so code exits due to deadlock(preemption)

	// // receive data
	// msg := <-messageChan

	// fmt.Println(msg)

	// solved using go-routines or using buffered channels
	numChan := make(chan int)

	go processNum(numChan)

	for i := 0; i < 3; i++ {
		numChan <- rand.Intn(100)
	}

	close(numChan)

	sumChan := make(chan int)
	go sum(sumChan, 1, 2)
	fmt.Println("Sum:", <-sumChan) // blocking

	close(sumChan)

	// synchronisation between go-routines
	doneChan := make(chan bool)

	go task(doneChan)

	<-doneChan // wait for task to complete

	fmt.Println("Task completed!")
	close(doneChan)

	// buffered channels are non-blocking
	emailChan := make(chan string, 3)
	emailChan <- "abc@gmail.com"
	emailChan <- "xyz@gmail.com"
	emailChan <- "pqr@gmail.com"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)

	close(emailChan)

	// muliple channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "Hello"
	}()

	// put in loop, because any of select-case gets evaulated
	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			fmt.Println("Received number:", num)
		case msg := <-chan2:
			fmt.Println("Received message:", msg)
		}
	}

	close(chan1)
	close(chan2)

	// in function parameters we can make a channel sender or reciever only, otherwise channels are duplex
	// func (receiveFromChan <-chan int, sendToChan chan<- bool)  {
	// 	receiveFromChan <- 10 // error
	// 	<-sendToChan // error
	// }

}
